package cli

import (
	"fmt"
	"github.com/kintohub/kinto-cli-go/internal/config"
	"github.com/kintohub/kinto-cli-go/internal/controller"
	"github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
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
		createRegisterCommand(controller),
		createLoginCommand(controller),
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

	// Search config in home directory with name "kinto.yaml"
	const configName = "kinto.yaml"
	config.AddConfigPath(home)
	config.SetConfigName(configName)
	config.SetConfigType("yaml")
	config.AutomaticEnv()
	config.CreateConfig(home,configName)
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
		Long:  `Helps create a new KintoHub account`,
		Run: func(cmd *cobra.Command, args []string) {
			controller.Init()
		},
	}
}

func createRegisterCommand(controller controller.ControllerInterface) *cobra.Command {
	return &cobra.Command{
		Use:   "register",
		Short: "Creates a new account on KintoHub",
		Long:  `Helps create a new kintoHub account`,
		Run: func(cmd *cobra.Command, args []string) {
			controller.Register()
		},
	}
}

func createLoginCommand(controller controller.ControllerInterface) *cobra.Command {
	return &cobra.Command{
		Use:   "login",
		Short: "Log in an existing KintoHub account",
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


