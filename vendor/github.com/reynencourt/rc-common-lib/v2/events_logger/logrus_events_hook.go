package events_logger

import (
	"context"
	"encoding/json"

	dc "github.com/reynencourt/rc-common-lib/v2/proto/data_collector"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/balancer/roundrobin"
)

type RCEventsHook struct {
	clientConn *grpc.ClientConn
}

func NewRCEventsHook(address string) (*RCEventsHook, error) {
	rmConn, err := grpc.Dial(
		address,
		grpc.WithInsecure(),
		grpc.WithBalancerName(roundrobin.Name),
	)
	if err != nil {
		logrus.WithError(err).Error("Something went wrong while calling DC grpc server")
		return nil, err
	}
	logrus.Info("Connection made")
	return &RCEventsHook{clientConn: rmConn}, nil
}

func (hook *RCEventsHook) Fire(entry *logrus.Entry) error {
	data := entry.Data
	eventsRequest := dc.EventsRequests{}
	byteData, err := json.Marshal(data)
	if err != nil {
		logrus.WithError(err).Error("Failed to marshal the data	")
		return err
	}
	err = json.Unmarshal(byteData, &eventsRequest)
	if err != nil {
		logrus.WithError(err).Error("Failed to unmarshal the eventsRequest")
		return err
	}

	switch entry.Level {
	case logrus.PanicLevel:
		logrus.Info("Skipping the panic event")

	case logrus.FatalLevel:
		logrus.Info("Skipping the FatalLevel event")

	case logrus.ErrorLevel:
		dcClient := dc.NewDataCollectorServiceClient(hook.clientConn)
		eventResponse, err := dcClient.SendEvents(context.Background(), &eventsRequest)
		if err != nil {
			logrus.WithError(err).Error("Error occured while SendEvents()")
		}
		logrus.Info(eventResponse)

	case logrus.InfoLevel:
		dcClient := dc.NewDataCollectorServiceClient(hook.clientConn)
		eventResponse, err := dcClient.SendEvents(context.Background(), &eventsRequest)
		if err != nil {
			logrus.WithError(err).Error("Error occured while SendEvents()")
			return err
		}
		logrus.Info(eventResponse)

	case logrus.WarnLevel:
		logrus.Info("Skipping the WarnLevel event")

	case logrus.DebugLevel, logrus.TraceLevel:
		logrus.Info("Skipping the DebugLevel, TraceLevel event")

	default:
		return nil
	}
	return nil
}

func (hook *RCEventsHook) Levels() []logrus.Level {
	return logrus.AllLevels
}
