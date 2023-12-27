package main

import (
	"github.com/gin-gonic/gin"
	"github.com/obnahsgnaw/application/pkg/url"
	"github.com/obnahsgnaw/http"
	"github.com/obnahsgnaw/http/engine"
	"log"
	http2 "net/http"
)

func main() {
	e, _ := http.Default(url.Host{Ip: "127.0.0.1", Port: 9000}, &engine.Config{
		Name:     "test",
		LogDebug: true,
	})
	e.Engine().GET("/b", func(context *gin.Context) {
		context.String(http2.StatusOK, "ok")
	})
	log.Println("http server start...")
	e.RunAndServ()
}
