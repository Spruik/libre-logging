package libreLogger

import (
	"github.com/Spruik/libre-configuration"
	"github.com/Spruik/libre-configuration/shared"
	"github.com/Spruik/libre-logging/api"
	"github.com/Spruik/libre-logging/internal/core/service"
	"github.com/Spruik/libre-logging/internal/implementation"
	"log"
	"net/http"
)

func Initialize(configComponentName string) {
	var err error
	var defDest string
	defDest, err = libreConfig.GetConfigService().GetConfigEntryWithDefault(configComponentName, "defaultDestination", "CONSOLE")
	var defLevel string
	defLevel, err = libreConfig.GetConfigService().GetConfigEntryWithDefault(configComponentName, "defaultLevel", "DEBUG")
	var loggersStanza *shared.ConfigItem
	loggersStanza, err = libreConfig.GetConfigService().GetConfigStanza(configComponentName, "loggers")
	if err != nil {
		panic("Error looking for 'loggers' section in component configuration")
	}
	var topic string
	var dest = defDest
	var level = defLevel
	for _, listEntry := range loggersStanza.Children {
		//each item here is an anonymous list stanza, so we go next level down of the logger config
		loggerStanza := listEntry.Children[0]
		loggerName := loggerStanza.Name
		topic = loggerName //default in case no topic given
		for _, item := range loggerStanza.Children {
			switch item.Name {
			case "topic":
				topic = item.Value
			case "destination":
				dest = item.Value
			case "level":
				level = item.Value
				//TODO - other config attrs
			}
		}
		t := implementation.NewLoggerLocalInternal(loggerName, topic, level, dest)
		service.LoggerMap[loggerName] = service.NewLoggerLocalService(t)
		log.Printf("Built logger %s with topic %s, level %s, and destination %s", loggerName, topic, level, dest)
	}
	topic = "********"
	t := implementation.NewLoggerLocalInternal(topic, topic, defLevel, dest)
	service.LoggerMap[topic] = service.NewLoggerLocalService(t)
	log.Printf("Built default logger %s with topic %s", topic, topic)
}

func GetLogger(name string) *service.LoggerLocalService {
	ret := service.LoggerMap[name]
	if ret == nil {
		ret = service.LoggerMap["********"]
		ret.Warn("###################  MISSING CONFIGURATION ###################")
		ret.Warnf("### Request for logger named '%s' which is not configured", name)
		ret.Warnf("###   consider adding: {\"%s\": {\"topic\":\"<topic>\", \"level\": \"<level>\"}} to the loggers", name)
		ret.Warn("##############################################################")
	}
	return ret
}

func GetRESTAPIEntryPoints() map[string]func(w http.ResponseWriter, r *http.Request) {
	return map[string]func(w http.ResponseWriter, r *http.Request){
		"/loggers":                api.LoggersLink,
		"/loggers/{level}":        api.ChangeAllLoggersLink,
		"/loggers/{name}/{level}": api.ChangeOneLoggerLink,
	}
}
