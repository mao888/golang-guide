# Gin源码阅读与分析

> 很典型的一个web框架

先看简单的demo：

```go
package main

import "github.com/gin-gonic/gin"

func main() {
        r := gin.Default()
        r.GET("/ping", func(c *gin.Context) {
                c.JSON(200, gin.H{
                        "message": "pong",
                })
        })
        r.Run() // listen and serve on 0.0.0.0:8080
}
```

- 先看 `gin.Default`:

```go
// Default returns an Engine instance with the Logger and Recovery middleware already attached.
func Default() *Engine {
        debugPrintWARNINGDefault()
        engine := New()
        engine.Use(Logger(), Recovery())
        return engine
}
```

- 看 `engine := New()` 所返回的结构体：

```go
func New() *Engine {
        debugPrintWARNINGNew()
        engine := &Engine{
                RouterGroup: RouterGroup{
                        Handlers: nil,
                        basePath: "/",
                        root:     true,
                },
                FuncMap:                template.FuncMap{},
                RedirectTrailingSlash:  true,
                RedirectFixedPath:      false,
                HandleMethodNotAllowed: false,
                ForwardedByClientIP:    true,
                AppEngine:              defaultAppEngine,
                UseRawPath:             false,
                UnescapePathValues:     true,
                MaxMultipartMemory:     defaultMultipartMemory,
                trees:                  make(methodTrees, 0, 9),
                delims:                 render.Delims{Left: "{{", Right: "}}"},
                secureJsonPrefix:       "while(1);",
        }
        engine.RouterGroup.engine = engine
        engine.pool.New = func() interface{} {
                return engine.allocateContext()
        }
        return engine
}
```

- 看 `engine.Use`:

```go
func (engine *Engine) Use(middleware ...HandlerFunc) IRoutes {
        engine.RouterGroup.Use(middleware...)
        engine.rebuild404Handlers()
        engine.rebuild405Handlers()
        return engine
}
```

`engine.RouterGroup.Use`:

```go
func (group *RouterGroup) Use(middleware ...HandlerFunc) IRoutes {
        group.Handlers = append(group.Handlers, middleware...)
        return group.returnObj()
    }
```

- 接下来回到demo，看 `r.Run()`:

```go
func (engine *Engine) Run(addr ...string) (err error) {
        defer func() { debugPrintError(err) }()

        address := resolveAddress(addr)
        debugPrint("Listening and serving HTTP on %s\n", address)
        err = http.ListenAndServe(address, engine)
        return
}
```

看过 `net/http` 的同学应该知道，在Go里只要实现 `ServeHTTP` 就可以，所以我们找一下：

```go
func (engine *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
        c := engine.pool.Get().(*Context)
        c.writermem.reset(w)
        c.Request = req
        c.reset()

        engine.handleHTTPRequest(c)

        engine.pool.Put(c)
}
```

这就是处理流程，请求来了，从 `engine.pool` 里拿一个空的context，丢到 `engine.handleHTTPRequest` 处理，然后回收。

- 看 `engine.handleHTTPRequest`:

```go
func (engine *Engine) handleHTTPRequest(c *Context) {
        httpMethod := c.Request.Method
        path := c.Request.URL.Path
        unescape := false
        if engine.UseRawPath && len(c.Request.URL.RawPath) > 0 {
                path = c.Request.URL.RawPath
                unescape = engine.UnescapePathValues
        }

        // Find root of the tree for the given HTTP method
        t := engine.trees
        for i, tl := 0, len(t); i < tl; i++ {
                if t[i].method == httpMethod {
                        root := t[i].root
                        // Find route in tree
                        handlers, params, tsr := root.getValue(path, c.Params, unescape)
                        if handlers != nil {
                                c.handlers = handlers
                                c.Params = params
                                c.Next()
                                c.writermem.WriteHeaderNow()
                                return
                        }
                        if httpMethod != "CONNECT" && path != "/" {
                                if tsr && engine.RedirectTrailingSlash {
                                        redirectTrailingSlash(c)
                                        return
                                }
                                if engine.RedirectFixedPath && redirectFixedPath(c, root, engine.RedirectFixedPath) {
                                        return
                                }
                        }
                        break
                }
        }

        if engine.HandleMethodNotAllowed {
                for _, tree := range engine.trees {
                        if tree.method != httpMethod {
                                if handlers, _, _ := tree.root.getValue(path, nil, unescape); handlers != nil {
                                        c.handlers = engine.allNoMethod
                                        serveError(c, 405, default405Body)
                                        return
                                }
                        }
                }
        }
        c.handlers = engine.allNoRoute
        serveError(c, 404, default404Body)
}
```

