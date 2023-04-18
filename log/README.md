# log

log provides generic logger interface so your application does not need to know which logger adapter you use. This simplify testing and keep the code clean.

## How to Use

`Logger` interface contains commonly used logging usecases. Inject `Logger` into your object.

```go
// service.go
type service struct {
  logger log.Logger
}

func NewService(logger log.Logger) *service {
  return &service{logger: logger}
}

func (s *service) DoSomething() error {
  // do your thing
  if err != nil {
    s.logger.Error("some error happened")
    return err
  }
}

// main.go
import (
  "github.com/sirupsen/logrus"

	gogox_logrus "github.com/raymondwongso/gogox/log/logrus"
)

func main() {
  logrus := logrus.New()
  logrus.SetFormatter(&logrus.JSONFormatter{})
  logrus.SetLevel(logrus.DebugLevel)

  logger := gogox_logrus.New(logrus, nil)

  service := service.NewService(logger)
}
```

Currently supported adapters:
1. Logrus
2. Zap
3. Zerolog (Will be available in 1.x)
4. Nop

## Metadata

You can add `metadata` into your log entries, commonly used to store additional error info. Use `[severity]w` method.

```go
s.logger.Errorw("some error happened", log.Metadata{"error": err.Error(), "user_id": 123})

// result
{"error":"some error happened","level":"error","msg":"some error happened","service":"api_logrus","time":"2023-02-03T11:32:54+07:00","user_id":123}
```

You can also add `baseMetadata` into your logger instance, commonly used to store global metadata for all your log entries.

```go
logger := gogox_logrus.New(logrus, log.Metadata{"service": "api", "version": "1.2.1"})

s.logger.Errorw("some error happened", log.Metadata{"error": err.Error(), "user_id": 123})

// result
{"error":"some error come up","level":"error","msg":"some error come up","service":"api","time":"2023-02-03T11:35:25+07:00","user_id":1,"version":"1.2.1"}
```

## Context

Context is useful for passing base metadata into all your log entries. Use `NewContext` to inject your log metadata to context.

```go
ctx := log.NewContext(context.Background(), log.Metadata{"user_id": 123})
```

Use `MetadataFromContext` to extract your log metadata from context

```go
md := log.MetadataFromContext(ctx)
```