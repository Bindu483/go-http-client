package utils

import (
	"bytes"
	"context"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc/metadata"
	"math/rand"
	"net/http"
	"os/exec"
	"time"
)

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func RandStringRunes(n int) string {
	rand.Seed(time.Now().UnixNano())
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

type HttpClient interface {
	Do(r http.Request) (http.Response, error)
}

func HttpRequest(methodType string, Url string, body []byte) (*http.Response, error) {
	req, err := http.NewRequest(methodType, Url, bytes.NewReader(body))
	if err != nil {
		return nil, err
	}
	var client http.Client
	client.Timeout = 15 * time.Second
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func ExtractHeader(ctx context.Context) (string, string) {
	headers, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		logrus.Info("could not get metadata")
	}
	logrus.Info("Headers are: ", headers)
	traceIDs := headers["x-traceid"]
	eventTypes := headers["x-eventtype"]
	traceID := ""
	eventType := ""
	for _, v := range traceIDs {
		traceID = v
	}

	for _, v := range eventTypes {
		eventType = v
	}
	logrus.Info("TraceID inside EH : ", traceID)
	logrus.Info("EventType inside EH : ", eventType)
	return traceID, eventType
}

func Shell(cmds []string) error {

	for _, cmd := range cmds {

		logrus.Println(cmd)
		exeCmd := exec.Command("/bin/sh", "-c", cmd)

		out, err := exeCmd.CombinedOutput()
		if err != nil {
			logrus.Println(string(out))
			return err
		}
	}

	return nil

}
