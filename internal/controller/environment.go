package controller

import (
	"github.com/docker/distribution/context"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/gookit/color"
	"github.com/kintohub/kinto-cli-go/internal/config"
	"github.com/kintohub/kinto-cli-go/internal/utils"
	"github.com/olekukonko/tablewriter"
	"os"

	_ "github.com/olekukonko/tablewriter"
	"google.golang.org/grpc/metadata"
)

func (c *Controller) Environment() {

	bearer := "Bearer " + config.GetString("authToken")
	md := metadata.Pairs("Authorization", bearer)
	ctx := metadata.NewOutgoingContext(context.Background(), md)

	env, err := c.clustersClient.GetClusterEnvironments(ctx, &empty.Empty{})
	clusters, err := c.clustersClient.GetClusters(ctx, &empty.Empty{})

	if err != nil {
		utils.TerminateWithError(err)
	} else {

		clusterDetail := make(map[string]string)

		for _, c := range clusters.Clusters {
			clusterDetail[c.Id] = c.DisplayName
		}
		table := tablewriter.NewWriter(os.Stdout)
		table.SetHeader([]string{
			"Env Id",
			"Name",
			"Region",
		})
		for _, c := range env.Envs {
			table.Append([]string{
				c.Id,
				c.Name,
				clusterDetail[c.ClusterId],
			})
		}

		if len(env.GetEnvs()) != 0 {
			table.Render()
		} else {
			color.Red.Printf("\nNo environment/s found!\n")
		}

	}

}
