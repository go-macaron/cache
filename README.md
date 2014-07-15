cache
=====

Middleware cache is the cache manager of [Macaron](https://github.com/Unknwon/macaron). It can use many cache adapters, including memory, redis, and memcache.

[API Reference](https://gowalker.org/github.com/macaron-contrib/cache)

## Usage

```go
import (
	"github.com/Unknwon/macaron"
	"github.com/macaron-contrib/cache"
)

func main() {
  	m := macaron.Classic()
  	m.Use(cache.Cacher())
	
	m.Get("/", func(c cache.Cache) string {
		c.Put("cache", "cache middleware", 120)
		return c.Get("cache")
	})

	m.Run()
}
```

To use redis or memcache as adapter, you should import their init functions:

```go
import (
	_ "github.com/macaron-contrib/cache/redis"
	_ "github.com/macaron-contrib/cache/memcache"
)
```

## Options

`cache.Cacher` comes with a variety of configuration options:

```go
// ...
m.Use(cahce.Cacher(cache.CacheOptions{
	Adapter:	"memory", // Name of adapter.
	Interval:	60, // GC interval for memory adapter.
	Conn:		"127.0.0.1:11211", // Connection string for non-memory adapter.
}))
// ...
```

## License

This project is under Apache v2 License. See the [LICENSE](LICENSE) file for the full license text.