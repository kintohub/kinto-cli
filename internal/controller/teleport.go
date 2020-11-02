package controller

import (
	"github.com/kintohub/kinto-cli/internal/api"
	"github.com/kintohub/kinto-cli/internal/utils"
)

func (c *Controller) Teleport() {

	utils.CheckLogin()
	utils.StartSpinner()

	var envDetails []api.EnvDetails


	envDetails = c.GetAvailableEnvNames(true)

	utils.StopSpinner()
	selectedEnvId, clusterId := SelectionPrompt(envDetails)
	utils.StartSpinner()
	var blocksToForward []api.RemoteConfig
	var blockNameToTeleport string
	blocksToForward, blockNameToTeleport = c.GetBlocksToTeleport(selectedEnvId)

	if len(blocksToForward) != 0 {
		utils.StopSpinner()
		c.api.StartTeleport(blocksToForward, selectedEnvId, clusterId, blockNameToTeleport)

	} else {
		utils.WarningMessage("No service/s found in this environment to teleport into!")
	}

}

