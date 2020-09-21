package controller

import (
	"github.com/kintohub/kinto-cli/internal/api"
	"github.com/kintohub/kinto-cli/internal/utils"
)

func (c *Controller) Access() {

	utils.CheckLogin()
	utils.StartSpinner()

	var envDetails []api.EnvDetails

	envDetails = c.GetAvailableEnvNames(false)

	utils.StopSpinner()
	selectedEnvId, clusterId := SelectionPrompt(envDetails)
	utils.StartSpinner()
	var blocksToForward []api.RemoteConfig
	blocksToForward = c.GetBlocksToForward(selectedEnvId)

	if len(blocksToForward) != 0 {
		utils.StopSpinner()
		c.api.StartAccess(blocksToForward, selectedEnvId, clusterId)

	} else {
		utils.WarningMessage("No service/s found in this environment to teleport into!")
	}

}
