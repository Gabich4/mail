package main

import (
	"bodyshop/service"
	"log"
)

func main() {
	service := service.New()

	if err := service.Server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
