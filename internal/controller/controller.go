package controller

import (
	"github.com/kintohub/kinto-cli/internal/api"
)

type ControllerInterface interface {
	Init(masterHost string)
	Login()
	Environment()
	EnvironmentAccess(envId ...string)
	Version()
	Services(envId ...string)
	ServiceAccess(envId string,blockId string)
	Access()
	Teleport()
	Status()
}

type Controller struct {
	api api.ApiInterface
}

func NewControllerOrDie(api api.ApiInterface) ControllerInterface {
	return &Controller{
		api: api,
	}
}
