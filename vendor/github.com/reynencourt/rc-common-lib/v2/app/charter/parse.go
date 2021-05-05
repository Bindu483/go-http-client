package charter

import (
	"fmt"
	"github.com/pkg/errors"
	"k8s.io/client-go/kubernetes/scheme"
	"sigs.k8s.io/yaml"
	"strings"
)

func (r RenderedChart) Parse() (ParsedHelmChart, error) {
	var parsedHelmChart = make(ParsedHelmChart, len(r))
	for k, v := range r {
		if strings.Contains(k, ".yaml") || strings.Contains(k, ".yml") {
			var parsedK8sManifests = make([]K8sObject, len(v))
			for i, obj := range v {
				if o := strings.TrimSpace(obj); o != "" {
					var yamlData map[string]interface{}
					err := yaml.Unmarshal([]byte(obj), &yamlData)
					if err != nil {
						return nil, errors.New(fmt.Sprintf("failed to parse template %s err: %s", k, err.Error()))
					}
					if yamlData == nil || len(yamlData) == 0 {
						continue
					}
					k8sObject, err := deserializeK8sObject([]byte(o))
					if err != nil {
						return nil, err
					}
					parsedK8sManifests[i] = k8sObject
				}
			}
			parsedHelmChart[k] = parsedK8sManifests
		}
	}
	return parsedHelmChart, nil
}

func deserializeK8sObject(rawObj []byte) (K8sObject, error) {
	decode := scheme.Codecs.UniversalDeserializer().Decode
	obj, _, err := decode(rawObj, nil, nil)
	return K8sObject(obj), err
}

func serializeK8sObject(object K8sObject) ([]byte, error) {
	return yaml.Marshal(object)
}
