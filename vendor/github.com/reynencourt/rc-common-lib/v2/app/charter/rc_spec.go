package charter

import (
	"github.com/pkg/errors"
	"github.com/reynencourt/rc-common-lib/v2/proto/rc_spec"
	"sigs.k8s.io/yaml"
)

func (c *HelmChart) GetRcSpec() (*rc_spec.RCSpec, error) {
	for _, v := range c.Files {
		if v.Name == "rcspec.yaml" || v.Name == "rcspec.yml" {
			var rcSpec rc_spec.RCSpec
			err := yaml.Unmarshal(v.Data, &rcSpec)
			if err != nil {
				return nil, err
			}
			return &rcSpec, nil
		}
	}
	return nil, errors.New("rcspec not found")
}

func (c *HelmChart) GetRcPipe() (*rc_spec.RCPipeline, error) {
	for _, v := range c.Files {
		if v.Name == "rcpipeline.yaml" || v.Name == "rcpipeline.yml" {
			var rcPipeline rc_spec.RCPipeline
			err := yaml.Unmarshal(v.Data, &rcPipeline)
			if err != nil {
				return nil, err
			}
			return &rcPipeline, nil
		}
	}
	return nil, errors.New("rcpipeline not found")
}
