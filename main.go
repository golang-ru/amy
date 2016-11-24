package main

import (
	"fmt"
	"log"

	"github.com/golang-ru/amy/core"
	"github.com/golang-ru/amy/handler"
	"github.com/golang-ru/amy/service"
)

func main() {
	config, err := core.NewConfig()
	if err != nil {
		log.Fatalf("[Config] Undefined error: %v\n", err)
	}

	services := service.NewService(config)

	handlers := []handler.IHandler{
		handler.NewJobber(services),
	}

	fmt.Println("Running handlers...")

	for _, h := range handlers {
		if err := h.Run(); err != nil {
			log.Fatalf("[Main] Undefined error: %v\n", err)
		}
	}
}
