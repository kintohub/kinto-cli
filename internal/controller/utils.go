package controller

import (
	"fmt"
	"github.com/kintohub/kinto-cli/internal/api"
	"github.com/kintohub/kinto-cli/internal/config"
	"github.com/kintohub/kinto-cli/internal/types"
	"github.com/kintohub/kinto-cli/internal/utils"
)

//get env from its id
func (c *Controller) GetEnvWithId(envId string) *types.ClusterEnvironment {

	envs, err := c.api.GetClusterEnvironments()
	if err != nil {
		utils.TerminateWithError(err)
		return nil
	}

	for _, env := range envs {

		if env.Id == envId {
			return env
		}
	}
	return nil
}

//Get All teleport-able/Accessible Environment names
func (c *Controller) GetAvailableEnvNames(enableLocalGitCheck bool) []api.EnvDetails {

	if enableLocalGitCheck {
		utils.CheckLocalGitOrDie()
	}

	var envDetails []api.EnvDetails
	serialNumber := 1

	envs, err := c.api.GetClusterEnvironments()
	if err != nil {
		utils.TerminateWithError(err)
		return nil
	}

	for _, env := range envs {
		blocks, err := c.api.GetBlocks(env.Id)
		if err != nil {
			utils.TerminateWithError(err)
			return nil
		}

		for _, block := range blocks {
			latestRelease := utils.GetLatestSuccessfulRelease(block.Releases)

			if latestRelease != nil && enableLocalGitCheck &&
				utils.CompareGitUrl(latestRelease.BuildConfig.Repository.Url) {
				envDetails = append(
					envDetails,
					api.EnvDetails{
						EnvName:   fmt.Sprintf("%d. %s", serialNumber, env.Name),
						EnvId:     env.Id,
						ClusterId: env.ClusterId})
				serialNumber++
				break

			} else if latestRelease != nil && !enableLocalGitCheck {
				envDetails = append(
					envDetails,
					api.EnvDetails{
						EnvName:   fmt.Sprintf("%d. %s", serialNumber, env.Name),
						EnvId:     env.Id,
						ClusterId: env.ClusterId})
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

//get list of services to port-forward
func (c *Controller) GetBlocksToForward(envId string) []api.RemoteConfig {
	var blocksToForward []api.RemoteConfig
	count := 0

	blocks, err := c.api.GetBlocks(envId)
	if err != nil {
		utils.TerminateWithError(err)
		return nil
	}

	for _, block := range blocks {
		latestRelease := utils.GetLatestSuccessfulRelease(block.Releases)

		if latestRelease != nil && utils.CanPortForwardToRelease(latestRelease) {
			port := config.DefaultClientAccessPort + count

			if utils.CheckIfPortOpen(port, true) {
				remote := api.RemoteConfig{
					FromHost: "localhost",
					FromPort: port,
					ToHost:   block.Name,
					ToPort:   utils.GetBlockPort(block.Name, latestRelease),
				}
				blocksToForward = append(blocksToForward, remote)
				count += 1
			}

		}
	}

	return blocksToForward
}

//get list of services to teleport
func (c *Controller) GetBlocksToTeleport(envId string) ([]api.RemoteConfig, string) {
	utils.CheckLocalGitOrDie()
	var blocksToForward []api.RemoteConfig
	var blockNameToTeleport string
	count := 0

	blocks, err := c.api.GetBlocks(envId)
	if err != nil {
		utils.TerminateWithError(err)
		return nil, ""
	}

	for _, block := range blocks {
		latestRelease := utils.GetLatestSuccessfulRelease(block.Releases)

		if latestRelease != nil && utils.CanPortForwardToRelease(latestRelease) {

			if utils.CompareGitUrl(latestRelease.BuildConfig.Repository.Url) {
				remote := api.RemoteConfig{
					FromHost: config.DefaultTeleportInterfacePort, // server listen to all interfaces
					FromPort: 3000,
					// https://github.com/kintohub/kinto-kube-core/blob/master/internal/store/kube/chisel.go#L35
					ToHost: "localhost",
					ToPort: config.DefaultClientTeleportPort,
					// TODO make it configurable, the user must run their local service on port 8080
				}
				blocksToForward = append(blocksToForward, remote)
				count++
				blockNameToTeleport = block.Name
			} else {
				port := config.DefaultClientAccessPort + count
				if utils.CheckIfPortOpen(port, true) {
					remote := api.RemoteConfig{
						FromHost: "localhost",
						FromPort: port,
						ToHost:   block.Name,
						ToPort:   utils.GetBlockPort(block.Name, latestRelease),
					}
					blocksToForward = append(blocksToForward, remote)
					count++
				}

			}
		}
	}

	return blocksToForward, blockNameToTeleport
}
