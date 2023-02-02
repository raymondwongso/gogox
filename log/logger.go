package log

import "golang.org/x/exp/maps"

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
}

// Metadata defines metadata for logger, which will be included whenever log is written.
type Metadata map[string]interface{}

// MergeMetadata creates new metadata as a result from merging between md1 and md2.
func MergeMetadata(md1, md2 Metadata) Metadata {
	res := Metadata{}
	maps.Copy(res, md1)
	maps.Copy(res, md2)
	return res
}
