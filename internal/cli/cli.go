package cli

import (
	"fmt"
	"github.com/kintohub/kinto-cli/internal/config"
	"github.com/kintohub/kinto-cli/internal/controller"
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
	initConfig()

	var rootCmd = &cobra.Command{
		Use:   "kinto",
		Short: "Kinto helps developers ship and iterate full stack apps with ease",
		Long: `KintoHub comes with a complete suite of tools to build, deploy, debug and optimize apps.
Documentation is available at https://docs.kintohub.com`,
	}

	return &Cli{
		rootCmd: rootCmd,
	}
}

func (c *Cli) GetHostFlag() string {
	host := config.GetMasterHost()
	return host
}

func initConfig() {
	// Find home directory.
	home, err := homedir.Dir()
	if err != nil {
		fmt.Printf("could not find home directory with error %v", err)
		os.Exit(1)
	}

	// Search config in home directory with name "kinto.yaml"
	// If it does not exists create one
	const configName = config.CliConfigName
	config.AddConfigPath(home)
	config.SetConfigName(configName)
	config.SetConfigType("yaml")
	config.AutomaticEnv()
	config.CreateConfig(home, configName)

}

func (c *Cli) Execute(controller controller.ControllerInterface) {
	c.rootCmd.AddCommand(
		createInitCommand(controller),
		createVersionCommand(controller),
		createLoginCommand(controller),
		createEnvironmentCommand(controller),
		createServicesCommand(controller),
		createAccessCommand(controller),
		createTeleportCommand(controller),
		createStatusCommand(controller),
	)

	if err := c.rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
func createInitCommand(controller controller.ControllerInterface) *cobra.Command {
	initCmd := &cobra.Command{
		Use:   "init",
		Short: "Init the CLI",
		Long:  `Create a 'kinto.yaml' file in your home directory`,
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			controller.Init(args[0])
		},
	}
	initCmd.SetUsageTemplate("\nUsage:\nSet new master host:\n\t" +
		"kinto init [host]\n\nReset master host:\n\tkinto init default\n")
	return initCmd
}

func createLoginCommand(controller controller.ControllerInterface) *cobra.Command {
	return &cobra.Command{
		Use:   "login",
		Short: "Login to an existing KintoHub account",
		Long:  `Allows you to log into an already existing KintoHub account`,
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
	envCmd := &cobra.Command{
		Use:     "env",
		Aliases: []string{"envs", "environment", "environments"},
		Short:   "List all the Environment IDs and their regions",
		Long:    `Get a list of all the Environment ID names and their regions`,
		Run: func(cmd *cobra.Command, args []string) {
			controller.Environment()
		},
	}

	envCmd.AddCommand(createEnvironmentAccessCommand(controller))
	return envCmd
}

func createEnvironmentAccessCommand(controller controller.ControllerInterface) *cobra.Command {
	return &cobra.Command{
		Use:     "access",
		Aliases: []string{"envs", "environment", "environments"},
		Short:   "Port-Forward services in an environment to your local machine",
		Long:    `Port-Forward all the services in an environment to your local machine.`,
		Args:    cobra.MaximumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			controller.EnvironmentAccess(args...)
		},
	}
}

func createServicesCommand(controller controller.ControllerInterface) *cobra.Command {
	svsCmd := &cobra.Command{
		Use:     "svs",
		Aliases: []string{"service", "services"},
		Short:   "List your services",
		Long:    `Get a list of all services within an environment`,
		Args:    cobra.MaximumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			controller.Services(args...)
		},
	}

	svsCmd.AddCommand(createServiceAccessCommand(controller))
	return svsCmd
}

func createServiceAccessCommand(controller controller.ControllerInterface) *cobra.Command {
	return &cobra.Command{
		Use:     "access",
		Aliases: []string{"envs", "environment", "environments"},
		Short:   "Port-Forward your services to your local machine",
		Long:    `Port-Forward all services in your environment to your local machine. 
Requires environment ID and service ID. This commands needs to be called from within a Git repo.`,
		Args:    cobra.ExactArgs(2),
		Run: func(cmd *cobra.Command, args []string) {
			controller.ServiceAccess(args[0], args[1])
		},
	}
}

func createAccessCommand(controller controller.ControllerInterface) *cobra.Command {
	accessCmd := &cobra.Command{
		Use:   "access",
		Short: "Port-forward your remote services to your local machine",
		Run: func(cmd *cobra.Command, args []string) {
			controller.Access()
		},
	}
	return accessCmd
}

func createTeleportCommand(controller controller.ControllerInterface) *cobra.Command {
	accessCmd := &cobra.Command{
		Use:   "teleport",
		Short: "Teleport into your remote services",
		Long: `Teleport allows you to teleport your local setup into KintoHub.
The teleported service's traffic will be redirected to your local machine and 
the rest of the services will be port-forwarded.
This commands needs to be called from within a Git repo.`,
		Run: func(cmd *cobra.Command, args []string) {
			controller.Teleport()
		},
	}
	return accessCmd
}

func createStatusCommand(controller controller.ControllerInterface) *cobra.Command {
	return &cobra.Command{
		Use: "status",
		Short: `List the environments on which the current local git repo is deployed.`,
		Long: `Get a list of all the environments on which the current local git repo is deployed.
This command should be run from within a Git repo.`,
		Run: func(cmd *cobra.Command, args []string) {
			controller.Status()
		},
	}
}
