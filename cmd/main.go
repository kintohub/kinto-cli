package main

import (
	"github.com/kintohub/kinto-cli-go/internal/api"
	"github.com/kintohub/kinto-cli-go/internal/cli"
	"github.com/kintohub/kinto-cli-go/internal/controller"
)

func main() {
	cli := cli.NewCliOrDie()
	api := api.NewApiOrDie(cli.GetHostFlag())
	controller := controller.NewControllerOrDie(api)
	cli.Execute(controller)
}
