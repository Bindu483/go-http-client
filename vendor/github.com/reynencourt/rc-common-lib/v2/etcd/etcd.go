package etcd

import (
	"context"
	"crypto/tls"
	"encoding/json"
	"errors"
	recipe "github.com/velann21/etcd/contrib/recipes"
	"time"

	"github.com/velann21/etcd/pkg/transport"
	"google.golang.org/grpc"

	"github.com/hashicorp/vault/api"
	v3 "github.com/velann21/etcd/clientv3"
)

const (
	ClientPEMPath = "/mnt/etcd/ssl/certs/client.pem"
	ClientKeyPath = "/mnt/etcd/ssl/certs/client-key.pem"
	TrustedCAPath = "/mnt/etcd/ssl/certs/ca.pem"
)

type Client struct {
	client *v3.Client
	queue  map[string]*recipe.Queue
}

func New(url []string, certPath string, keyPath string, caPath string) (*Client, error) {
	var tlsConfig *tls.Config
	var err error
	if certPath != "" {
		tlsInfo := transport.TLSInfo{
			CertFile:      certPath,
			KeyFile:       keyPath,
			TrustedCAFile: caPath,
		}

		tlsConfig, err = tlsInfo.ClientConfig()
		if err != nil {
			return nil, err
		}
	}

	cfg := v3.Config{
		Endpoints:   url,
		DialTimeout: 10 * time.Second,
		DialOptions: []grpc.DialOption{grpc.WithBlock()},
		TLS:         tlsConfig,
	}

	c, err := v3.New(cfg)
	if err != nil {
		return nil, err
	}
	if c == nil {
		return nil, errors.New("could not connect to etcd")
	}

	return &Client{
		client: c,
		queue:  make(map[string]*recipe.Queue),
	}, nil
}

func (c *Client) GetClient() *v3.Client {
	return c.client
}

func (c *Client) FetchKeys() (*api.InitResponse, error) {

	var secret api.InitResponse

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	resp, err := c.client.Get(ctx, "/rc-token")
	if err != nil {
		return nil, err
	}

	if resp.Count == 0 {
		return nil, nil
	}

	if len(resp.Kvs) == 0 {
		return nil, nil
	}

	if resp.Kvs == nil {
		return nil, errors.New("node is null")
	}

	err = json.Unmarshal(resp.Kvs[0].Value, &secret)
	if err != nil {
		return nil, err
	}

	return &secret, nil
}

func (c *Client) AddKeys(secret api.InitResponse) error {

	data, err := json.Marshal(&secret)
	if err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

	_, err = c.client.Put(ctx, "/rc-token", string(data))
	cancel()

	if err != nil {
		return err
	}

	return nil
}

func (c *Client) NewQueue(prefix string) (*recipe.Queue, error) {
	queue := recipe.NewQueue(c.client, prefix)
	return queue, nil
}
