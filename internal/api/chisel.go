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

func (a *Api) StartAccess(blocksToForward []RemoteConfig, envId string, clusterId string) {

	// Default time to cancel is 30 minutes for our nginx gateway
	// TODO: Move to env var / build arg
	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(time.Minute*30))
	streamResponse, err := a.getKubeCoreService(clusterId, envId).StartTeleport(
		ctx, &types.TeleportRequest{EnvId: envId})
	defer cancel()

	if err != nil {
		utils.TerminateWithError(err)
	}
	StartChisel(blocksToForward, streamResponse)

}

func (a *Api) StartTeleport(blocksToForward []RemoteConfig, envId string, clusterId string, blockName string) {

	// Default time to cancel is 30 minutes for our nginx gateway
	// TODO: Move to env var / build arg
	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(time.Minute*30))
	streamResponse, err := a.getKubeCoreService(clusterId, envId).StartTeleport(
		ctx, &types.TeleportRequest{EnvId: envId, BlockName: blockName})
	defer cancel()

	if err != nil {
		utils.TerminateWithError(err)
	}

	StartChisel(blocksToForward, streamResponse)

}

func StartChisel(blocksToForward []RemoteConfig, streamResponse types.KintoKubeCoreService_StartTeleportClient) {

	host, err := streamResponse.Recv()

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
		Server:           fmt.Sprintf("https://%s", host.Data.Host),
		Auth:             host.Data.Credentials,
		Remotes:          remotes,
		KeepAlive:        10 * time.Second,
	})
	defer chiselClient.Close()

	if err != nil {
		utils.TerminateWithError(err)
	}

	// Run infinite stream connection
	go func() {
		_, err := streamResponse.Recv()
		if err != nil {
			utils.TerminateWithError(err)
		}
	}()

	chiselClient.Logger.Info = false

	fmt.Println("")
	utils.InfoMessage("Starting Tunnel")

	// Run chisel client in background
	go func() {
		err = chiselClient.Run()
		if err != nil {
			utils.TerminateWithError(err)
		}
	}()

	if err != nil {
		utils.TerminateWithError(err)
	}

	for _, remote := range blocksToForward {
		utils.InfoMessage(fmt.Sprintf("Forwarding > %s:%d => %s:%d",
			remote.FromHost, remote.FromPort, remote.ToHost, remote.ToPort))
	}

	utils.SuccessMessage("Connected!")
	utils.NoteMessage("\nPress any key to close the tunnel")
	fmt.Scanln()
	utils.NoteMessage("Connection Closed")
}
