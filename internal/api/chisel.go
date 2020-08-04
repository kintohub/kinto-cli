package api

import (
	"fmt"
	"github.com/gookit/color"
	chclient "github.com/jpillora/chisel/client"
	"github.com/kintohub/kinto-cli-go/internal/utils"
	"time"
)

func (a *Api) CreateTunnel() {
	c, err := chclient.NewClient(&chclient.Config{
		KeepAlive:        time.Second,
		MaxRetryInterval: time.Second,
		Server:           "https://chisel-5f28e.asia1.kinto.io",
		Remotes: []string{
			"postgresql:5432:postgresql:5432",
			"mongodb-0:27019:mongodb-0:27017",
			//"mongo-1:mongo-1:27017",
			//"mongo-2:mongo-2:27017",
		},
	})

	if err != nil {
		utils.TerminateWithError(err)
	}

	color.Green.Printf("\nStarting tunnel!\n")

	// Run indefinitely.
	go c.Run()
	defer c.Close()

	//TODO: this should not be here.
	fmt.Println("Press any key to close tunnel")
	fmt.Scanln()

}
