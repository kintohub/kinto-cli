package controller

import (
	"github.com/gookit/color"
	"github.com/kintohub/kinto-cli-go/internal/utils"
	"github.com/olekukonko/tablewriter"
	"os"
)

func (c *Controller) Status() {

	localGitUrl := utils.GetLocalGitUrl()
	var count = 0
	envs, err := c.api.GetClusterEnvironments()

	if err != nil {
		utils.TerminateWithError(err)
	} else {

		table := tablewriter.NewWriter(os.Stdout)
		table.SetAutoMergeCells(true)
		table.SetRowLine(true)
		table.SetHeader([]string{
			"Env Name",
			"Service Name",
		})

		for _, env := range envs {
			blocks, err := c.api.GetBlocks(env.Id)
			if err != nil {

			}
			for _, block := range blocks {
				for _, release := range block.Releases {
					if release.BuildConfig.Repository.Url == localGitUrl {
						count = count + 1 /* To avoid rendering the table multiple times
						if the repo is deployed more than once on KintoHub. */
						table.Append([]string{
							env.Name,
							block.Name,
						})
					}
				}
			}
		}

		if count > 0 {
			color.Green.Printf("\nRepo is deployed to these environments:\n")
			table.Render()
		} else {
			color.Yellow.Printf("Current Repo is not deployed on KintoHub!")
		}
	}
}
