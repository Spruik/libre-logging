package api

import (
	"errors"
	"fmt"
	"github.com/Spruik/libre-logging/interfaces"
	"github.com/Spruik/libre-logging/internal/core/service"
	"github.com/gorilla/mux"
	"net/http"
	"sort"
)

func VisitAllLoggers(visitor interfaces.LibreLoggerVisitor) {
	for _, logger := range service.LoggerMap {
		logger.AcceptVisitor(visitor)
	}
}

func VisitOneLogger(name string, visitor interfaces.LibreLoggerVisitor) error {
	var err error = nil
	logger := service.LoggerMap[name]
	if logger != nil {
		logger.AcceptVisitor(visitor)
	} else {
		err = errors.New(fmt.Sprintf("No logger named '%s'", name))
	}
	return err
}

type loggerData struct {
	name        string
	topic       string
	level       string
	destination string
}
type loggerDataVisitor struct {
	results map[string]loggerData
}

func (s *loggerDataVisitor) LibreLoggerVisit(logger interfaces.LoggerLocalIF) {
	s.results[logger.GetName()] = loggerData{
		name:        logger.GetName(),
		topic:       logger.GetTopic(),
		level:       logger.GetLevel(),
		destination: logger.GetDestination(),
	}
}

func LoggersLink(w http.ResponseWriter, r *http.Request) {
	_ = r
	var respStr = ""
	visitor := loggerDataVisitor{results: map[string]loggerData{}}
	VisitAllLoggers(&visitor)
	lgrNames := make([]string, 0, 0)
	for lgrName := range visitor.results {
		lgrNames = append(lgrNames, lgrName)
	}
	sort.Strings(lgrNames)
	for _, name := range lgrNames {
		data := visitor.results[name]
		respStr += fmt.Sprintf("%-30s : %-10s : %-10s : %s\n", name, data.topic, data.level, data.destination)
	}
	respStr += "\n\n"
	respStr += "    Use /loggers/{level}          to change all levels\n"
	respStr += "    Use /loggers/{name}/{level}   to change one logger's levels\n"
	_, _ = fmt.Fprintf(w, "%s", respStr)
}

type loggerChangeVisitor struct {
	newLevel string
}

func (s *loggerChangeVisitor) LibreLoggerVisit(logger interfaces.LoggerLocalIF) {
	logger.SetLoggingLevel(s.newLevel)
}

func ChangeOneLoggerLink(w http.ResponseWriter, r *http.Request) {
	name := mux.Vars(r)["name"]
	lvl := mux.Vars(r)["level"]
	visitor := loggerChangeVisitor{newLevel: lvl}
	respStr := ""
	err := VisitOneLogger(name, &visitor)
	if err != nil {
		respStr += err.Error()
	} else {
		respStr += "logging level changed"
	}
	_, _ = fmt.Fprintf(w, "%s", respStr)
}

func ChangeAllLoggersLink(w http.ResponseWriter, r *http.Request) {
	lvl := mux.Vars(r)["level"]
	respStr := ""
	visitor := loggerChangeVisitor{newLevel: lvl}
	VisitAllLoggers(&visitor)
	respStr += "logging level changed"
	_, _ = fmt.Fprintf(w, "%s", respStr)
}
