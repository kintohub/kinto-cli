package controller

import (
	"fmt"
	"github.com/kintohub/kinto-cli/internal/api"
	"github.com/kintohub/kinto-cli/internal/utils"
	"github.com/olekukonko/tablewriter"
	"os"
)

func createTable() *tablewriter.Table {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetRowLine(true)
	table.SetHeader([]string{
		"Name",
		"Service ID",
	})
	table.SetHeaderColor(tablewriter.Colors{tablewriter.Bold, tablewriter.FgWhiteColor},
		tablewriter.Colors{tablewriter.Bold, tablewriter.FgWhiteColor})

	return table
}

func (c *Controller) Services(envId ...string) {

	utils.StartSpinner()
	utils.CheckLogin()
	table := createTable()

	var envName []string
	var envDetails []api.EnvDetails

	if len(envId) != 0 {
		blocks, err := c.api.GetBlocks(envId[0])

		if err != nil {
			utils.TerminateWithError(err)
		}

		if len(blocks) != 0 {
			for _, block := range blocks {
				table.Append([]string{
					block.DisplayName,
					block.Id,
				})
			}
			utils.StopSpinner()
			table.Render()
		} else {
			utils.WarningMessage("No services/s found!")
		}

	} else {

		count := 1
		envs, err := c.api.GetClusterEnvironments()
		if err != nil {
			utils.TerminateWithError(err)
		}

		for _, env := range envs {
			envName = append(envName, fmt.Sprintf("%d. %s", count, env.Name))
			envDetail := api.EnvDetails{EnvName: env.Name, EnvId: env.Id}
			envDetails = append(envDetails, envDetail)
			count += 1
		}

		if len(envDetails) != 0 {
			utils.StopSpinner()
			selectedEnvId := SelectionPrompt(envName, envDetails)
			c.showSelectedEnvServices(selectedEnvId)

		} else {
			utils.WarningMessage("No Env/s found!")
		}
	}
}

func (c *Controller) showSelectedEnvServices(envId string) {

	utils.StartSpinner()
	blocks, err := c.api.GetBlocks(envId)
	table := createTable()
	if err != nil {
		utils.TerminateWithError(err)
	}

	if len(blocks) != 0 {
		for _, block := range blocks {
			table.Append([]string{
				block.DisplayName,
				block.Id,
			})
		}
		utils.StopSpinner()
		table.Render()
	} else {
		utils.WarningMessage("No services/s found!")
	}

}
