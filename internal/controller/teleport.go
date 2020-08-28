package controller

import (
	"github.com/kintohub/kinto-cli/internal/api"
	"github.com/kintohub/kinto-cli/internal/config"
	"github.com/kintohub/kinto-cli/internal/utils"
)

func (c *Controller) Teleport() {
	utils.CheckLogin()
	utils.StartSpinner()

	utils.GetGitDetails()
	envs, err := c.api.GetClusterEnvironments()

	if err != nil {
		utils.TerminateWithError(err)
	}

	var envName []string
	var clusterId string
	envDetails := make(map[string]string)
	for _, env := range envs {

		blocks, err := c.api.GetBlocks(env.Id)
		if err != nil {
			utils.TerminateWithError(err)
		}

		for _, block := range blocks {
			latestRelease := utils.GetLatestSuccessfulRelease(block.Releases)

			/* Initial Env filter is done on basis of localGitUrl to present
			a list of Env that have any Svs associated with the local Git Repo. */

			if latestRelease != nil {
				if utils.GetGitDetails(latestRelease.BuildConfig.Repository.Url) {
					envName = append(envName, env.Name)
					envDetails[env.Name] = env.Id
					clusterId = env.ClusterId
				}
			}
		}
	}

	/* Need to have a Slice "envName" along with envDetails as a parameter
	due to the limitations of the external package used. */
	if len(envDetails) != 0 {
		utils.StopSpinner()
		selectedEnvId := TeleportPrompt(envName, envDetails)
		c.configureTeleport(selectedEnvId, clusterId)

	} else {
		utils.WarningMessage("No environment/s found to teleport into!")
	}

}

func (c *Controller) configureTeleport(envId string, clusterId string) {
	utils.StartSpinner()
	var blocksToForward []api.RemoteConfig
	inc := 0
	utils.GetGitDetails()
	blocks, err := c.api.GetBlocks(envId)
	if err != nil {
		utils.TerminateWithError(err)
	}

	for _, block := range blocks {
		latestRelease := utils.GetLatestSuccessfulRelease(block.Releases)

		if latestRelease != nil {

			if !utils.GetGitDetails(latestRelease.BuildConfig.Repository.Url) {
				port := config.LocalPort + inc
				remote := api.RemoteConfig{FromHost: "localhost", FromPort: utils.CheckPort(port),
					ToHost: block.Name, ToPort: utils.GetBlockPort(block)}
				blocksToForward = append(blocksToForward, remote)
				inc += 1
			}
		}
	}

	if len(blocksToForward) != 0 {

		utils.StopSpinner()
		c.api.StartTeleport(blocksToForward, envId, clusterId)

	} else {
		utils.WarningMessage("No service/s found in this environment to teleport into!")
	}

}
