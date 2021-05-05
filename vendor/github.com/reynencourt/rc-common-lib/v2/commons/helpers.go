package commons

import (
	"bytes"
	"errors"
	"net/http"
	"time"
)

func Contains(nodeNames []string, nodeName string) bool {
	for _, v := range nodeNames {
		if v == nodeName {
			return true
		}
	}
	return false
}

func HttpRequest(methodType string, Url string, body []byte, headers map[string]string) (*http.Response, error) {
	req, err := http.NewRequest(methodType, Url, bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}
	var client http.Client
	client.Timeout = 15 * time.Second
	for k, v := range headers {
		req.Header.Set(k, v)
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

var (
	// ErrParamMissing required field missing error
	InvalidRequest = errors.New("Invalid Request")
	SomethingWrong = errors.New("Something went wrong")
)
