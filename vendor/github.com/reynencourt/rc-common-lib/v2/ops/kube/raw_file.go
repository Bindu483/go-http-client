package kube

import (
	"fmt"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"os"
	"os/exec"
)

const ManifestPath = "/k8s/manifests"

func (k Kube) CreateFromFile(d []byte) error {
	logrus.Info(string(d))
	err := os.MkdirAll(ManifestPath, os.ModePerm)
	if err != nil {
		return err
	}
	manifestFileName := fmt.Sprintf("%s/%s.yml", ManifestPath, "manifest")
	err = ioutil.WriteFile(manifestFileName, d, os.ModePerm)
	if err != nil {
		return err
	}
	kubeConfigFileName := fmt.Sprintf("%s/%s.yml", ManifestPath, "kubeconfig")
	err = ioutil.WriteFile(kubeConfigFileName, k.config, os.ModePerm)
	if err != nil {
		return err
	}

	defer func() {
		_ = os.Remove(manifestFileName)
		_ = os.Remove(kubeConfigFileName)
	}()

	command := exec.Command("kubectl", "create", "-f", manifestFileName)
	command.Env = append(os.Environ(), fmt.Sprintf("KUBECONFIG=%s", kubeConfigFileName))
	logrus.Info(command.String())
	out, err := command.CombinedOutput()
	if err != nil {
		return errors.Wrap(err, string(out))
	}
	return nil
}

func (k Kube) DeleteFile(d []byte) error {
	err := os.MkdirAll(ManifestPath, os.ModePerm)
	if err != nil {
		return err
	}
	manifestFileName := fmt.Sprintf("%s/%s.yml", ManifestPath, "manifest")
	err = ioutil.WriteFile(manifestFileName, d, os.ModePerm)
	if err != nil {
		return err
	}
	logrus.Info(string(d))
	kubeConfigFileName := fmt.Sprintf("%s/%s.yml", ManifestPath, "kubeconfig")
	err = ioutil.WriteFile(kubeConfigFileName, k.config, os.ModePerm)
	if err != nil {
		return err
	}
	defer func() {
		_ = os.Remove(manifestFileName)
		_ = os.Remove(kubeConfigFileName)
	}()
	command := exec.Command("kubectl", "delete", "-f", manifestFileName)
	command.Env = append(os.Environ(), fmt.Sprintf("KUBECONFIG=%s", kubeConfigFileName))
	logrus.Info(command.String())
	out, err := command.CombinedOutput()
	if err != nil {
		return errors.Wrap(err, string(out))
	}
	return nil
}
