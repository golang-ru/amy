package main

import (
	"fmt"
	"log"

	"github.com/golang-ru/amy/handler"
)

func main() {
	handlers := []handler.IHandler{
		handler.NewJobber(),
	}

	fmt.Println("Running handlers...")

	for _, h := range handlers {
		if err := h.Run(); err != nil {
			log.Fatalf("Undefined error: %v\n", err)
		}
	}
}
