package controller

import (
	"github.com/kintohub/kinto-cli/internal/api"
	"github.com/kintohub/kinto-cli/internal/config"
	"github.com/kintohub/kinto-cli/internal/utils"
)

func (c *Controller) Access() {

	utils.CheckLogin()
	utils.StartSpinner()

	var envDetails []api.EnvDetails

	envDetails = c.GetAvailableEnvNames(true)

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

			if utils.CanPortForwardToRelease(latestRelease) {
				remote := api.RemoteConfig{
					FromHost: "localhost",
					FromPort: utils.CheckIfPortAvailable(config.LocalPort),
					ToHost:   block.Name,
					ToPort:   utils.GetBlockPort(block.Name, latestRelease),
				}
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
