package config

import (
	"fmt"
	"github.com/kintohub/kinto-cli-go/internal/utils"
	"github.com/spf13/viper"
)

const (
	Version = "v0.1"
	KintoMasterHost = "master.vegeta.kintohub.net:443"
)



func AddConfigPath(path string) {
	viper.AddConfigPath(path)
}

func SetConfigName(configName string) {
	viper.SetConfigName(configName)
}

func SetConfigType(configType string) {
	viper.SetConfigType(configType)
}

func AutomaticEnv() {
	viper.AutomaticEnv()
}

func CreateConfig(path string, configName string) {
	if err := viper.ReadInConfig(); err != nil {
		// Create new config file
		err := viper.WriteConfigAs(fmt.Sprintf("%s/%s",
			path,
			configName,
		))
		if err != nil {
			utils.TerminateWithError(err)
		}
	}
}

func WriteConfig() {
	err := viper.WriteConfig()
	if err != nil {
		utils.TerminateWithError(err)
	}
}

func GetString(string string) string {
	str := viper.GetString(string)
	return str
}

func SetAuthToken(token string) {
	viper.Set("authToken", token)
}

func SetEmail(email string) {
	viper.Set("email", email)
}
