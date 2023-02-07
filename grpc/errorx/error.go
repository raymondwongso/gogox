package errorx

import (
	"github.com/raymondwongso/gogox/errorx"
	"github.com/raymondwongso/gogox/grpclib/protobuf"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/runtime/protoiface"
)

// GrpcError defines GRPC error wrapper for gogox error
type GrpcError struct {
	Code            codes.Code
	UnderlyingError *errorx.Error
}

// Error implements error interface
func (e *GrpcError) Error() string {
	return e.UnderlyingError.Error()
}

// GRPCStatus implements GRPCStatus interface that is required by grpc server handler
func (e *GrpcError) GRPCStatus() *status.Status {
	return status.New(e.Code, e.UnderlyingError.Error())
}

// GRPCStatusWithDetails return status with details containing the gogox error instance
func (e *GrpcError) GRPCStatusWithDetails() (*status.Status, error) {
	grpcStatus := e.GRPCStatus()
	details := make([]protoiface.MessageV1, 0)

	errorPb := protobuf.Error{Code: e.UnderlyingError.Code, Message: e.UnderlyingError.Message}
	for _, detail := range e.UnderlyingError.Details {
		errorPb.Details = append(errorPb.Details, &protobuf.Detail{
			Field:   detail.Field,
			Message: detail.Message,
		})
	}
	details = append(details, &errorPb)

	return grpcStatus.WithDetails(details...)
}
