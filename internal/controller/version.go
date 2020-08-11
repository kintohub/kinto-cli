package controller

import (
	"fmt"
	"github.com/kintohub/kinto-cli-go/internal/config"
	"github.com/kintohub/kinto-cli-go/internal/utils"
)

func (c Controller) Version() {
	utils.InfoMessage(fmt.Sprintf("Kinto Command Line Interface (CLI) %s", config.Version))
}
