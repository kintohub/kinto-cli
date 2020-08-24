package api

import (
	"context"
	"fmt"
	chclient "github.com/jpillora/chisel/client"
	"github.com/kintohub/kinto-cli/internal/types"
	"github.com/kintohub/kinto-cli/internal/utils"
	"io"
	"strconv"
	"sync"
	"time"
)

func (a *Api) StartTeleport(blocksToForward []RemoteConfig, envId string, clusterId string) {

	var host *types.TeleportResponse
	var wg sync.WaitGroup
	resp, err := a.getKubeCoreService(clusterId, envId).StartTeleport(
		context.Background(), &types.TeleportRequest{EnvId: envId})

	if err != nil {
		utils.TerminateWithError(err)
	}

	for {
		host, err = resp.Recv()
		if err == io.EOF {
			utils.TerminateWithCustomError("stream has no data!")
		}
		if err != nil {
			utils.TerminateWithError(err)
		}

		var remotes []string
		var err error

		for _, remote := range blocksToForward {
			remotes = append(remotes, fmt.Sprintf(remote.FromHost+":"+strconv.Itoa(remote.FromPort)+
				":"+remote.ToHost+":"+strconv.Itoa(remote.ToPort)))
		}

		chiselClient, err := chclient.NewClient(&chclient.Config{

			MaxRetryInterval: 1 * time.Second,
			MaxRetryCount:    5,
			Server:           "https://" + host.Data.Host,
			Auth:             host.Data.Credentials,
			Remotes:          remotes,
		})
		if err != nil {
			utils.TerminateWithError(err)
		}

		//chiselClient.Logger.Info = false

		utils.NoteMessage("Starting Tunnel")

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
		defer resp.CloseSend()
		utils.NoteMessage("Press any key to close the tunnel")
		fmt.Scanln()

	}

}
