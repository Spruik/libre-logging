package service

import "github.com/Spruik/libre-logging/interfaces"

var LoggerMap = map[string]*LoggerLocalService{}

type LoggerLocalService struct {
	loggerinterfaces interfaces.LoggerLocalIF
}

func NewLoggerLocalService(interfaces interfaces.LoggerLocalIF) *LoggerLocalService {
	return &LoggerLocalService{
		loggerinterfaces: interfaces,
	}
}

func (s *LoggerLocalService) Debug(msg ...interface{}) {
	s.loggerinterfaces.Debug(msg...)
}
func (s *LoggerLocalService) Debugf(format string, arg ...interface{}) {
	s.loggerinterfaces.Debugf(format, arg...)
}
func (s *LoggerLocalService) Info(msg ...interface{}) {
	s.loggerinterfaces.Info(msg...)
}
func (s *LoggerLocalService) Infof(format string, arg ...interface{}) {
	s.loggerinterfaces.Infof(format, arg...)
}
func (s *LoggerLocalService) Warn(msg ...interface{}) {
	s.loggerinterfaces.Warn(msg...)
}
func (s *LoggerLocalService) Warnf(format string, arg ...interface{}) {
	s.loggerinterfaces.Warnf(format, arg...)
}
func (s *LoggerLocalService) Error(msg ...interface{}) {
	s.loggerinterfaces.Error(msg...)
}
func (s *LoggerLocalService) Errorf(format string, arg ...interface{}) {
	s.loggerinterfaces.Errorf(format, arg...)
}

//func (s *LoggerLocalService) NewLogger(topic string) interfaces2.LoggerLocalIF {
//	return s.loggerinterfaces.NewLogger(topic)
//}
func (s *LoggerLocalService) SetLoggingLevel(level string) {
	s.loggerinterfaces.SetLoggingLevel(level)
}
func (s *LoggerLocalService) AcceptVisitor(visitor interfaces.LibreLoggerVisitor) {
	s.loggerinterfaces.AcceptVisitor(visitor)
}

func (s *LoggerLocalService) GetLevel() string {
	return s.loggerinterfaces.GetLevel()
}
func (s *LoggerLocalService) GetDestination() string {
	return s.loggerinterfaces.GetDestination()
}
func (s *LoggerLocalService) GetName() string {
	return s.loggerinterfaces.GetName()
}
func (s *LoggerLocalService) GetTopic() string {
	return s.loggerinterfaces.GetTopic()
}
