package libreLogger

import (
	"github.com/bruchte/libreConfig"
	"github.com/bruchte/libreLogger/internal/core/service"
	"github.com/bruchte/libreLogger/internal/implementation"
	"log"
)

var loggerMap = map[string]*service.LoggerLocalService{}

func Initialize(configComponentName string) {
	loggersStanza, err := libreConfig.GetConfigService().GetConfigStanza(configComponentName, "loggers")
	if err != nil {
		panic("Error looking for 'loggers' section in component configuration")
	}
	for _, listEntry := range loggersStanza.Children {
		//each item here is an anonymous list stanza, so we go next level down of the logger config
		loggerStanza := listEntry.Children[0]
		loggerName := loggerStanza.Name
		var topic string
		//TODO - other config attrs
		for _, item := range loggerStanza.Children {
			switch item.Name {
			case "topic":
				topic = item.Value
				//TODO - other config attrs
			}
		}
		t := implementation.NewLoggerLocalInternal(topic)
		loggerMap[loggerName] = service.NewLoggerLocalService(t)
		log.Printf("Built logger %s with topic %s", loggerName, topic)
	}
	topic := "********"
	t := implementation.NewLoggerLocalInternal(topic)
	loggerMap[topic] = service.NewLoggerLocalService(t)
	log.Printf("Built default logger %s with topic %s", topic, topic)
}

func GetLogger(name string) *service.LoggerLocalService {
	ret := loggerMap[name]
	if ret == nil {
		ret = loggerMap["********"]
	}
	return ret
}
