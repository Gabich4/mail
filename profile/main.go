package main

import (
	"profile/common"
	"profile/service"
)

func main() {
	service, err := service.New()
	if err != nil {
		common.Logger.Fatal(err)
	}

	if err := service.Serve(); err != nil {
		common.Logger.Fatal(err)
	}
}