大致的流程就是从路由里找出handler，然后进行处理。其中路由使用 `httprouter` 实现，使用的数据结构是基数树(radix tree)。

- 看 `c.Next()`:

```go
func (c *Context) Next() {
        c.index++
        for s := int8(len(c.handlers)); c.index < s; c.index++ {
                c.handlers[c.index](c)
        }
}
```

最开始的时候 `c.index` 为0值，所以会执行 `c.handlers` 里面的第一个handler，然后一个个执行下去。

- 我们来看看中间件的原理。先看demo：

```go
package main

import (
	"fmt"
	"os"
	"time"

	"github.com/gin-gonic/contrib/ginrus"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func main() {
	r := gin.New()

	r.Use(ginrus.Ginrus(logrus.StandardLogger(), time.RFC3339, true))

	// Example ping request.
	r.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong "+fmt.Sprint(time.Now().Unix()))
	})

	// Listen and Server in 0.0.0.0:8080
	r.Run(":8080")
}
```

来看一下 `r.Use`:

```go
func (engine *Engine) Use(middleware ...HandlerFunc) IRoutes {
        engine.RouterGroup.Use(middleware...)
        engine.rebuild404Handlers()
        engine.rebuild405Handlers()
        return engine
}
```

继续追:

```go
func (group *RouterGroup) Use(middleware ...HandlerFunc) IRoutes {
        group.Handlers = append(group.Handlers, middleware...)
        return group.returnObj()
}
```

然后发现没有然后了，那么当我们调用 `r.GET` 的时候发生了什么呢？追一下发现其实就是调用了 `RouterGroup.GET`

```go
func (group *RouterGroup) GET(relativePath string, handlers ...HandlerFunc) IRoutes {
        return group.handle("GET", relativePath, handlers)
}
```

继续：

```go
func (group *RouterGroup) handle(httpMethod, relativePath string, handlers HandlersChain) IRoutes {
        absolutePath := group.calculateAbsolutePath(relativePath)
        handlers = group.combineHandlers(handlers)
        group.engine.addRoute(httpMethod, absolutePath, handlers)
        return group.returnObj()
}
```

这里会把绝对路由算出来，然后加到树里。里面有一步 `handlers = group.combineHandlers(handlers)`:

```go
func (group *RouterGroup) combineHandlers(handlers HandlersChain) HandlersChain {
        finalSize := len(group.Handlers) + len(handlers)
        if finalSize >= int(abortIndex) {
                panic("too many handlers")
        }
        mergedHandlers := make(HandlersChain, finalSize)
        copy(mergedHandlers, group.Handlers)
        copy(mergedHandlers[len(group.Handlers):], handlers)
        return mergedHandlers
}
```

看到了 `group.Handlers` 嘛？上面的 `Use` 也有用到。所以其实 `Use` 和 `GET`, `POST` 等是一样的。 而中间件则是通过 `c.Next` 先执行接下来的handlers，然后返回再执行当前的。比如 `Ginrus`:

```go
ckage ginrus

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type loggerEntryWithFields interface {
	WithFields(fields logrus.Fields) *logrus.Entry
}

// Ginrus returns a gin.HandlerFunc (middleware) that logs requests using logrus.
//
// Requests with errors are logged using logrus.Error().
// Requests without errors are logged using logrus.Info().
//
// It receives:
//   1. A time package format string (e.g. time.RFC3339).
//   2. A boolean stating whether to use UTC time zone or local.
func Ginrus(logger loggerEntryWithFields, timeFormat string, utc bool) gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		// some evil middlewares modify this values
		path := c.Request.URL.Path
		c.Next()

		end := time.Now()
		latency := end.Sub(start)
		if utc {
			end = end.UTC()
		}

		entry := logger.WithFields(logrus.Fields{
			"status":     c.Writer.Status(),
			"method":     c.Request.Method,
			"path":       path,
			"ip":         c.ClientIP(),
			"latency":    latency,
			"user-agent": c.Request.UserAgent(),
			"time":       end.Format(timeFormat),
		})

		if len(c.Errors) > 0 {
			// Append error field if this is an erroneous request.
			entry.Error(c.Errors.String())
		} else {
			entry.Info()
		}
	}
}
```

到此，Over 