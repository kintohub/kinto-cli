package controller

import (
	"github.com/kintohub/kinto-cli/internal/utils"
	"github.com/olekukonko/tablewriter"
	"os"
)

//Get list of all available environments in an account
func (c *Controller) Environment() {

	utils.CheckLogin()
	utils.StartSpinner()

	envs, er := c.api.GetClusterEnvironments()
	if er != nil {
		utils.TerminateWithError(er)
		return
	}

	clusters, err := c.api.GetClusters()
	if err != nil {
		utils.TerminateWithError(err)
		return
	}

	clusterDetail := make(map[string]string)

	for _, c := range clusters {
		clusterDetail[c.Id] = c.DisplayName
	}
	table := tablewriter.NewWriter(os.Stdout)
	table.SetRowLine(true)
	table.SetHeader([]string{
		"Env Id",
		"Cluster Id",
		"Name",
		"Region",
	})
	table.SetHeaderColor(tablewriter.Colors{tablewriter.Bold, tablewriter.FgWhiteColor},
		tablewriter.Colors{tablewriter.Bold, tablewriter.FgWhiteColor},
		tablewriter.Colors{tablewriter.Bold, tablewriter.FgWhiteColor},
		tablewriter.Colors{tablewriter.Bold, tablewriter.FgWhiteColor})

	for _, c := range envs {
		table.Append([]string{
			c.Id,
			c.ClusterId,
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
