package charter

import (
	"bytes"
	"errors"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
)

var ErrChartAlreadyExist = errors.New("chart already exist")

func UploadChart(url, chartPath string) error {
	var client http.Client
	client.Timeout = 15 * time.Second

	chartData, err := ioutil.ReadFile(chartPath)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", url, bytes.NewReader(chartData))
	if err != nil {
		return err
	}

	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	if resp.StatusCode == 201 {
		return nil
	} else if resp.StatusCode == 409 {
		return ErrChartAlreadyExist
	} else {
		return errors.New("Unknown error " + strconv.Itoa(resp.StatusCode))
	}
}
