package interfaces

//The LoggerLocalIF interfaces defines the logging functions to be provided by any local logging
type LoggerLocalIF interface {
	Debug(msg ...interface{})
	Debugf(format string, arg ...interface{})
	Info(msg ...interface{})
	Infof(format string, arg ...interface{})
	Warn(msg ...interface{})
	Warnf(format string, arg ...interface{})
	Error(msg ...interface{})
	Errorf(format string, arg ...interface{})

	//NewLogger(topic string) LoggerLocalIF
	SetLoggingLevel(level string)

	AcceptVisitor(visitor LibreLoggerVisitor)
	GetLevel() string
	GetDestination() string
	GetName() string
	GetTopic() string
}

type LibreLoggerVisitor interface {
	LibreLoggerVisit(logger LoggerLocalIF)
}
