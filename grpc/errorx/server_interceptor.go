package errorx

import (
	"context"

	"github.com/raymondwongso/gogox/errorx"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
)

// UnaryServerInterceptor intercepts a GRPC server response and wrap the error proto
func UnaryServerInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		res, err := handler(ctx, req)

		if err != nil {
			grpcErr, ok := err.(*GrpcError)

			if !ok {
				grpcErr = &GrpcError{
					Code:            codes.Internal,
					UnderlyingError: errorx.New(errorx.CodeInternal, err.Error()),
				}
			}

			status, err := grpcErr.GRPCStatusWithDetails()
			if err != nil {
				return res, err
			}

			return res, status.Err()
		}

		return res, err
	}
}
