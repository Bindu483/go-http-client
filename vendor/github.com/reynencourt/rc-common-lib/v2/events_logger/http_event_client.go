package events_logger

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/url"
	"path"
	"time"
)

type httpEventClient struct {
	collectorHost string
	client        http.Client
}

func NewHttpEventClient(host string, timeout int) *httpEventClient {
	return &httpEventClient{
		collectorHost: host,
		client:        http.Client{Timeout: time.Duration(timeout) * time.Second},
	}
}

func (c *httpEventClient) Push(t *Transaction) error {
	u, err := url.Parse(c.collectorHost)
	if err != nil {
		return err
	}
	u.Path = path.Join(u.Path, "/api/v1/data-collector/collector/distributed_tracing")
	buff, err := json.Marshal(t)
	if err != nil {
		return err
	}

	r, err := http.NewRequest("POST", u.String(), bytes.NewReader(buff))
	if err != nil {
		return err
	}

	res, err := c.client.Do(r)
	if err != nil {
		return err
	}

	if res.StatusCode != http.StatusOK {
		return errors.New("failed to push event")
	}
	return nil
}
