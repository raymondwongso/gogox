package gateway

import (
	"context"
	"fmt"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/raymondwongso/gogox/errorx"
	"github.com/raymondwongso/gogox/grpc/protobuf"
	"github.com/raymondwongso/gogox/log"
	"google.golang.org/grpc/status"
)

// ErrorHandler defines handler for error response
type ErrorHandler struct {
	logger log.Logger
}

func NewErrorHandler(logger log.Logger) *ErrorHandler {
	return &ErrorHandler{logger: logger}
}

// ErrroxProtoErrorHandler convert error proto message into gogox error.
func (h *ErrorHandler) ErrorxProtoErrorHandler(ctx context.Context, mux *runtime.ServeMux, marshaler runtime.Marshaler, w http.ResponseWriter, r *http.Request, err error) {
	s, ok := status.FromError(err)

	logMd := log.Metadata{
		"url":   r.URL.Path,
		"error": err.Error(),
	}

	if !ok {
		h.logger.Error("error parsing status grpc", logMd)
		runtime.DefaultHTTPError(ctx, mux, marshaler, w, r, err)
		return
	}

	sdetails := s.Proto().GetDetails()
	if len(sdetails) == 0 {
		h.logger.Error("error details not found", logMd)
		runtime.DefaultHTTPError(ctx, mux, marshaler, w, r, err)
		return
	}

	// sdetails will only contain 1 gogox error
	for _, detail := range sdetails {
		fmt.Printf("detail type url: %s\n", detail.TypeUrl)

		pbErr := protobuf.Error{}
		unmarshalErr := detail.UnmarshalTo(&pbErr)
		if unmarshalErr == nil {
			gogoxErr := &errorx.Error{
				Code:    pbErr.Code,
				Message: pbErr.Message,
				Details: []*errorx.Details{},
			}

			for _, d := range pbErr.Details {
				gogoxErr.Details = append(gogoxErr.Details, &errorx.Details{
					Field:   d.Field,
					Message: d.Message,
				})
			}

			buf, marshalErr := marshaler.Marshal(gogoxErr)
			if marshalErr != nil {
				h.logger.Error("error marshal gogox error", logMd)
				runtime.DefaultHTTPError(ctx, mux, marshaler, w, r, err)
				return
			}

			w.WriteHeader(runtime.HTTPStatusFromCode(s.Code()))
			_, _ = w.Write(buf)
			return
		}
	}

	runtime.DefaultHTTPError(ctx, mux, marshaler, w, r, err)
}
