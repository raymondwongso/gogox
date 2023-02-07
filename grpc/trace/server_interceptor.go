package trace

import (
	"context"

	"github.com/raymondwongso/gogox/log"
	"github.com/raymondwongso/gogox/trace"
	"google.golang.org/grpc"
)

// UnaryServerInterceptor intercepts a GRPC server response and inject the trace ID.
func UnaryServerInterceptor(traceField string) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		traceID := trace.TraceFromContext(ctx)
		if traceID == "" {
			traceID = trace.New()
		}

		ctx = trace.NewContext(ctx, trace.New())

		logMd := log.MetadataFromContext(ctx)
		logMd[traceField] = traceID

		ctx = log.NewContext(ctx, logMd)

		return handler(ctx, req)
	}
}
