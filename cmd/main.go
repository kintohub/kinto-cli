package main

import (
	"github.com/kintohub/kinto-cli-go/internal/cli"
	"github.com/kintohub/kinto-cli-go/internal/controller"
)

func main() {
	controller := controller.InitControllerOrDie()
	cli := cli.NewCliOrDie(controller)
	cli.Execute()
}
