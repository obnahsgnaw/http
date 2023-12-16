package main

import (
	"github.com/obnahsgnaw/http"
	"log"
)

func main() {
	e, _ := http.New(&http.Config{
		Name:     "test",
		LogDebug: false,
	})

	log.Println("ok")
	log.Fatal(e.Run(":9000"))
}
