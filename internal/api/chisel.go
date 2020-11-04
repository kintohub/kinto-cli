package api

import (
	"context"
	"fmt"
	chclient "github.com/jpillora/chisel/client"
	"github.com/kintohub/kinto-cli/internal/config"
	"github.com/kintohub/kinto-cli/internal/types"
	"github.com/kintohub/kinto-cli/internal/utils"
	"io"
	"strconv"
	"strings"
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
	StartChisel(blocksToForward, streamResponse, false)

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

	StartChisel(blocksToForward, streamResponse, true)

}

func StartChisel(blocksToForward []RemoteConfig, streamResponse types.KintoKubeCoreService_StartTeleportClient, isTeleport bool) {

	host, err := streamResponse.Recv()

	if err == io.EOF {
		utils.TerminateWithCustomError("stream has no data!")
	}
	if err != nil {
		utils.TerminateWithError(err)
	}

	var remotes []string

	for _, remote := range blocksToForward {
		newRemote := fmt.Sprintf(
			"%s:%s:%s:%s",
			remote.FromHost,
			strconv.Itoa(remote.FromPort),
			remote.ToHost,
			strconv.Itoa(remote.ToPort),
		)
		remotes = append(remotes, newRemote)
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
	for _, remote := range blocksToForward {
		if strings.Contains(remote.FromHost, "R:0.0.0.0") {
			utils.InfoMessage(fmt.Sprintf("Teleporting > %s:%d => %s:%d",
				remote.ToHost, remote.ToPort, remote.FromHost, remote.FromPort))
		} else {
			utils.InfoMessage(fmt.Sprintf("Forwarding  >  %s:%d => %s:%d",
				remote.FromHost, remote.FromPort, remote.ToHost, remote.ToPort))
		}
	}
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

	//Show start server message only if teleporting
	// For the lack of a better way, this is a temporary workaround for validating connection status
	if isTeleport {
		fmt.Println("")
		//TODO: make teleport port configurable
		utils.WarningMessage(
			fmt.Sprintf("Please start your local server at PORT => %d",
				config.DefaultTeleportPort))
		for utils.CheckIfPortOpen(config.DefaultTeleportPort, false) {
			time.Sleep(1 * time.Second)
			utils.StartSpinner()
		}
	}

	utils.StartSpinner()
	utils.SuccessMessage("Connected!")
	fmt.Println("")
	utils.NoteMessage("Press any key to close the tunnel")
	fmt.Scanln()
	utils.NoteMessage("Connection Closed")

}
