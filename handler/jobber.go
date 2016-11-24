package handler

import (
	"fmt"
	"time"
)

type Jobber struct {
}

func NewJobber() *Jobber {
	return &Jobber{}
}

func (j Jobber) Run() error {
	for {
		fmt.Println("I'm working, of course!")

		time.Sleep(5 * time.Second)
	}
}
