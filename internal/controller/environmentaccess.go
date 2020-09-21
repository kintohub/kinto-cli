package controller

import (
	"github.com/kintohub/kinto-cli/internal/api"
	"github.com/kintohub/kinto-cli/internal/utils"
)

func (c *Controller) EnvironmentAccess(envId ...string) {

	utils.CheckLogin()
	utils.StartSpinner()

	var clusterId string
	var selectedEnvId string
	var blocksToForward []api.RemoteConfig

	if len(envId) == 0 {

		envDetails := c.GetAvailableEnvNames(false)
		utils.StopSpinner()
		selectedEnvId, clusterId = SelectionPrompt(envDetails)
		utils.StartSpinner()
		blocksToForward = c.GetBlocksToForward(selectedEnvId)

	} else {

		env := c.GetEnvFromId(envId[0])
		blocksToForward = c.GetBlocksToForward(env.Id)
		clusterId = env.ClusterId
		selectedEnvId = env.Id

	}

	if len(blocksToForward) != 0 {
		utils.StopSpinner()
		c.api.StartAccess(blocksToForward, selectedEnvId, clusterId)

	} else {
		utils.WarningMessage("No Accessible service/s found!")
	}

}
