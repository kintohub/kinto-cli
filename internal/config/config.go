package config

import (
	"fmt"
	"github.com/gookit/color"
	enterpriseTypes "github.com/kintohub/kinto-enterprise/pkg/types"
	"github.com/spf13/viper"
	"os"
)

const (
	authTokenKey           = "authToken"
	emailKey               = "emailKey"
	publicClustersKey      = "publicClusters"
	clusterEnvironmentsKey = "clusterEnvironments"
	LocalPort              = 5360
	ChiselHost             = "https://chisel-5f194.vegeta.kintohub.net"
)

var Version = "v0.1" //Needs to be a non-const for passing version at build time

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
			color.Red.Println("An error occurred: %v", err)
			//To avoid cyclic dependency error while importing
			os.Exit(1)
		}
	}
}

func Save() {
	err := viper.WriteConfig()
	if err != nil {
		color.Red.Println("An error occurred: %v", err)
		os.Exit(1)
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

func GetEmail() string {
	return viper.GetString(emailKey)
}

func SetPublicClusters(publicClusters []*enterpriseTypes.PublicClusterInfo) {
	publicClustersMap := map[string]interface{}{}

	for _, publicCluster := range publicClusters {
		publicClustersMap[publicCluster.Id] = publicCluster
	}

	viper.Set(publicClustersKey, publicClustersMap)
}

func GetPublicClusterInfo(clusterId string) *enterpriseTypes.PublicClusterInfo {
	publicClustersMap := viper.GetStringMap(publicClustersKey)

	if publicClustersMap != nil {
		if publicCluster, ok := publicClustersMap[clusterId]; ok {
			return publicCluster.(*enterpriseTypes.PublicClusterInfo)
		}
	}

	return nil
}

func SetClusterEnvironments(clusterEnvs []*enterpriseTypes.ClusterEnvironment) {
	clusterEnvironmentsMap := map[string]interface{}{}

	for _, clusterEnv := range clusterEnvs {
		clusterEnvironmentsMap[clusterEnv.Id] = clusterEnv
	}

	viper.Set(publicClustersKey, clusterEnvironmentsMap)
}

func GetClusterEnvironment(envId string) *enterpriseTypes.ClusterEnvironment {
	clusterEnvironmentsMap := viper.GetStringMap(clusterEnvironmentsKey)

	if clusterEnvironmentsMap != nil {
		if clusterEnv, ok := clusterEnvironmentsMap[envId]; ok {
			return clusterEnv.(*enterpriseTypes.ClusterEnvironment)
		}
	}

	return nil
}
