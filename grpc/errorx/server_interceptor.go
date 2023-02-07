package errorx

import (
	"context"

	"google.golang.org/grpc"
)

// UnaryServerInterceptor intercepts a GRPC server response and wrap the error proto
func UnaryServerInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		res, err := handler(ctx, req)

		if err != nil {
			grpcErr, ok := err.(*GrpcError)
			if ok {
				status, err := grpcErr.GRPCStatusWithDetails()
				if err != nil {
					return res, err
				}

				return res, status.Err()
			}
		}

		return res, err
	}
}
