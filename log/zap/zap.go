package zap

import (
	"github.com/raymondwongso/gogox/log"
	"go.uber.org/zap"
)

// Logger is a wrapper of zap.Logger that implements Logger interface.
// baseMd will be included in each log call.
type Logger struct {
	zap    *zap.Logger
	baseMd log.Metadata
}

// New creates new Logger instance with base metadata.
func New(zap *zap.Logger, md log.Metadata) *Logger {
	return &Logger{zap: zap, baseMd: md}
}

func (l *Logger) Trace(msg string, args ...interface{}) {
	// do nothing as zap does not support trace level
}

func (l *Logger) Tracew(msg string, md log.Metadata, args ...interface{}) {
	// do nothing as zap does not support trace level
}

func (l *Logger) Debug(msg string, args ...interface{}) {
	if l.zap == nil {
		return
	}

	l.zap.Log(zap.DebugLevel, msg, buildZapFields(l.baseMd)...)
}

func (l *Logger) Debugw(msg string, md log.Metadata, args ...interface{}) {
	if l.zap == nil {
		return
	}

	l.zap.Log(zap.DebugLevel, msg, buildZapFields(log.MergeMetadata(l.baseMd, md))...)
}

func (l *Logger) Info(msg string, args ...interface{}) {
	if l.zap == nil {
		return
	}

	l.zap.Log(zap.InfoLevel, msg, buildZapFields(l.baseMd)...)
}

func (l *Logger) Infow(msg string, md log.Metadata, args ...interface{}) {
	if l.zap == nil {
		return
	}

	l.zap.Log(zap.InfoLevel, msg, buildZapFields(log.MergeMetadata(l.baseMd, md))...)
}

func (l *Logger) Warn(msg string, args ...interface{}) {
	if l.zap == nil {
		return
	}

	l.zap.Log(zap.WarnLevel, msg, buildZapFields(l.baseMd)...)
}

func (l *Logger) Warnw(msg string, md log.Metadata, args ...interface{}) {
	if l.zap == nil {
		return
	}

	l.zap.Log(zap.WarnLevel, msg, buildZapFields(log.MergeMetadata(l.baseMd, md))...)
}

func (l *Logger) Error(msg string, args ...interface{}) {
	if l.zap == nil {
		return
	}

	l.zap.Log(zap.ErrorLevel, msg, buildZapFields(l.baseMd)...)
}

func (l *Logger) Errorw(msg string, md log.Metadata, args ...interface{}) {
	if l.zap == nil {
		return
	}

	l.zap.Log(zap.ErrorLevel, msg, buildZapFields(log.MergeMetadata(l.baseMd, md))...)
}

func (l *Logger) Fatal(msg string, args ...interface{}) {
	if l.zap == nil {
		return
	}

	l.zap.Log(zap.FatalLevel, msg, buildZapFields(l.baseMd)...)
}

func (l *Logger) Fatalw(msg string, md log.Metadata, args ...interface{}) {
	if l.zap == nil {
		return
	}

	l.zap.Log(zap.FatalLevel, msg, buildZapFields(log.MergeMetadata(l.baseMd, md))...)
}

func (l *Logger) Panic(msg string, args ...interface{}) {
	if l.zap == nil {
		return
	}

	l.zap.Log(zap.PanicLevel, msg, buildZapFields(l.baseMd)...)
}

func (l *Logger) Panicw(msg string, md log.Metadata, args ...interface{}) {
	if l.zap == nil {
		return
	}

	l.zap.Log(zap.PanicLevel, msg, buildZapFields(log.MergeMetadata(l.baseMd, md))...)
}

func buildZapFields(md log.Metadata) (res []zap.Field) {
	for k, v := range md {
		res = append(res, zap.Any(k, v))
	}
	return res
}
