package service

import (
	"github.com/bruchte/libreLogger/internal/core/port"
)

type LoggerLocalService struct {
	loggerPort port.LoggerLocalIF
}

func NewLoggerLocalService(port port.LoggerLocalIF) *LoggerLocalService {
	return &LoggerLocalService{
		loggerPort: port,
	}
}

func (s *LoggerLocalService) Debug(msg ...interface{}) {
	s.loggerPort.Debug(msg...)
}
func (s *LoggerLocalService) Debugf(format string, arg ...interface{}) {
	s.loggerPort.Debugf(format, arg...)
}
func (s *LoggerLocalService) Info(msg ...interface{}) {
	s.loggerPort.Info(msg...)
}
func (s *LoggerLocalService) Infof(format string, arg ...interface{}) {
	s.loggerPort.Infof(format, arg...)
}
func (s *LoggerLocalService) Warn(msg ...interface{}) {
	s.loggerPort.Warn(msg...)
}
func (s *LoggerLocalService) Warnf(format string, arg ...interface{}) {
	s.loggerPort.Warnf(format, arg...)
}
func (s *LoggerLocalService) Error(msg ...interface{}) {
	s.loggerPort.Error(msg...)
}
func (s *LoggerLocalService) Errorf(format string, arg ...interface{}) {
	s.loggerPort.Errorf(format, arg...)
}

func (s *LoggerLocalService) NewLogger(topic string) port.LoggerLocalIF {
	return s.loggerPort.NewLogger(topic)
}
