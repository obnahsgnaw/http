package main

import (
	"github.com/obnahsgnaw/application/pkg/url"
	"github.com/obnahsgnaw/http"
	"log"
)

func main() {
	e, _ := http.New(&http.Config{
		Name:     "test",
		LogDebug: true,
	})

	log.Println("server start...")
	//log.Fatal(e.Run(":9000"))
	pe := http.NewPortedEngine(e, url.Host{Port: 9000})
	log.Fatal(pe.Run())
}
