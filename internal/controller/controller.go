package controller

import (
	"github.com/kintohub/kinto-cli-go/internal/api"
)

type ControllerInterface interface {
	Init()
	Register()
	Login()
	Environment()
	Version()
}

type Controller struct {
	api api.ApiInterface
}

func NewControllerOrDie(api api.ApiInterface) ControllerInterface {
	return &Controller{
		api: api,
	}
}
