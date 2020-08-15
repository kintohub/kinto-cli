package api

import (
	"fmt"
	chclient "github.com/jpillora/chisel/client"
	"github.com/kintohub/kinto-cli-go/internal/config"
	"github.com/kintohub/kinto-cli-go/internal/utils"
	"strconv"
	"sync"
	"time"
)

var chiselClient *chclient.Client

func (a *Api) StartTeleport(blocksToForward []RemoteConfig) {

	var remotes []string
	var err error
	var wg sync.WaitGroup

	for _, remote := range blocksToForward {
		remotes = append(remotes, fmt.Sprintf(remote.FromHost+":"+strconv.Itoa(remote.FromPort)+
			":"+remote.ToHost+":"+strconv.Itoa(remote.ToPort)))
	}

	chiselClient, err = chclient.NewClient(&chclient.Config{
		KeepAlive:        time.Second,
		MaxRetryInterval: time.Second,
		Server:           config.ChiselHost,
		Remotes:          remotes,
	})
	if err != nil {
		utils.TerminateWithError(err)
	}

	chiselClient.Logger.Info = false

	utils.WarningMessage("\nStarting Tunnel")

	go func() {
		err = chiselClient.Run()
	}()
	wg.Wait()

	if err != nil {
		utils.TerminateWithError(err)
	}

	for _, remote := range blocksToForward {
		utils.InfoMessage(fmt.Sprintf("> Forwarding: %s:%d => %s:%d",
			remote.FromHost, remote.FromPort, remote.ToHost, remote.ToPort))
	}

	utils.SuccessMessage("✓ Connected!")
	defer chiselClient.Close()
	utils.WarningMessage("\nPress any key to close the tunnel")
	fmt.Scanln()
}
