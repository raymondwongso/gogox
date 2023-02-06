# stats

Stats provides usecase for scoring your app metrics

## How to Use

`stats.Stats` defines standard stats interface.

```go
type service struct {
  stats stats.Stats
}

func (s *service) DoSomething() {
  // something happened
  if err != nil {
    s.stats.Increment("something_failed", stats.Option{
      Tags: stats.Tags{"method": "DoSomething"}
    })
  }
}
```

You can inject stats implementor.
```go

import (
  gogox_prom "github.com/raymondwongso/gogox/stats/prometheus"
)

func main() {
  // e.g: prometheus stats, you can add base tags for your stats.
  prom := gogox_prom.New("mynamespace", stats.Tags{"service":"gogox_service"})

  service := service.NewService(prom)
}
```
