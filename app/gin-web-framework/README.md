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

# References
Official documentation of Gin (Simple without principles explanation): [https://gin-gonic.com/docs/](https://gin-gonic.com/docs/)

Official examples of Gin: [https://github.com/gin-gonic/examples](https://github.com/gin-gonic/examples)


