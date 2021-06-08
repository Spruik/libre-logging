package port

//The LoggerLocalIF interface defines the logging functions to be provided by any local logging
type LoggerLocalIF interface {
	Debug(msg ...interface{})
	Debugf(format string, arg ...interface{})
	Info(msg ...interface{})
	Infof(format string, arg ...interface{})
	Warn(msg ...interface{})
	Warnf(format string, arg ...interface{})
	Error(msg ...interface{})
	Errorf(format string, arg ...interface{})

	NewLogger(topic string) LoggerLocalIF
}
