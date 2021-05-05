package container

import (
	"fmt"
	"os/exec"
	"strings"

	"github.com/pkg/errors"
)

func (i *ContainerImage) ChangeRegistryDomainAndPush(registryDomain string) error {
	tmp := *i
	tmp.Registry = registryDomain
	s := fmt.Sprintf("docker --config /etc/docker tag %v %v", i.ToString(), tmp.ToString())
	cmd := exec.Command("/bin/sh", "-c", s)
	o, err := cmd.CombinedOutput()
	if err != nil {
		return errors.Wrap(err, string(o))
	}
	return tmp.Push()
}

func (i *ContainerImage) Inspect() ([]byte, error) {
	s := fmt.Sprintf("docker --config /etc/docker inspect %v", i.ToString())
	cmd := exec.Command("/bin/sh", "-c", s)
	o, err := cmd.CombinedOutput()
	if err != nil {
		return nil, errors.Wrap(err, string(o))
	}
	return o, nil
}

func (i *ContainerImage) Pull() error {
	s := fmt.Sprintf("docker --config /etc/docker pull %v", i.ToString())
	cmd := exec.Command("/bin/sh", "-c", s)
	o, err := cmd.CombinedOutput()
	if err != nil {
		return errors.Wrap(err, string(o))
	}
	return nil
}

func (i *ContainerImage) ToString() string {
	var t string
	if i.Tag == "" {
		t = "latest"
	} else {
		t = i.Tag
	}
	if i.Registry == "" {
		return fmt.Sprintf("%s:%s", i.Repository, t)
	}
	return fmt.Sprintf("%s/%s:%s", i.Registry, i.Repository, t)
}

func (i *ContainerImage) Remove() error {
	s := fmt.Sprintf("docker --config /etc/docker image remove %v", i.ToString())
	cmd := exec.Command("/bin/sh", "-c", s)
	e, err := cmd.CombinedOutput()
	if err != nil {
		return errors.Wrap(err, string(e))
	}
	return nil
}

func (i *ContainerImage) Push() error {
	s := fmt.Sprintf("docker --config /etc/docker push %v", i.ToString())
	cmd := exec.Command("/bin/sh", "-c", s)
	e, err := cmd.CombinedOutput()
	if err != nil {
		return errors.Wrap(err, string(e))
	}

	return nil
}

func parseApp(app string) (string, string) {

	imageElements := strings.Split(app, ":")

	if len(imageElements) == 1 {
		return app, "latest"
	}

	return imageElements[0], imageElements[1]
}

func ParseImage(imageRaw string) *ContainerImage {
	var image ContainerImage

	imageElements := strings.Split(imageRaw, "/")

	if len(imageElements) == 1 {

		app, version := parseApp(imageElements[0])

		image.Registry = "docker.io"
		image.Tag = version
		image.Repository = app

		return &image

	} else if len(imageElements) > 0 {

		if strings.Contains(imageElements[0], ".") {
			image.Registry = imageElements[0]

			join := strings.Join(imageElements[1:], "/")

			repositoryDetails := strings.Split(join, ":")

			if len(repositoryDetails) == 1 {
				image.Repository = repositoryDetails[0]
				image.Tag = "latest"
			} else if len(repositoryDetails) == 2 {
				image.Repository = repositoryDetails[0]
				image.Tag = repositoryDetails[1]
			}

		} else {
			image.Registry = "docker.io"
			repositoryDetails := strings.Split(imageRaw, ":")

			if len(repositoryDetails) == 1 {
				image.Repository = imageRaw
				image.Tag = "latest"
			} else if len(repositoryDetails) == 2 {
				image.Repository = repositoryDetails[0]
				image.Tag = repositoryDetails[1]
			}
		}
	}

	return &image
}

func (i *ContainerImage) TagImage(j *ContainerImage) error {
	s := fmt.Sprintf("docker --config /etc/docker tag %s %s ", i.ToString(), j.ToString())
	cmd := exec.Command("/bin/sh", "-c", s)
	e, err := cmd.CombinedOutput()
	if err != nil {
		return errors.Wrap(err, string(e))
	}
	return nil
}
