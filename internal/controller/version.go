package controller

import (
	"fmt"
	"github.com/kintohub/kinto-cli-go/internal/config"
)

func (c Controller) Version() {

	fmt.Printf("Kinto Command Line Interface (CLI) %s", config.Version)
}
