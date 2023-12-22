package main

import (
	"github.com/obnahsgnaw/application/pkg/url"
	"github.com/obnahsgnaw/http"
	"github.com/obnahsgnaw/http/engine"
	"log"
)

func main() {
	e, _ := http.Default(url.Host{Ip: "127.0.0.1", Port: 9000}, &engine.Config{
		Name:     "test",
		LogDebug: true,
	})

	log.Println("http server start...")
	log.Fatal(e.Run())
}
