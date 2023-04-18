# trace

Trace provides usecase for generating trace id, commonly used for tracing your request throughout its lifecycle.

## How to Use

`trace.New` creates new random trace string.

## Context

Use `NewContext` to inject your trace to context.

```go
ctx := trace.NewContext(context.Background(), "your-trace-id")
```

Use `TraceFromContext` to extract your trace from context

```go
trace := trace.TraceFromContext(ctx)
fmt.Println(trace)
// your-trace-id
```
