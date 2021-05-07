package errorlib

import (
	"encoding/json"
)

type localerror struct {
	Err  string `json:"msg"`
	Code string `json:"code"`
}

func (e *localerror) JSON() []byte {
	d, _ := json.Marshal(e)
	return d
}

func (e *localerror) Error() string {

	d, err := json.Marshal(e)
	if err != nil {
		return err.Error()
	}

	return string(d)
}

func Error(code string, err string) error {
	return &localerror{Err: err, Code: code}
}

func Wrap(code string, err error) error {
	return &localerror{Err: err.Error(), Code: code}
}
