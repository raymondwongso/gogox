package trace

import (
	"context"

	"github.com/raymondwongso/gogox/log"
	"github.com/raymondwongso/gogox/trace"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

// UnaryServerInterceptor intercepts a GRPC server response and inject the trace ID.
// traceField is used to inject traceID to log
// traceHeaderKey is metadata header key for traceID
func UnaryServerInterceptor(traceField, traceHeaderKey string) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		// get traceID from request header, if not found check request context.
		// otherwise, generate new trace
		traceID := valueFromMetadata(ctx, traceHeaderKey)
		if traceID == "" {
			traceID = trace.TraceFromContext(ctx)
			if traceID == "" {
				traceID = trace.New()
			}
		}

		ctx = trace.NewContext(ctx, traceID)

		logMd := log.MetadataFromContext(ctx)
		logMd[traceField] = traceID

		ctx = log.NewContext(ctx, logMd)

		return handler(ctx, req)
	}
}

func valueFromMetadata(ctx context.Context, key string) string {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return ""
	}

	ids, ok := md[key]
	if !ok || len(ids) == 0 {
		return ""
	}

	return ids[0]
}
