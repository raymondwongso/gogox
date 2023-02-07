# grpc

Contains various grpc components to help intergrate with gogox std lib.

## errorx

`GrpcError` is a wrapper for gogox's `Error`, to help parsing the error into grpc status.

## log

Provides interceptor to log the response.

## protobuf

Contains proto message for `GrpcError`

## trace

Provides interceptor to inject trace to each request.
