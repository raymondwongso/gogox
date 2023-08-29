package gateway

import (
	"context"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/raymondwongso/gogox/errorx"
	grpc_errorx "github.com/raymondwongso/gogox/grpc/errorx"
	"github.com/raymondwongso/gogox/grpc/protobuf"
	"github.com/raymondwongso/gogox/log"
	"google.golang.org/grpc/codes"
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
	logMd := log.Metadata{
		"url":   r.URL.Path,
		"error": err.Error(),
	}

	response, httpStatus := h.httpResponse(ctx, err, r, logMd)

	buf, marshalErr := marshaler.Marshal(response)
	if marshalErr != nil {
		h.logger.Errorw("error marshal gogox error", logMd)
		response = &grpc_errorx.GrpcError{
			Code:            codes.Internal,
			UnderlyingError: errorx.New(errorx.CodeInternal, marshalErr.Error()),
		}
		httpStatus = http.StatusInternalServerError
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(httpStatus)
	_, _ = w.Write(buf)
}

func (h *ErrorHandler) httpResponse(ctx context.Context, err error, r *http.Request, logMd log.Metadata) (interface{}, int) {
	s, ok := status.FromError(err)

	defaultErr := &errorx.Error{
		Code:    errorx.CodeInternal,
		Message: s.Message(),
		Details: []*errorx.Details{},
	}

	if !ok {
		h.logger.Errorw("error parsing status grpc", logMd)
		return defaultErr, http.StatusInternalServerError
	}

	httpStatus := runtime.HTTPStatusFromCode(s.Code())
	defaultCode, ok := errorx.DefaultCodeMap[s.Code()]
	if ok {
		defaultErr.Code = defaultCode
	}

	sdetails := s.Proto().GetDetails()
	if len(sdetails) == 0 {
		h.logger.Errorw("error details not found", logMd)
		return defaultErr, httpStatus
	}

	for _, detail := range sdetails {
		if detail.TypeUrl != "type.googleapis.com/raymondwongso.gogox.grpc.Error" {
			continue
		}

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

			return gogoxErr, httpStatus
		}
	}

	return defaultErr, httpStatus
}
