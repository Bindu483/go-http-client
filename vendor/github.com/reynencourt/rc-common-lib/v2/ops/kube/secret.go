package kube

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"github.com/pkg/errors"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)
import apiErrors "k8s.io/apimachinery/pkg/api/errors"

func (k *Kube) CreateSecret(ctx context.Context, secret *v1.Secret, ns string) error {
	_, err := k.clientset.CoreV1().Secrets(ns).Get(ctx, secret.Name, metav1.GetOptions{})
	if err != nil {
		if apiErrors.IsNotFound(err) {
			_, err = k.clientset.CoreV1().Secrets(ns).Create(ctx, secret, metav1.CreateOptions{})
			if err != nil {
				return err
			}
		}
	}
	return err
}

const DockerPullSecretName = "registry-auth"

// DockerConfigJSON represents a local docker auth config file
// for pulling images.
type DockerConfigJSON struct {
	Auths DockerConfig `json:"auths"`
	// +optional
	HttpHeaders map[string]string `json:"HttpHeaders,omitempty"`
}

// DockerConfig represents the config file used by the docker CLI.
// This config that represents the credentials that should be used
// when pulling images from specific image repositories.
type DockerConfig map[string]DockerConfigEntry

type DockerConfigEntry struct {
	Username string `json:"username,omitempty"`
	Password string `json:"password,omitempty"`
	Email    string `json:"email,omitempty"`
	Auth     string `json:"auth,omitempty"`
}

func (k *Kube) CreateDockerPullSecret(ctx context.Context, username, password, email, ns string, registries []string) error {
	dockerConfig := DockerConfigJSON{
		Auths: make(map[string]DockerConfigEntry),
	}

	for _, r := range registries {
		dockerConfig.Auths[r] = DockerConfigEntry{
			Username: username,
			Password: password,
			Email:    email,
			Auth:     encodeDockerConfigFieldAuth(username, password),
		}
	}

	d, err := json.Marshal(dockerConfig)
	if err != nil {
		return errors.Wrap(err, "failed create dockerconfigjson")
	}

	secret := &v1.Secret{}
	secret.Name = DockerPullSecretName
	secret.Type = v1.SecretTypeDockerConfigJson
	secret.Data = map[string][]byte{
		v1.DockerConfigJsonKey: d,
	}

	_, err = k.clientset.CoreV1().Secrets(ns).Create(ctx, secret, metav1.CreateOptions{})
	return err
}

func encodeDockerConfigFieldAuth(username, password string) string {
	fieldValue := username + ":" + password
	return base64.StdEncoding.EncodeToString([]byte(fieldValue))
}

func (k *Kube) GetSecret(ctx context.Context, name, ns string) (*v1.Secret, error) {
	s, err := k.clientset.CoreV1().Secrets(ns).Get(ctx, name, metav1.GetOptions{})
	if err != nil {
		return nil, err
	}
	return s, nil
}
