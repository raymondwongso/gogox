package log

import (
	"context"

	"golang.org/x/exp/maps"
)

type MetadataContextKeyType string

const (
	// MetadataContextKey is a custom context key for metadata
	MetadataContextKey MetadataContextKeyType = "metadata.context.key"
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

	md := ctxVal.(Metadata)
	res := Metadata{}
	maps.Copy(res, md)
	return res
}
