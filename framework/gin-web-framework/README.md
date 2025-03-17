# Motivation
This is to explore the capabilities of Gin Web Framework of Go

# Background
Go supports native web capabilities net/http, which has more limitations under complex scenarios that Gin can compensate.
In this case, net/http is only needed when `net/http/httptest` as API test for Gin and constants such as `http.StatusOK`

Gin is an HTTP web framework, which features a Martini-like API (是说开发Web功能的这个框架接口，不是说Martini API是另一种和Restful API同层级的API),
but with performance up to 40 times faster than Martini.

# Fundamentals
- `r := gin.Default()` returns an ` *Engine` instance with the *Logger* and *Recovery* middleware already attached whose core role is a **router**.
- Each request is referenced as a `c *gin.Context` instance in the handler function parameter, which is a wrapper of `http.Request` and `http.ResponseWriter`.
    - Core members include `c.Param`, `c.Query`, `c.ClientIP`, etc.
    - c.Param是路径参数，？前面的替换内容；c.Query是查询参数，？后面的内容。比如`/room/123?nick=alice`，c.Param["id"]是"123"，c.Query["nick"]是"alice"

# 核心对象
## gin.Context
- 它代表了一个HTTP请求和响应的上下文，是对net/http包中的Request和ResponseWriter的封装，提供了方便的方法来获取请求信息和写入响应信息。
- 它与HTTP格式相关但不完全一致，主要还是支持用户对HTTP标准格式处理后自定义添加字段。
### 核心方法
**c.Next()**
- 用于执行后续的处理函数(下一个中间件或者最后的应用程序)，如果不调用，后续的处理函数将不会被执行。
- 中间件的执行顺序与它们被注册的顺序(即r.Use()的顺序)一致。
- 最后一个中间件调用c.Next()后，会执行路由的处理程序。
- 注意c.Next()必须显式调用，否则请求将不会到达后续中间件或者路由处理程序，客户端可能会收到一个空的响应或超时。
**c.Abort()**
- 用于阻止后续的处理函数执行，但是当前处理函数仍然会继续执行。

# References
Official documentation of Gin (Simple without principles explanation): [https://gin-gonic.com/docs/](https://gin-gonic.com/docs/)

Official examples of Gin: [https://github.com/gin-gonic/examples](https://github.com/gin-gonic/examples)

