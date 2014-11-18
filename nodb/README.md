## Nodb cache provider

This provider is based on [Nodb](https://github.com/lunny/nodb).

## Usage

```go
import (
    "github.com/Unknwon/macaron"
    "github.com/macaron-contrib/cache"
    _ "github.com/macaron-contrib/cache/nodb"
)

func main() {
    m := macaron.Classic()
    m.Use(cahce.Cacher(cache.CacheOptions{
        Adapter:    "nodb", // Name of adapter.
        Conn:       "./cachedir", // cache dir
    }))
    
    m.Get("/", func(c cache.Cache) string {
        c.Put("cache", "cache middleware", 120)
        return c.Get("cache")
    })

    m.Run()
}
```