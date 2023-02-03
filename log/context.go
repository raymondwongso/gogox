package log

import "context"

type MetadataContextKeyType string

const (
	// MetadataContextKey is a custom context key
	MetadataContextKey MetadataContextKeyType = "log.metadata.key"
)

// NewContext create new context that is injected with log.Metadata
func NewContext(parentCtx context.Context, md Metadata) context.Context {
	if parentCtx == nil {
		parentCtx = context.Background()
	}

	if md == nil {
		md = Metadata{}
	}

	return context.WithValue(parentCtx, MetadataContextKey, md)
}

// MetadataFromContext extract log.Metadata from context
func MetadataFromContext(ctx context.Context) Metadata {
	if ctx == nil {
		return Metadata{}
	}

	ctxVal := ctx.Value(MetadataContextKey)
	if ctxVal == nil {
		return Metadata{}
	}

	return ctxVal.(Metadata)
}
