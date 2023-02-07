package nop

import "github.com/raymondwongso/gogox/log"

// Logger is an implementation of Logger interface with zero ops.
// commonly used for testing.
type Logger struct{}

// New creates new Nop logger.
func New() *Logger {
	return &Logger{}
}

func (l *Logger) Trace(msg string, args ...interface{})                   {}
func (l *Logger) Tracew(msg string, md log.Metadata, args ...interface{}) {}
func (l *Logger) Debug(msg string, args ...interface{})                   {}
func (l *Logger) Debugw(msg string, md log.Metadata, args ...interface{}) {}
func (l *Logger) Info(msg string, args ...interface{})                    {}
func (l *Logger) Infow(msg string, md log.Metadata, args ...interface{})  {}
func (l *Logger) Warn(msg string, args ...interface{})                    {}
func (l *Logger) Warnw(msg string, md log.Metadata, args ...interface{})  {}
func (l *Logger) Error(msg string, args ...interface{})                   {}
func (l *Logger) Errorw(msg string, md log.Metadata, args ...interface{}) {}
func (l *Logger) Fatal(msg string, args ...interface{})                   {}
func (l *Logger) Fatalw(msg string, md log.Metadata, args ...interface{}) {}
func (l *Logger) Panic(msg string, args ...interface{})                   {}
func (l *Logger) Panicw(msg string, md log.Metadata, args ...interface{}) {}

func (l *Logger) Log(level log.LogLevel, msg string, args ...interface{})                   {}
func (l *Logger) Logw(level log.LogLevel, msg string, md log.Metadata, args ...interface{}) {}
