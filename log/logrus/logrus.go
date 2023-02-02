package logrus

import (
	"github.com/raymondwongso/gogox/log"
	"github.com/sirupsen/logrus"
)

// Logger is a wrapper of logrus.Logger that implements Logger interface.
// baseMd will be included in each log call.
type Logger struct {
	logrus *logrus.Logger
	baseMd log.Metadata
}

// New creates new Logger instance with base metadata.
func New(logrus *logrus.Logger, md log.Metadata) *Logger {
	return &Logger{logrus: logrus, baseMd: md}
}

func (l *Logger) Trace(msg string, args ...interface{}) {
	if l.logrus == nil {
		return
	}

	l.logrus.WithFields(logrus.Fields(l.baseMd)).Logf(logrus.TraceLevel, msg, args...)
}

func (l *Logger) Tracew(msg string, md log.Metadata, args ...interface{}) {
	if l.logrus == nil {
		return
	}

	l.logrus.WithFields(logrus.Fields(
		log.MergeMetadata(l.baseMd, md),
	)).Logf(logrus.TraceLevel, msg, args...)
}

func (l *Logger) Debug(msg string, args ...interface{}) {
	if l.logrus == nil {
		return
	}

	l.logrus.WithFields(logrus.Fields(l.baseMd)).Logf(logrus.DebugLevel, msg, args...)
}

func (l *Logger) Debugw(msg string, md log.Metadata, args ...interface{}) {
	if l.logrus == nil {
		return
	}

	l.logrus.WithFields(logrus.Fields(
		log.MergeMetadata(l.baseMd, md),
	)).Logf(logrus.DebugLevel, msg, args...)
}

func (l *Logger) Info(msg string, args ...interface{}) {
	if l.logrus == nil {
		return
	}

	l.logrus.WithFields(logrus.Fields(l.baseMd)).Logf(logrus.InfoLevel, msg, args...)
}

func (l *Logger) Infow(msg string, md log.Metadata, args ...interface{}) {
	if l.logrus == nil {
		return
	}

	l.logrus.WithFields(logrus.Fields(
		log.MergeMetadata(l.baseMd, md),
	)).Logf(logrus.InfoLevel, msg, args...)
}

func (l *Logger) Warn(msg string, args ...interface{}) {
	if l.logrus == nil {
		return
	}

	l.logrus.WithFields(logrus.Fields(l.baseMd)).Logf(logrus.WarnLevel, msg, args...)
}

func (l *Logger) Warnw(msg string, md log.Metadata, args ...interface{}) {
	if l.logrus == nil {
		return
	}

	l.logrus.WithFields(logrus.Fields(
		log.MergeMetadata(l.baseMd, md),
	)).Logf(logrus.WarnLevel, msg, args...)
}

func (l *Logger) Error(msg string, args ...interface{}) {
	if l.logrus == nil {
		return
	}

	l.logrus.WithFields(logrus.Fields(l.baseMd)).Logf(logrus.ErrorLevel, msg, args...)
}

func (l *Logger) Errorw(msg string, md log.Metadata, args ...interface{}) {
	if l.logrus == nil {
		return
	}

	l.logrus.WithFields(logrus.Fields(
		log.MergeMetadata(l.baseMd, md),
	)).Logf(logrus.ErrorLevel, msg, args...)
}

func (l *Logger) Fatal(msg string, args ...interface{}) {
	if l.logrus == nil {
		return
	}

	l.logrus.WithFields(logrus.Fields(l.baseMd)).Logf(logrus.FatalLevel, msg, args...)
}

func (l *Logger) Fatalw(msg string, md log.Metadata, args ...interface{}) {
	if l.logrus == nil {
		return
	}

	l.logrus.WithFields(logrus.Fields(
		log.MergeMetadata(l.baseMd, md),
	)).Logf(logrus.FatalLevel, msg, args...)
}

func (l *Logger) Panic(msg string, args ...interface{}) {
	if l.logrus == nil {
		return
	}

	l.logrus.WithFields(logrus.Fields(l.baseMd)).Logf(logrus.PanicLevel, msg, args...)
}

func (l *Logger) Panicw(msg string, md log.Metadata, args ...interface{}) {
	if l.logrus == nil {
		return
	}

	l.logrus.WithFields(logrus.Fields(
		log.MergeMetadata(l.baseMd, md),
	)).Logf(logrus.PanicLevel, msg, args...)
}
