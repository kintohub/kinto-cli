package config

import (
	"fmt"
	"github.com/kintohub/kinto-cli-go/internal/utils"
	enterpriseTypes "github.com/kintohub/kinto-enterprise/pkg/types"
	"github.com/spf13/viper"
)

const (
	Version           = "v0.1"
	authTokenKey      = "authToken"
	emailKey          = "emailKey"
	publicClustersKey = "publicClusters"
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
	viper.Set(authTokenKey, token)
}

func GetAuthToken() string {
	return viper.GetString(authTokenKey)
}

func SetEmail(email string) {
	viper.Set(emailKey, email)
}

func SetPublicClusters(publicClusters []*enterpriseTypes.PublicClusterInfo) {
	m := map[string]interface{}{}

	for _, publicCluster := range publicClusters {
		m[publicCluster.Id] = publicCluster
	}

	viper.Set(publicClustersKey, publicClusters)
}

func GetPublicClusterInfo(clusterId string) *enterpriseTypes.PublicClusterInfo {
	v := viper.GetStringMap(publicClustersKey)

	if v != nil {
		if publicCluster, ok := v[clusterId]; ok {
			return publicCluster.(*enterpriseTypes.PublicClusterInfo)
		}
	}

	return nil
}
