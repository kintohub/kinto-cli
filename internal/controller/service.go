package controller

import (
	"github.com/kintohub/kinto-cli-go/internal/utils"
	"github.com/olekukonko/tablewriter"
	"os"
)

func (c *Controller) Services(envId string) {
	utils.StartSpinner()
	utils.CheckLogin()
	blocks, err := c.api.GetBlocks(envId)

	if err != nil {
		utils.TerminateWithError(err)
	}

	if len(blocks) != 0 {
		table := tablewriter.NewWriter(os.Stdout)
		table.SetHeader([]string{
			"Service Id",
			"Name",
		})

		for _, block := range blocks {
			table.Append([]string{
				block.Id,
				block.DisplayName,
			})
		}
		utils.StopSpinner()
		table.Render()
	} else {
		utils.WarningMessage("No services/s found!")
	}
}
