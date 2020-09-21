package controller

import (
	"github.com/kintohub/kinto-cli/internal/api"
	"github.com/kintohub/kinto-cli/internal/config"
	"github.com/kintohub/kinto-cli/internal/utils"
)

func (c *Controller) ServiceAccess(envId string, blockId string) {

	utils.CheckLogin()
	utils.StartSpinner()
	var blocksToForward []api.RemoteConfig

	env := c.GetEnvFromId(envId)

	blocks, err := c.api.GetBlocks(env.Id)
	if err != nil {
		utils.TerminateWithError(err)
	}
	for _, block := range blocks {
		if block.Id == blockId {

			latestRelease := utils.GetLatestSuccessfulRelease(block.Releases)

			if latestRelease != nil {
				remote := api.RemoteConfig{FromHost: "localhost", FromPort: utils.CheckPort(config.LocalPort),
					ToHost: block.Name, ToPort: utils.GetBlockPort(block)}
				blocksToForward = append(blocksToForward, remote)
			}
			break
		}
	}

	if len(blocksToForward) != 0 {
		utils.StopSpinner()
		c.api.StartAccess(blocksToForward, env.Id, env.ClusterId)

	} else {
		utils.WarningMessage("No Accessible service/s found!")
	}

}
