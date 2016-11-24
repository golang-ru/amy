package handler

import (
	"fmt"
	"time"

	"github.com/golang-ru/amy/service"
)

type Jobber struct {
	Services service.IService
}

func NewJobber(services service.IService) *Jobber {
	return &Jobber{Services: services}
}

func (j Jobber) Run() error {
	for {
		fmt.Println("I'm working, of course!")

		time.Sleep(5 * time.Second)
	}
}
