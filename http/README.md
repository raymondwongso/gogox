# http

Contains http components to integrate with gogox std lib

## log

`LoggingMiddleware` provides middleware to automatically log your HTTP request and response.

```go
func dummyHandler(w http.ResponseWriter, r *http.Request) {
  w.Write([]byte("OK"))
}

mux := http.NewServeMux()

logrus := logrus.New()
logrus.SetFormatter(&logrus.JSONFormatter{})
logrus.SetLevel(logrus.DebugLevel)
logger := gogox_logrus.New(logrus, nil)

mux.Handle("/", log.LoggingMiddleware(logger, dummyHandler, log.DefaultOptions()))
```

You can provide `options` to control the middlware.

```go
opts := log.DefaultOptions()
// Do not log if url is something
opts.ShouldLog = func(req *http.Request, status int) bool {
  if req.URL == "something" {
    return false
  }
  return true
}

mux.Handle("/", log.LoggingMiddleware(logger, dummyHandler, opts))
```

## trace

`TraceMiddlware` provides middleware to automatically inject traceID for your HTTP request. It also inject the `traceField` into log metadata.

```go
func dummyHandler(w http.ResponseWriter, r *http.Request) {
  w.Write([]byte("OK"))
}

mux := http.NewServeMux()

// your-trace-field is used for log metadata
// X-Trace-ID is used for setting request header
mux.Handle("/", trace.TraceMiddlware("your-trace-field", "X-Trace-ID", dummyHandler))
```

## gateway

Contains grpc gateway components to ingetate with gogox std lib.

`ErrorHandler` defines handler for error response. It will parse the status error into gogox grpc error `errorx.GrpcError`.