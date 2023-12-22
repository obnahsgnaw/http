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

	//log.Fatal(e.Run(":9000"))
	host := url.Host{Ip: "127.0.0.1", Port: 9000}
	//pe := http.NewPortedEngine(e, host)
	l, _ := http.NewListener(host)
	pe := http.NewListenerEngine(e, host, l)
	log.Println("server start...")
	log.Fatal(pe.Run())
}
