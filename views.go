package main

import (
	"fmt"
	"github.com/huawen0327/golee"
	"time"
)

var index = func(c *golee.Context) {
	c.HTMLResponse(golee.StatusOK, "<h1>index</h1>")
}

var hello = func(c *golee.Context) {
	c.StringResponse(golee.StatusOK, "hello %s, you're at %s, the parm is %s\n", c.GET("name"), c.Path, c.Parm)
}
var logger = func(c *golee.Context) {
	t := time.Now()
	c.Next()
	fmt.Printf("[%d] %s in %v", c.StatusCode, c.Req.RequestURI, time.Since(t))
}

var temp = func(c *golee.Context) {
	c.HTMLRender(golee.StatusOK, nil, "templates/index.html")
}
