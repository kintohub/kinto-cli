package controller

import (
	"fmt"
	"github.com/kintohub/kinto-cli/internal/api"
	"github.com/kintohub/kinto-cli/internal/config"
	"github.com/kintohub/kinto-cli/internal/utils"
)

//Teleports / Port-Forwards remote services to the local machine.
// Takes a flag as arg.
// if flag is true, teleport will show all the available services inside a user account.
// if false, it will only show the services belonging to the user's current local git repo
func (c *Controller) Teleport(teleportAllFlag bool) {

	utils.CheckLogin()
	utils.StartSpinner()

	var clusterId string
	var envDetails []api.EnvDetails
	teleportableServices := 0
	serialNumber := 1

	if !teleportAllFlag {
		utils.CheckLocalGitOrDie()
	}

	envs, err := c.api.GetClusterEnvironments()
	if err != nil {
		utils.TerminateWithError(err)
	}

	for _, env := range envs {
		teleportableServices = 0
		blocks, err := c.api.GetBlocks(env.Id)
		if err != nil {
			utils.TerminateWithError(err)
		}

		for _, block := range blocks {
			latestRelease := utils.GetLatestSuccessfulRelease(block.Releases)

			//only consider if latest release is "Active/Successful"
			//only increment the teleportableServices counter to keep track of
			//services to avoid duplication of envNames in SelectionPrompt
			if latestRelease != nil &&
				utils.CompareGitUrl(latestRelease.BuildConfig.Repository.Url) && !teleportAllFlag {
				teleportableServices++
			} else if latestRelease != nil && teleportAllFlag {
				teleportableServices++
			}
		}

		if teleportableServices > 0 {
			//serialNumber is appended before each EnvName in SelectionPrompt
			//to make each envId, envName combo unique.
			//so that SelectionPrompt can pass correct values to chisel server.
			envDetails = append(envDetails, api.EnvDetails{EnvName: fmt.Sprintf(
				"%d. %s", serialNumber, env.Name), EnvId: env.Id})
			clusterId = env.ClusterId
			serialNumber++
		}
	}

	//check if the appended struct has a non-zero length and then only pass the values to SelectionPrompt
	if len(envDetails) != 0 {
		utils.StopSpinner()
		selectedEnvId := SelectionPrompt(envDetails)
		c.configureTeleport(selectedEnvId, clusterId)

	} else {
		utils.WarningMessage("No environment/s found to teleport into!")
	}

}

//Initial setup for the data that is to be passed to the chisel server.
func (c *Controller) configureTeleport(envId string, clusterId string) {
	utils.StartSpinner()
	var blocksToForward []api.RemoteConfig
	count := 0
	utils.CheckLocalGitOrDie()
	blocks, err := c.api.GetBlocks(envId)
	if err != nil {
		utils.TerminateWithError(err)
	}

	for _, block := range blocks {
		latestRelease := utils.GetLatestSuccessfulRelease(block.Releases)

		if latestRelease != nil {

			//Here every service inside a selected environment is port forwarded except
			//the one that is present in the user's local Git repo
			if !utils.CompareGitUrl(latestRelease.BuildConfig.Repository.Url) {
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
