package controller

import (
	"github.com/kintohub/kinto-cli-go/internal/utils"
	"github.com/olekukonko/tablewriter"
	"os"
)

func (c *Controller) Status() {
	utils.CheckLogin()
	utils.StartSpinner()

	localGitUrl := utils.GetLocalGitUrl()
	var count = 0
	envs, err := c.api.GetClusterEnvironments()

	if err != nil {
		utils.TerminateWithError(err)
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetRowLine(true)
	table.SetHeader([]string{
		"Env Name",
		"Service Name",
	})
	table.SetHeaderColor(tablewriter.Colors{tablewriter.Bold, tablewriter.FgWhiteColor},
		tablewriter.Colors{tablewriter.Bold, tablewriter.FgWhiteColor})

	for _, env := range envs {
		blocks, err := c.api.GetBlocks(env.Id)
		if err != nil {
			utils.TerminateWithError(err)
		}
		for _, block := range blocks {
			latestRelease := utils.GetLatestSuccessfulRelease(block.Releases)

			if latestRelease.BuildConfig.Repository.Url == localGitUrl {
				count = count + 1 /* To avoid rendering the table multiple times
				if the repo is deployed more than once on KintoHub. */
				table.Append([]string{
					env.Name,
					block.Name,
				})
			}
		}
	}

	if count > 0 {
		utils.SuccessMessage("Repo is deployed to these environments:")
		table.Render()
	} else {
		utils.WarningMessage("Current Repo is not deployed on KintoHub!")
	}

}
