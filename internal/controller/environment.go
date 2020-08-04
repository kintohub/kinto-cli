package controller

import (
	"github.com/gookit/color"
	"github.com/kintohub/kinto-cli-go/internal/utils"
	"github.com/olekukonko/tablewriter"
	"os"

	_ "github.com/olekukonko/tablewriter"
)

func (c *Controller) Environment() {
	envs, err := c.api.GetClusterEnvironments()
	clusters, err := c.api.GetClusters()

	if err != nil {
		utils.TerminateWithError(err)
	} else {

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

		for _, c := range envs {
			table.Append([]string{
				c.Id,
				c.Name,
				clusterDetail[c.ClusterId],
			})
		}

		if len(envs) != 0 {
			table.Render()
		} else {
			color.Red.Printf("\nNo environment/s found!\n")
		}

	}
}
