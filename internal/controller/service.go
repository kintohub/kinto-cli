package controller

import (
	"github.com/kintohub/kinto-cli-go/internal/utils"
	"github.com/olekukonko/tablewriter"
	"os"
)

func (c *Controller) Services(envId string) {
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

		table.Render()
	} else {
		utils.WarningMessage("No services/s found!")
	}
}
