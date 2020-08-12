package controller

import (
	"fmt"
	"github.com/kintohub/kinto-cli-go/internal/config"
	"github.com/kintohub/kinto-cli-go/internal/utils"
	"strings"
)

func (c *Controller) Teleport() {

	utils.StartSpinner()

	localGitUrl := utils.GetLocalGitUrl()
	envs, err := c.api.GetClusterEnvironments()

	if err != nil {
		utils.TerminateWithError(err)
	} else {

		var envName []string
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
				if latestRelease.BuildConfig.Repository.Url == localGitUrl {
					envName = append(envName, env.Name)
					envDetails[env.Name] = env.Id
				}
			}
		}

		/* Need to have a Slice "envName" along with envDetails as a parameter
		due to the limitations of the external package used. */
		if len(envDetails) != 0 {
			utils.StopSpinner()
			selectedEnvId := utils.TeleportPrompt(envName, envDetails)
			c.startTunnel(selectedEnvId)

		} else {
			utils.WarningMessage("No environment/s found to port-forward!")
		}
	}
}

func (c *Controller) startTunnel(envId string) {
	utils.StartSpinner()


	 var remotes []string

	inc := 0
	localGitUrl := utils.GetLocalGitUrl()
	blocks, err := c.api.GetBlocks(envId)
	if err != nil {
		utils.TerminateWithError(err)
	}

	/* A second filtering using localGitUrl is needed to filter out the svs inside an Env as
	a previously filtered Env can still have svs that don't belong to the local Git Repo */
	for _, block := range blocks {
		latestRelease := utils.GetLatestSuccessfulRelease(block.Releases)

		if latestRelease.BuildConfig.Repository.Url == localGitUrl {
			port := config.LocalPort + inc

			remote := fmt.Sprintf("%s:%d:%s:%s", "localhost",
				utils.CheckPort(port), block.Name, "80")
			remotes = append(remotes, remote)
			inc += 1
		}
	}
	if len(remotes) != 0 {

		utils.StopSpinner()
		chiselClient := c.api.CreateTeleport(remotes)

		utils.WarningMessage("\nStarting Tunnel")

		for _, proxy := range remotes {
			split := strings.Split(proxy, ":")
			fromHost, fromPort, toBlock, toPort := split[0], split[1], split[2], split[3]
			utils.InfoMessage(fmt.Sprintf("> Forwarding: %s:%s => %s:%s",
				fromHost, fromPort, toBlock, toPort))
		}

		go chiselClient.Run()
		utils.SuccessMessage("✓ Connected!")
		defer chiselClient.Close()
		utils.WarningMessage("\nPress any key to close the tunnel")
		fmt.Scanln()

	} else {
		utils.StopSpinner()
		utils.WarningMessage("No service/s found in this environment to port-forward!")
	}

}
