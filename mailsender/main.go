package main

import (
	"log"
	"mailsender/service"
)

func main() {
	service, err := service.New()
	if err != nil {
		log.Fatal(err)
	}

	if err := service.Server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
