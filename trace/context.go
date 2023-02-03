package trace

import "context"

type TraceContextKeyType string

const (
	// TraceContextKey is a custom context key for trace
	TraceContextKey TraceContextKeyType = "trace.context.key"
)

// NewContext create new context that is injected with trace
func NewContext(parentCtx context.Context, trace string) context.Context {
	if parentCtx == nil {
		parentCtx = context.Background()
	}

	return context.WithValue(parentCtx, TraceContextKey, trace)
}

// TraceFromContext extract trace from context
func TraceFromContext(ctx context.Context) string {
	if ctx == nil {
		return ""
	}

	ctxVal := ctx.Value(TraceContextKey)
	if ctxVal == nil {
		return ""
	}

	return ctxVal.(string)
}
