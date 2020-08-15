package cli

import (
	"errors"
	"fmt"
	"github.com/kintohub/kinto-cli-go/internal/config"
	"github.com/kintohub/kinto-cli-go/internal/controller"
	"github.com/kintohub/kinto-cli-go/internal/utils"
	"github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"os"
)

type CliInterface interface {
	GetHostFlag() string
	Execute(controller controller.ControllerInterface)
}

type Cli struct {
	rootCmd *cobra.Command
}

func NewCliOrDie() CliInterface {
	cobra.OnInitialize(initConfig)

	var rootCmd = &cobra.Command{
		Use:   "kinto",
		Short: "Kinto helps developers ship and iterate full stack apps with ease",
		Long: `KintoHub comes with a complete suite of tools to build, deploy, debug and optimize apps.
               Documentation is available at https://docs.kintohub.com`,
	}

	rootCmd.PersistentFlags().StringP(
		"host", "", "master.vegeta.kintohub.net:443", "target kintohub host")

	return &Cli{
		rootCmd: rootCmd,
	}
}

func (c *Cli) GetHostFlag() string {
	host := c.rootCmd.PersistentFlags().Lookup("host")

	if host == nil {
		utils.TerminateWithError(errors.New("internal error - host flag was not setup correcting"))
	}

	return host.Value.String()
}

func initConfig() {
	// Find home directory.
	home, err := homedir.Dir()
	if err != nil {
		fmt.Printf("could not find home directory with error %v", err)
		os.Exit(1)
	}

	// Search config in home directory with name "kinto.yaml"
	const configName = "kinto.yaml"
	config.AddConfigPath(home)
	config.SetConfigName(configName)
	config.SetConfigType("yaml")
	config.AutomaticEnv()
	config.CreateConfig(home, configName)
}

func (c *Cli) Execute(controller controller.ControllerInterface) {
	c.rootCmd.AddCommand(
		createVersionCommand(controller),
		createLoginCommand(controller),
		createEnvironmentCommand(controller),
		createServicesCommand(controller),
		createTeleportCommand(controller),
		createStatusCommand(controller),
	)

	if err := c.rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func createLoginCommand(controller controller.ControllerInterface) *cobra.Command {
	return &cobra.Command{
		Use:   "login",
		Short: "Login to an existing KintoHub account",
		Long:  `Helps you to log in an already existing KintoHub account`,
		Run: func(cmd *cobra.Command, args []string) {
			controller.Login()
		},
	}
}

func createVersionCommand(controller controller.ControllerInterface) *cobra.Command {
	return &cobra.Command{
		Use:   "version",
		Short: "Prints the version number of Kinto CLI",
		Long:  `All software has versions. This is Kinto's!`,
		Run: func(cmd *cobra.Command, args []string) {
			controller.Version()
		},
	}
}

func createEnvironmentCommand(controller controller.ControllerInterface) *cobra.Command {
	return &cobra.Command{
		Use:   "env",
		Short: "List all the Environment IDs and their regions",
		Long:  `Get a list of all the Environment ID names and their regions`,
		Run: func(cmd *cobra.Command, args []string) {
			controller.Environment()
		},
	}
}

func createServicesCommand(controller controller.ControllerInterface) *cobra.Command {
	return &cobra.Command{
		Use:   "svs",
		Short: "List services",
		Long:  `Get a list of all services within an environment`,
		Args: func(cmd *cobra.Command, args []string) error {
			if len(args) != 1 {
				return errors.New("requires envId argument")
			}

			return nil
		},
		Run: func(cmd *cobra.Command, args []string) {
			controller.Services(args[0])
		},
	}
}

func createTeleportCommand(controller controller.ControllerInterface) *cobra.Command {
	return &cobra.Command{
		Use:   "teleport",
		Short: "Port-forward your remote services to your local machine",
		Run: func(cmd *cobra.Command, args []string) {
			controller.Teleport()
		},
	}
}

func createStatusCommand(controller controller.ControllerInterface) *cobra.Command {
	return &cobra.Command{
		Use: "status",
		Short: `List environments & services where the current repo is deployed. 
				This commands needs to be called from within a Git repo.`,
		Long: `Get a list of all environments & services where the current Git repo is deployed to. 
				This command should be run from within a Git repo.`,
		Run: func(cmd *cobra.Command, args []string) {
			controller.Status()
		},
	}
}
