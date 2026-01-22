package svc

import "common/internal/config"

type ServiceContext struct {
	Config config.Config
	App    config.App
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		App:    c.App,
	}
}
