package controller

import (
	"fmt"
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
	var envDetails []api.EnvDetails
	count := 1
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

					envName = append(envName, fmt.Sprintf("%d. %s", count, env.Name))
					envDetail := api.EnvDetails{EnvName: env.Name, EnvId: env.Id}
					envDetails = append(envDetails, envDetail)
					clusterId = env.ClusterId
					count += 1
				}
			}
		}
	}

	/* Need to have a Slice "envName" along with envDetails as a parameter
	due to the limitations of the external package used. */
	if len(envDetails) != 0 {
		utils.StopSpinner()
		selectedEnvId := SelectionPrompt(envName, envDetails)
		c.configureTeleport(selectedEnvId, clusterId)

	} else {
		utils.WarningMessage("No environment/s found to teleport into!")
	}

}

func (c *Controller) configureTeleport(envId string, clusterId string) {
	utils.StartSpinner()
	var blocksToForward []api.RemoteConfig
	count := 0
	utils.GetGitDetails()
	blocks, err := c.api.GetBlocks(envId)
	if err != nil {
		utils.TerminateWithError(err)
	}

	for _, block := range blocks {
		latestRelease := utils.GetLatestSuccessfulRelease(block.Releases)

		if latestRelease != nil {

			if !utils.GetGitDetails(latestRelease.BuildConfig.Repository.Url) {
				port := config.LocalPort + count
				remote := api.RemoteConfig{FromHost: "localhost", FromPort: utils.CheckPort(port),
					ToHost: block.Name, ToPort: utils.GetBlockPort(block)}
				blocksToForward = append(blocksToForward, remote)
				count += 1
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
