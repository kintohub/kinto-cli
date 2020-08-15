package controller

import (
	"github.com/kintohub/kinto-cli-go/internal/utils"
	"github.com/olekukonko/tablewriter"
	"os"
)

func (c *Controller) Environment() {

	utils.LoginCheck()

	utils.StartSpinner()
	envs, err := c.api.GetClusterEnvironments()
	clusters, err := c.api.GetClusters()

	if err != nil {
		utils.TerminateWithError(err)
	}

	clusterDetail := make(map[string]string)

	for _, c := range clusters {
		clusterDetail[c.Id] = c.DisplayName
	}
	table := tablewriter.NewWriter(os.Stdout)
	table.SetRowLine(true)
	table.SetHeader([]string{
		"Env Id",
		"Name",
		"Region",
	})
	table.SetHeaderColor(tablewriter.Colors{tablewriter.Bold, tablewriter.FgWhiteColor},
		tablewriter.Colors{tablewriter.Bold, tablewriter.FgWhiteColor},
		tablewriter.Colors{tablewriter.Bold, tablewriter.FgWhiteColor})

	for _, c := range envs {
		table.Append([]string{
			c.Id,
			c.Name,
			clusterDetail[c.ClusterId],
		})
	}

	if len(envs) != 0 {
		utils.SuccessMessage("Available environments:")
		table.Render()
	} else {
		utils.WarningMessage("No environment/s found!")
	}

}
