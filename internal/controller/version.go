package controller

import (
	"fmt"
	"github.com/kintohub/kinto-cli-go/internal/consts"
)

func (c Controller) Version() {
	fmt.Printf("Kinto Command Line Interface (CLI) %s", consts.Version)
}
