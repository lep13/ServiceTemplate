package main

import (
	"fmt"
	"log"
	"service-template/cmd/api"
)

const (
	host = "localhost"
	port = "8080"
)

func main() {

	err := api.ServeHTTP(fmt.Sprintf("%s:%s", host, port))
	if err != nil {
		log.Fatalf("[ERROR]: %s", err)
	}
}
