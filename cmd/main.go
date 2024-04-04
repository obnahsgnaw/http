package main

import (
	"github.com/gin-gonic/gin"
	"github.com/obnahsgnaw/http"
	"github.com/obnahsgnaw/http/engine"
	"log"
	http2 "net/http"
)

func main() {
	e, err := http.Default("127.0.0.1", 9011, &engine.Config{
		Name: "test",
	})
	if err != nil {
		log.Fatal(err.Error())
	}
	defer e.Close()
	e.Engine().GET("/b", func(context *gin.Context) {
		context.String(http2.StatusOK, "ok")
	})
	log.Println("http server start...")
	e.RunAndServ()
}
