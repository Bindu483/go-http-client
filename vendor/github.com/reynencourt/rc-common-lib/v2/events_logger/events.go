package events_logger

import (
	"encoding/json"
	"strings"
	"time"
)

type EventClient interface {
	Push(t *Transaction) error
}

type ActionType string

var ActionTypeRequest ActionType = "Request"
var ActionTypeResponse ActionType = "Response"

type EventMetadata map[string]string
type EventType string

type Transaction struct {
	Time        time.Time     `json:"time"`
	ActionType  ActionType    `json:"actionType"`
	TraceID     string        `json:"traceId,omitempty"`
	EventType   string        `json:"eventType"`
	Metadata    EventMetadata `json:"metadata"`
	ErrorCode   string        `json:"errorCode,omitempty"`
	ServiceName string        `json:"serviceName"`
	client      EventClient   `json:"-"`
}

type Event struct {
	Transaction Transaction
	service     string
	client      EventClient
}

func NewEvent(sName string, client EventClient) *Event {
	return &Event{client: client, service: sName}
}

func (d *Event) Request() *Transaction {
	return &Transaction{
		ActionType:  ActionTypeRequest,
		ServiceName: d.service,
		client:      d.client,
	}
}

func (d *Event) Response() *Transaction {
	return &Transaction{
		ActionType:  ActionTypeResponse,
		ServiceName: d.service,
		client:      d.client,
	}
}

func (t *Transaction) WithMetadata(meta EventMetadata) *Transaction {
	t.Metadata = meta
	return t
}

func (t *Transaction) WithTraceId(traceId string) *Transaction {
	t.TraceID = traceId
	return t
}

func (t *Transaction) Info(event string) {
	t.EventType = event
	t.Time = time.Now().UTC()
	_, err := json.Marshal(t)
	if err != nil {
	}
	go t.client.Push(t)
}

func (t *Transaction) Error(event string, err error) {
	t.ErrorCode = strings.Split(err.Error(), ":")[0]
	t.EventType = event
	t.Time = time.Now().UTC()
	go t.client.Push(t)
}
