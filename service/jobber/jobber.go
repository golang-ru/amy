package jobber

import "github.com/golang-ru/amy/core"

type IJobber interface {
	Test() error
}

type service struct {
	config core.IConfig
}

func NewService(config core.IConfig) IJobber {
	return &service{config: config}
}
