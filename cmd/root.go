package cmd

import (
	"fmt"
	"github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
)

var cfgFile string

func init() {
	//cobra.OnInitialize(initConfig)

	rootCmd.AddCommand(ProxyCmd)

	//rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.kinto.yaml)")
	rootCmd.PersistentFlags().StringP("gateway", "g", "gateway", "kinto-gateway hostname")
	viper.BindPFlag("gateway", rootCmd.PersistentFlags().Lookup("gateway"))
	viper.SetDefault("gateway", "api.kintohub.com")
}

var rootCmd = &cobra.Command{
	Use:   "kinto",
	Short: "Kinto cli tools allow you to quickly debug and interact with KintoHub",
}

func initConfig() {
	// Don't forget to read config either from cfgFile or from home directory!
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Search config in home directory with name ".cobra" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".kinto")
	}

	if err := viper.ReadInConfig(); err != nil {
		fmt.Println("Can't read config:", err)
		os.Exit(1)
	} else {
		//TODO: Create default config
	}
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
