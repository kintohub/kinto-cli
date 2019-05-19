package cmd

import (
	"errors"
	"github.com/spf13/cobra"
	"kinto-cli-go/pkg/kintohub"
	"strings"
)

var ProxyCmd = &cobra.Command{
	Use:     "proxy",
	Example: "proxy {block-name}:{port}",
	Short:   "Proxy a specific kintoblock's traffic a local port",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) == 1 {
			splitStr := strings.Split(args[0], ":")
			if len(splitStr) != 2 {
				return errors.New("Invalid proxy argument. Must provide argument {block-name}:{port}")
			}
		} else {
			return errors.New("Did not provide a proxy target and destination {block-name}:${port}")
		}

		return nil
	},
	Run: HandleProxyAction,
}

func HandleProxyAction(cmd *cobra.Command, args []string) {
	splitStr := strings.Split(args[0], ":")
	blockName := splitStr[0]
	forwardPort := splitStr[1]

	api := kintohub.InitApi("token")
	api.CreateProxy("api.kintohub.com", blockName, forwardPort)
}
