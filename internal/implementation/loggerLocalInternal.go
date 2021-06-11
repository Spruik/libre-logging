package implementation

import (
	"fmt"
	"github.com/bruchte/libreConfig"
	"github.com/bruchte/libreLogger/internal/core/port"
	"log"
	"os"
	"runtime"
	"strings"
)

type LoggerLocalInternal struct {
	//inherit config
	libreConfig.ConfigurationEnabler
	destination    string
	logger         *log.Logger
	level          string
	levelHierarchy map[string]int
}

func NewLoggerLocalInternal(topic string, dest string) *LoggerLocalInternal {
	var err error
	//config?
	t := "[" + topic + "]"
	for len(t) < 10 {
		t += " "
	}
	var d = os.Stdout
	if dest == "CONSOLE" {
		d = os.Stdout
	} else if strings.Index(dest, "FILE:") == 0 {
		d, err = os.Create(dest[5:])
		if err != nil {
			log.Printf("ERROR:  Failed to open logger destination file %s [%s]", dest[5:], err)
			d = os.Stdout
		}
	}
	l := log.New(d, t, log.Ltime|log.Lmicroseconds)
	lh := map[string]int{"ERROR": 10, "WARN": 20, "INFO": 30, "DEBUG": 40}
	return &LoggerLocalInternal{
		logger:         l,
		destination:    dest,
		levelHierarchy: lh,
	}
}

func getCallerString() string {
	_, file, line, ok := runtime.Caller(5)
	if ok {
		file = file[strings.LastIndex(file, "/")+1:]
		return fmt.Sprintf("(%s:%d)", file, line)
	}
	return "(no call info avaialble)"
}

func (s *LoggerLocalInternal) Debug(msg ...interface{}) {
	s.logDirect("DEBUG", msg...)
}

func (s *LoggerLocalInternal) Debugf(format string, arg ...interface{}) {
	s.logFormatted("DEBUG", format, arg...)
}

func (s *LoggerLocalInternal) Info(msg ...interface{}) {
	s.logDirect("INFO", msg...)
}

func (s *LoggerLocalInternal) Infof(format string, arg ...interface{}) {
	s.logFormatted("INFO", format, arg...)
}

func (s *LoggerLocalInternal) Warn(msg ...interface{}) {
	s.logDirect("WARN", msg...)
}

func (s *LoggerLocalInternal) Warnf(format string, arg ...interface{}) {
	s.logFormatted("WARN", format, arg...)
}

func (s *LoggerLocalInternal) Error(msg ...interface{}) {
	s.logDirect("ERROR", msg...)
}

func (s *LoggerLocalInternal) Errorf(format string, arg ...interface{}) {
	s.logFormatted("ERROR", format, arg...)
}

func (s *LoggerLocalInternal) NewLogger(topic string) port.LoggerLocalIF {
	return NewLoggerLocalInternal(topic, s.destination)
}

func (s *LoggerLocalInternal) SetLoggingLevel(level string) {
	s.level = strings.ToUpper(level)
}

func (s *LoggerLocalInternal) logDirect(level string, msg ...interface{}) {
	if s.shouldLog(level) {
		msg[0] = level + "|" + getCallerString() + fmt.Sprintf("%s", msg[0])
		s.logger.Println(msg...)
	}
}
func (s *LoggerLocalInternal) logFormatted(level string, format string, arg ...interface{}) {
	if s.shouldLog(level) {
		s.logger.Printf(level+"|"+getCallerString()+format, arg...)
	}
}

func (s *LoggerLocalInternal) shouldLog(level string) bool {
	return s.levelHierarchy[level] <= s.levelHierarchy[s.level]
}
