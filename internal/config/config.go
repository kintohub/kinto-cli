package config

import "github.com/spf13/viper"

const (
	Version         = "v0.1"
	KintoMasterHost = "master.vegeta.kintohub.net:443"
)

func SetAuthToken(token string)  {
	viper.Set("authToken", token)
}

func SetEmail(email string)  {
	viper.Set("email", email)
}