package controller

import (
	"github.com/kintohub/kinto-cli/internal/api"
)

type ControllerInterface interface {
	Init(masterHost string)
	Login()
	Environment()
	Version()
	Services(envId ...string)
	Teleport(teleportAllFlag bool)
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
