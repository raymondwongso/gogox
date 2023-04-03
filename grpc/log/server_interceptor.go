package log

import (
	"context"
	"fmt"
	"path"
	"time"

	"github.com/raymondwongso/gogox/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var (
	levelMapping = map[codes.Code]log.LogLevel{
		codes.OK:                 log.DebugLevel,
		codes.Canceled:           log.DebugLevel,
		codes.Unknown:            log.InfoLevel,
		codes.InvalidArgument:    log.DebugLevel,
		codes.DeadlineExceeded:   log.InfoLevel,
		codes.NotFound:           log.DebugLevel,
		codes.AlreadyExists:      log.DebugLevel,
		codes.PermissionDenied:   log.InfoLevel,
		codes.Unauthenticated:    log.InfoLevel,
		codes.ResourceExhausted:  log.DebugLevel,
		codes.FailedPrecondition: log.DebugLevel,
		codes.Aborted:            log.DebugLevel,
		codes.OutOfRange:         log.DebugLevel,
		codes.Unimplemented:      log.WarnLevel,
		codes.Internal:           log.WarnLevel,
		codes.Unavailable:        log.WarnLevel,
		codes.DataLoss:           log.WarnLevel,
	}
)

// UnaryServerInterceptor intercepts a GRPC server response and inject
func UnaryServerInterceptor(logger log.Logger, opts options) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		startTime := time.Now()
		logMd := newLogMetadata(ctx, info, startTime)

		resp, err = handler(ctx, req)

		if !opts.ShouldLog(info, err) {
			return resp, err
		}

		code := status.Code(err)
		level := levelMapping[code]
		duration_ms := float32(time.Since(startTime).Nanoseconds()/1000) / 1000

		logMd["grpc.request"] = req
		logMd["grpc.code"] = code
		logMd["grpc.time_ms"] = duration_ms

		if err != nil {
			logMd["error"] = err.Error()
		}

		fullMethod := info.FullMethod
		service := path.Dir(fullMethod)[1:]
		method := path.Base(fullMethod)
		logger.Logw(level, fmt.Sprintf("[%s].%s finished with code: %s", service, method, code), logMd)

		return resp, err
	}
}

func newLogMetadata(ctx context.Context, info *grpc.UnaryServerInfo, startTime time.Time) log.Metadata {
	fullMethod := info.FullMethod
	service := path.Dir(fullMethod)[1:]
	method := path.Base(fullMethod)

	logMd := log.MetadataFromContext(ctx)

	logMd["system"] = "grpc"
	logMd["span.kind"] = "server"
	logMd["grpc.service"] = service
	logMd["grpc.method"] = method
	logMd["grpc.start_time"] = startTime.Format(time.RFC3339)

	if d, ok := ctx.Deadline(); ok {
		logMd["grpc.request.deadline"] = d.Format(time.RFC3339)
	}

	return logMd
}
