# cache

`cache` provides `Cache` interface containing commonly used cache usecases. Notably:
1. `Get` get the value of certain cache key.
2. `Set` set the value to certain cache key with expiration.
3. `Del` delete a certain cache key.

## How to Use
There are 2 provided cache adapter: `redis` and `memcache`. Define your object and inject the `Cache` implementor into it. Populate the implementor with your choosen adapter.

```go

import (
  gogox_cache "github.com/raymondwongso/cache"
  gogox_redis "github.com/raymondwongso/cache/redis"

  goredis "github.com/go-redis/redis/v8"
)

type Service struct {
  cache gogox_cache.Cache
}

func (s *Service) DoSomething(ctx context.Context) error {
  // do something
  type someStruct struct {}
  res := someStruct{}

  err := s.cache.Get(ctx, "some-key", &res)
  // do something with res
  return err
}

func main () {
  goredisClient := goredis.NewClient(&goredis.Options{
		Addr: "your-redis-instance-addr",
		Password: "your-redis-instance-pass",
	})

  redisCache := gogox_redis.New(
    goredisCLient,
  )

  service := &Service{cache: redisCache}
  service.DoSomething(context.Background())
}
```

You can easily change the adapter to `memcache` like this:
```go

import (
  "github.com/bradfitz/gomemcache/memcache"
)

func main() {
  memcacheClient := memcache.New("your-memcached-instance-addr:port")

  memcacheCache := gogox_memcache.New(
    memcacheClient,
  )

  service := &Service{cache: memcacheCache}
}
```