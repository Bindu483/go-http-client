package utils

import (
	"bytes"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func MakeHTTPRequest(url, method string, payload []byte, headers map[string]string) ([]byte, error) {
	req, err := http.NewRequest(method, url, bytes.NewReader(payload))
	if err != nil {
		return nil, err
	}
	for k, v := range headers {
		req.Header.Add(k, v)
	}
	c := http.Client{
		Timeout: 10 * time.Second,
	}

	res, err := c.Do(req)
	if err != nil {
		return nil, err
	}

	resBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	if res.StatusCode >= 400 {
		return nil, errors.New(fmt.Sprintf("Err: %s StatusCode:%d", string(resBody), res.StatusCode))
	}

	return resBody, nil
}
