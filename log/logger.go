package log

//go:generate mockgen -destination=mock/logger.go -package=logmock -source=logger.go

// LogLevel defines log level
type LogLevel int

const (
	TraceLevel LogLevel = iota
	DebugLevel
	InfoLevel
	WarnLevel
	ErrorLevel
	FatalLevel
	PanicLevel
)

// Logger defines generic interface for all logger.
type Logger interface {
	Trace(msg string, args ...interface{})
	Tracew(msg string, md Metadata, args ...interface{})
	Debug(msg string, args ...interface{})
	Debugw(msg string, md Metadata, args ...interface{})
	Info(msg string, args ...interface{})
	Infow(msg string, md Metadata, args ...interface{})
	Warn(msg string, args ...interface{})
	Warnw(msg string, md Metadata, args ...interface{})
	Error(msg string, args ...interface{})
	Errorw(msg string, md Metadata, args ...interface{})
	Fatal(msg string, args ...interface{})
	Fatalw(msg string, md Metadata, args ...interface{})
	Panic(msg string, args ...interface{})
	Panicw(msg string, md Metadata, args ...interface{})

	Log(level LogLevel, msg string, args ...interface{})
	Logw(level LogLevel, msg string, md Metadata, args ...interface{})
}
