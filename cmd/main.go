package main

import (
	cli2 "github.com/kintohub/kinto-cli-go/internal/cli"
	"github.com/kintohub/kinto-cli-go/internal/controller"
)

func main() {
	c := controller.InitControllerOrDie()
	cli := cli2.NewCliOrDie(c)
	cli.Execute()
}
