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

	logger *log.Logger
}

func NewLoggerLocalInternal(topic string) *LoggerLocalInternal {

	//config?
	t := "[" + topic + "]"
	for len(t) < 10 {
		t += " "
	}
	l := log.New(os.Stdout, t, 0)
	return &LoggerLocalInternal{
		logger: l,
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
	return NewLoggerLocalInternal(topic)
}

func (s *LoggerLocalInternal) logDirect(level string, msg ...interface{}) {
	msg[0] = level + "|" + getCallerString() + fmt.Sprintf("%s", msg[0])
	s.logger.Println(msg...)
}
func (s *LoggerLocalInternal) logFormatted(level string, format string, arg ...interface{}) {
	s.logger.Printf(level+"|"+getCallerString()+format, arg...)
}
