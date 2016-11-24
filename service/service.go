package service

import (
	"github.com/golang-ru/amy/core"
	"github.com/golang-ru/amy/service/jobber"
)

type IService interface {
	Jobber() jobber.IJobber
}

type service struct {
	config core.IConfig
}

func NewService(config core.IConfig) IService {
	return &service{config: config}
}

func (s *service) Jobber() jobber.IJobber {
	return jobber.NewService(s.config)
}
