# log

Log provides generic logger that is pluggable

## How to Use

`Logger` interface contains commonly used logging usecases. Inject `Logger` into your object. Use adapters provided (logrus or zap)

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
  logrus.SetLevel(logrus.ErrorLevel)

  logger := gogox_logrus.New(logrus, nil)

  service := service.NewService(logger)
}
```

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
Use `NewContext` to inject your log metadata to context.

```go
ctx := log.NewContext(context.Background(), log.Metadata{"user_id": 123})
```

Use `MetadataFromContext` to extract your log metadata from context

```go
md := log.MetadataFromContext(ctx)
```