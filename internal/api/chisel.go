package api

import (
	"fmt"
	"github.com/gookit/color"
	chclient "github.com/jpillora/chisel/client"
	"github.com/kintohub/kinto-cli-go/internal/config"
	"github.com/kintohub/kinto-cli-go/internal/utils"
	"strings"
	"time"
)

func (a *Api) CreateTeleport(remotes []string) {

	c, err := chclient.NewClient(&chclient.Config{
		KeepAlive:        time.Second,
		MaxRetryInterval: time.Second,
		Server:           config.ChiselHost,
		Remotes:          remotes,
	})

	if err != nil {
		utils.TerminateWithError(err)
	}
	c.Logger.Info = false

	color.Yellow.Printf("\nStarting tunnel!\n")
	for _, proxy := range remotes {
		split := strings.Split(proxy, ":")
		fromHost, fromPort, toBlock, toPort := split[0], split[1], split[2], split[3]
		color.Magenta.Printf("# Forwarding:")
		color.Gray.Printf(" %s:%s => %s:%s\n", fromHost, fromPort, toBlock, toPort)
	}

	// Run indefinitely.
	go c.Run()

	color.Green.Printf("Connected!\n")

	defer c.Close()

	//TODO: this should not be here.
	fmt.Println("\nPress any key to close the tunnel")
	fmt.Scanln()

}
