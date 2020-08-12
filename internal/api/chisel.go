package api

import (
	chclient "github.com/jpillora/chisel/client"
	"github.com/kintohub/kinto-cli-go/internal/config"
	"github.com/kintohub/kinto-cli-go/internal/utils"
	"time"
)

func (a *Api) CreateTeleport(remotes []string) *chclient.Client {

	chiselClient, err := chclient.NewClient(&chclient.Config{
		KeepAlive:        time.Second,
		MaxRetryInterval: time.Second,
		Server:           config.ChiselHost,
		Remotes:          remotes,
	})
	if err != nil {
		utils.TerminateWithError(err)
	}
	chiselClient.Logger.Info = false
	return chiselClient
}
