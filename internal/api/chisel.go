package api

import (
	"context"
	"fmt"
	chclient "github.com/jpillora/chisel/client"
	"github.com/kintohub/kinto-cli/internal/types"
	"github.com/kintohub/kinto-cli/internal/utils"
	"io"
	"strconv"
	"time"
)

func (a *Api) StartTeleport(blocksToForward []RemoteConfig, envId string, clusterId string) {

	var host *types.TeleportResponse

	resp, err := a.getKubeCoreService(clusterId, envId).StartTeleport(
		context.Background(), &types.TeleportRequest{EnvId: envId})

	if err != nil {
		utils.TerminateWithError(err)
	}

	host, err = resp.Recv()

	if err == io.EOF {
		utils.TerminateWithCustomError("stream has no data!")
	}
	if err != nil {
		utils.TerminateWithError(err)
	}

	var remotes []string

	for _, remote := range blocksToForward {
		remotes = append(remotes, fmt.Sprintf(remote.FromHost+":"+strconv.Itoa(remote.FromPort)+
			":"+remote.ToHost+":"+strconv.Itoa(remote.ToPort)))
	}

	chiselClient, err := chclient.NewClient(&chclient.Config{
		MaxRetryInterval: 1 * time.Second,
		MaxRetryCount:    50,
		Server:           "https://" + host.Data.Host,
		Auth:             host.Data.Credentials,
		Remotes:          remotes,
		KeepAlive:        10 * time.Second,
	})
	if err != nil {
		utils.TerminateWithError(err)
	}

	//chiselClient.Logger.Info = false

	utils.NoteMessage("Starting Tunnel")

	// Run chisel client in background
	go func() {
		err = chiselClient.Run()
		if err != nil {
			utils.TerminateWithError(err)
		}
	}()

	// Run infinite stream connection
	go func() {
		_, err := resp.Recv()
		if err != nil {
			utils.TerminateWithError(err)
		}
	}()

	if err != nil {
		utils.TerminateWithError(err)
	}

	for _, remote := range blocksToForward {
		utils.InfoMessage(fmt.Sprintf("> Forwarding: %s:%d => %s:%d",
			remote.FromHost, remote.FromPort, remote.ToHost, remote.ToPort))
	}

	utils.SuccessMessage("✓ Connected!")
	defer chiselClient.Close()
	utils.NoteMessage("Press any key to close the tunnel")
	fmt.Scanln()
}
