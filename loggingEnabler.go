package libreLogger

import "github.com/bruchte/libreLogger/internal/core/port"

type LoggingEnabler struct {
	hook   string
	logger port.LoggerLocalIF
}

func (s *LoggingEnabler) SetLoggerConfigHook(hook string) {
	s.hook = hook
	s.logger = GetLogger(hook)
}

func (s *LoggingEnabler) LogDebug(msg ...interface{}) {
	s.logger.Debug(msg...)
}

func (s *LoggingEnabler) LogDebugf(format string, arg ...interface{}) {
	s.logger.Debugf(format, arg...)
}

func (s *LoggingEnabler) LogInfo(msg ...interface{}) {
	s.logger.Info(msg...)
}

func (s *LoggingEnabler) LogInfof(format string, arg ...interface{}) {
	s.logger.Infof(format, arg...)
}

func (s *LoggingEnabler) LogError(msg ...interface{}) {
	s.logger.Error(msg...)
}

func (s *LoggingEnabler) LogErrorf(format string, arg ...interface{}) {
	s.logger.Errorf(format, arg...)
}
