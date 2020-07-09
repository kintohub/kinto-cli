package cli

import (
	"fmt"
	"github.com/kintohub/kinto-cli-go/internal/controller"
	"github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
)

type CliInterface interface {
	Execute()
}

type Cli struct {
	rootCmd *cobra.Command
}

func NewCliOrDie(controller controller.ControllerInterface) CliInterface {
	cobra.OnInitialize(initConfig)

	var rootCmd = &cobra.Command{
		Use:   "kinto",
		Short: "Kinto helps developers ship and iterate full stack apps with ease",
		Long: `KintoHub comes with a complete suite of tools to build, deploy, debug and optimize apps.
               Documentation is available at https://docs.kintohub.com`,
	}

	rootCmd.AddCommand(
		createVersionCommand(controller),
		createInitCommand(controller),
		createLogoutCommand(controller),
	)

	return &Cli{
		rootCmd: rootCmd,
	}
}

func initConfig() {
	// Find home directory.
	home, err := homedir.Dir()
	if err != nil {
		fmt.Printf("could not find home directory with error %v", err)
		os.Exit(1)
	}

	// Search config in home directory with name ".cobra" (without extension).
	viper.AddConfigPath(home)
	viper.SetConfigName(".kintocli")

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("using config file:", viper.ConfigFileUsed())
	}
}

func (c *Cli) Execute() {
	if err := c.rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func createInitCommand(controller controller.ControllerInterface) *cobra.Command {
	return &cobra.Command{
		Use:   "init",
		Short: "Login or registers to KintoHub",
		Long:  `Helps create a new kintohub account`,
		Run: func(cmd *cobra.Command, args []string) {
			controller.Init()
		},
	}
}

func createLogoutCommand(controller controller.ControllerInterface) *cobra.Command {
	return &cobra.Command{
		Use:   "logout",
		Short: "initializes the git repository for kintohub",
		Long:  `Authenticates to KintoHub and detects or creates a .kinto file`,
		Run: func(cmd *cobra.Command, args []string) {
			controller.Logout()
		},
	}
}

func createVersionCommand(controller controller.ControllerInterface) *cobra.Command {
	return &cobra.Command{
		Use:   "version",
		Short: "Print the version number of Kinto CLI",
		Long:  `All software has versions. This is Kinto's!`,
		Run: func(cmd *cobra.Command, args []string) {
			controller.Version()
		},
	}
}
