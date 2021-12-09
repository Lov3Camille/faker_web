package main

import (
	core "faker_web/core"
	"log"
	"net/http"
	"time"
)

func onlyForV2() core.HandlerFunc {
	return func(c *core.Context) {
		t := time.Now()
		log.Printf("[%d] %s in %v for group v2", c.StatusCode, c.Req.RequestURI, time.Since(t))
	}
}

func main() {
	r := core.Default()
	r.GET("/", func(c *core.Context) {
		c.String(http.StatusOK, "吴诗墨最可爱！\n")
	})

	r.GET("/hello", func(c *core.Context) {
		c.String(http.StatusOK, "hello %s, you're at %s\n", c.Query("name"), c.Path)
	})

	r.GET("/hello/:name", func(c *core.Context) {
		c.String(http.StatusOK, "hello %s, you're at %s\n", c.Param("name"), c.Path)
	})

	r.GET("/panic", func(c *core.Context) {
		names := []string{"faker"}
		c.String(http.StatusOK, names[100])
	})

	r.Run(":9999")
}
