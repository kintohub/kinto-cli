package controller

import (
	"fmt"
	"github.com/gookit/color"
	"github.com/kintohub/kinto-cli-go/internal/config"
	"github.com/kintohub/kinto-cli-go/internal/utils"
)

func (c *Controller) Teleport(envId string) {

	var remotes []string
	inc := 0
	localGitUrl := utils.GetLocalGitUrl()
	blocks, err := c.api.GetBlocks(envId)
	if err != nil {
		utils.TerminateWithError(err)
	}

	for _, block := range blocks {
		latestRelease := utils.GetLatestSuccessfulRelease(block.Releases)

		if latestRelease.BuildConfig.Repository.Url == localGitUrl {
			port := config.LocalPort + inc
			remote := fmt.Sprintf("%s:%d:%s:%s", "localHost",
				utils.CheckPort(port), block.Name, "80")
			remotes = append(remotes, remote)
			inc += 1
		}
	}
	if len(remotes) != 0 {
		c.api.CreateTeleport(remotes)
	} else {
		color.Yellow.Printf("\nNo service/s found in this environment to port-forward!\n")
	}

}
