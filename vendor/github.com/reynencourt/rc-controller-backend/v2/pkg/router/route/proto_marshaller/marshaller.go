package proto_marshaller

import (
	"context"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
	"net/http"
)

var marshaller = protojson.MarshalOptions{
	UseEnumNumbers:  true,
	UseProtoNames:   true,
	EmitUnpopulated: true,
}

func Marshal(m proto.Message) ([]byte, error) {
	d, err := marshaller.Marshal(m)
	if err != nil {
		return nil, err
	}
	return d, err
}

func EncodeProtoMessageResponseJSON(_ context.Context, w http.ResponseWriter, response interface{}) error {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	out := response.([]byte)
	_, err := w.Write(out)
	return err
}
