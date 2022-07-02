package main

import (
	"log"

	"github.com/Sraik25/go-hexagonal_http_api/02-01-architectured-gin-healthcheck/cmd/api/bootstrap"
)

func main() {
	if err := bootstrap.Run(); err != nil {
		log.Fatal(err)
	}
}
