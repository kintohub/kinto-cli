package controller

import (
	"github.com/kintohub/kinto-cli/internal/utils"
)

//allows to teleport a particular service inside an env while port-forwarding the remaining services.
func (c *Controller) Teleport() {

	utils.CheckLogin()
	utils.StartSpinner()

	envDetails := c.GetAvailableEnvNames(true)

	utils.StopSpinner()
	selectedEnvId, clusterId := SelectionPrompt(envDetails)
	utils.StartSpinner()

	blocksToForward, blockNameToTeleport := c.GetBlocksToTeleport(selectedEnvId)

	if len(blocksToForward) != 0 {
		utils.StopSpinner()
		c.api.StartTeleport(blocksToForward, selectedEnvId, clusterId, blockNameToTeleport)

	} else {
		utils.WarningMessage("No service/s found in this environment to teleport into!")
	}

}
