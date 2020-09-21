package controller

import (
	"fmt"
	"github.com/kintohub/kinto-cli/internal/api"
	"github.com/kintohub/kinto-cli/internal/config"
	"github.com/kintohub/kinto-cli/internal/types"
	"github.com/kintohub/kinto-cli/internal/utils"
)

func (c *Controller) GetEnvFromId(envId string) *types.ClusterEnvironment {

	envs, err := c.api.GetClusterEnvironments()
	if err != nil {
		utils.TerminateWithError(err)
	}

	for _, env := range envs {

		if env.Id == envId {
			return env
		}
	}
	return nil
}

func (c *Controller) GetBlockFromId(blockId string) *types.Block {

	envs, err := c.api.GetClusterEnvironments()
	if err != nil {
		utils.TerminateWithError(err)
	}

	for _, env := range envs {
		blocks, err := c.api.GetBlocks(env.Id)
		if err != nil {
			utils.TerminateWithError(err)
		}

		for _, block := range blocks {

			if block.Id == blockId {
				return block
			}
		}
	}
	return nil
}

//Get All teleport-able/Accessible Environment names
func (c *Controller) GetAvailableEnvNames(enableLocalGit bool) []api.EnvDetails {

	if enableLocalGit {
		utils.CheckLocalGitOrDie()
	}

	var envDetails []api.EnvDetails
	serialNumber := 1

	envs, err := c.api.GetClusterEnvironments()
	if err != nil {
		utils.TerminateWithError(err)
	}

	for _, env := range envs {
		blocks, err := c.api.GetBlocks(env.Id)
		if err != nil {
			utils.TerminateWithError(err)
		}

		for _, block := range blocks {
			latestRelease := utils.GetLatestSuccessfulRelease(block.Releases)

			if latestRelease != nil && enableLocalGit &&
				utils.CompareGitUrl(latestRelease.BuildConfig.Repository.Url) {

				envDetails = append(envDetails, api.EnvDetails{EnvName: fmt.Sprintf(
					"%d. %s", serialNumber, env.Name), EnvId: env.Id, ClusterId: env.ClusterId})
				serialNumber++
				break

			} else if latestRelease != nil && !enableLocalGit {

				envDetails = append(envDetails, api.EnvDetails{EnvName: fmt.Sprintf(
					"%d. %s", serialNumber, env.Name), EnvId: env.Id, ClusterId: env.ClusterId})
				serialNumber++
				break
			}
		}

	}

	if len(envDetails) != 0 {
		return envDetails
	} else {
		utils.WarningMessage("No Accessible environment/s found!")
	}

	return nil
}

func (c *Controller) GetBlocksToForward(envId string) []api.RemoteConfig {
	var blocksToForward []api.RemoteConfig
	count := 0

	blocks, err := c.api.GetBlocks(envId)
	if err != nil {
		utils.TerminateWithError(err)
	}

	for _, block := range blocks {
		latestRelease := utils.GetLatestSuccessfulRelease(block.Releases)

		if latestRelease != nil {
			port := config.LocalPort + count
			remote := api.RemoteConfig{FromHost: "localhost", FromPort: utils.CheckPort(port),
				ToHost: block.Name, ToPort: utils.GetBlockPort(block)}
			blocksToForward = append(blocksToForward, remote)
			count += 1
			fmt.Println(block.Name)
		}
	}

	return blocksToForward
}

func (c *Controller) GetBlocksToTeleport(envId string) ([]api.RemoteConfig, string) {
	utils.CheckLocalGitOrDie()
	var blocksToForward []api.RemoteConfig
	var blockName string
	count := 0

	blocks, err := c.api.GetBlocks(envId)
	if err != nil {
		utils.TerminateWithError(err)
	}

	for _, block := range blocks {
		latestRelease := utils.GetLatestSuccessfulRelease(block.Releases)

		if latestRelease != nil {

			if utils.CompareGitUrl(latestRelease.BuildConfig.Repository.Url) {
				remote := api.RemoteConfig{FromHost: "R:localhost",
					FromPort: 3000,
					ToHost:   "localhost", ToPort: 3000}
				blocksToForward = append(blocksToForward, remote)
				count++
				blockName = block.Name
			} else {
				remote := api.RemoteConfig{FromHost: "localhost",
					FromPort: utils.CheckPort(config.LocalPort + count),
					ToHost:   block.Name, ToPort: utils.GetBlockPort(block)}
				blocksToForward = append(blocksToForward, remote)
				count++
			}

		}
	}

	return blocksToForward, blockName
}
