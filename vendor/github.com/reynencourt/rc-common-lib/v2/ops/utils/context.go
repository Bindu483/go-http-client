package utils

import (
	"sync"

	"github.com/reynencourt/rc-common-lib/v2/events_logger"
)

var once sync.Once
var applicationContext GlobalApplicationContext

type GlobalApplicationContext struct {
	EventsPusher *events_logger.Event
}

func GetEventsObject(sName string) *events_logger.Event {
	once.Do(
		func() {
			c := events_logger.NewHttpEventClient("http://localhost:9000", 10)
			e := events_logger.NewEvent(sName, c)
			applicationContext = GlobalApplicationContext{e}
		})
	return applicationContext.EventsPusher
}

func NewEventsObject(serviceName string) *events_logger.Event {
	c := events_logger.NewHttpEventClient("http://localhost:9000", 10)
	e := events_logger.NewEvent(serviceName, c)
	return e
}
