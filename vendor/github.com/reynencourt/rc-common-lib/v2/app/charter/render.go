package charter

import (
	"errors"
	"fmt"
	"github.com/reynencourt/helm/v3/pkg/action"
	"github.com/reynencourt/helm/v3/pkg/chart"
	"github.com/reynencourt/helm/v3/pkg/chartutil"
	kubefake "github.com/reynencourt/helm/v3/pkg/kube/fake"
	"github.com/reynencourt/helm/v3/pkg/releaseutil"
	"github.com/reynencourt/helm/v3/pkg/storage"
	"github.com/reynencourt/helm/v3/pkg/storage/driver"
	"io/ioutil"
	"regexp"
	"strings"
)

//values should be both rc_spec and platform environment variables
func (c *HelmChart) Render(values map[string]interface{}) (RenderedChart, error) {
	helmChart := chart.Chart(*c)
	var (
		namespace   string
		releaseName string
		err         error
	)
	coalescedValues, err := c.CoalesceDeploymentValues(values)
	if err != nil {
		return nil, err
	}
	for k, v := range coalescedValues {
		if k == "Release" {
			nestedMap, ok := v.(chartutil.Values)
			if ok {
				releaseName = nestedMap["Name"].(string)
				namespace = nestedMap["Namespace"].(string)
				break
			} else {
				fmt.Println("damn it .....")
			}
		}
	}

	actionConfig := &action.Configuration{
		Releases:     storage.Init(driver.NewMemory()),
		KubeClient:   &kubefake.FailingKubeClient{PrintingKubeClient: kubefake.PrintingKubeClient{Out: ioutil.Discard}},
		Capabilities: chartutil.DefaultCapabilities,
		Log:          func(format string, v ...interface{}) {},
	}

	client := action.NewInstall(actionConfig)
	client.Namespace = namespace
	client.ReleaseName = releaseName
	client.ClientOnly = true

	release, err := client.Run(&helmChart, coalescedValues)
	if err != nil {
		return nil, err
	}

	if release.Manifest == "" {
		return nil, errors.New("failed to render chart")
	}

	manifestNameRegex := regexp.MustCompile("# Source: [^/]+/(.+)")

	splitManifests := releaseutil.SplitManifests(strings.TrimSpace(release.Manifest))
	manifestsKeys := make([]string, 0, len(splitManifests))
	for k := range splitManifests {
		manifestsKeys = append(manifestsKeys, k)
	}

	var mapped = map[string][]string{}
	for _, manifest := range manifestsKeys {
		submatch := manifestNameRegex.FindStringSubmatch(splitManifests[manifest])

		if len(submatch) != 0 {
			//key := strings.ReplaceAll(submatch[0], fmt.Sprintf("# Source: %s/",c.Metadata.Name), "")
			mapped[submatch[1]] = append(mapped[submatch[1]], splitManifests[manifest])
		}
	}

	for _, m := range release.Hooks {
		key := strings.ReplaceAll(m.Path, fmt.Sprintf("%s/", c.Metadata.Name), "")
		mapped[key] = append(mapped[key], m.Manifest)
	}

	return mapped, nil

}
