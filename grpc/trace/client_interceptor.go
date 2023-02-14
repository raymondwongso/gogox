package trace

import (
	"context"

	"github.com/raymondwongso/gogox/log"
	"github.com/raymondwongso/gogox/trace"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

// UnaryClientInterceptor intercepts a GRPC client request and inject the trace ID.
// traceField is used to inject traceID to log
func UnaryClientInterceptor(traceField, traceHeaderKey string) grpc.UnaryClientInterceptor {
	return func(ctx context.Context, method string, req interface{}, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
		// get traceID from request header, if not found check request context.
		// otherwise, generate new trace
		traceID := valueFromMetadata(ctx, traceHeaderKey, metadata.FromOutgoingContext)
		if traceID == "" {
			traceID = trace.TraceFromContext(ctx)
			if traceID == "" {
				traceID = trace.New()
			}
		}

		ctx = metadata.AppendToOutgoingContext(ctx, traceField, traceID)

		ctx = trace.NewContext(ctx, traceID)

		logMd := log.MetadataFromContext(ctx)
		logMd[traceField] = traceID

		ctx = log.NewContext(ctx, logMd)

		return invoker(ctx, method, req, reply, cc, opts...)
	}
}
