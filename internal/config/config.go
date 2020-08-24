package config

import (
	"fmt"
	"github.com/gookit/color"
	"github.com/kintohub/kinto-cli/internal/types"
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
		//To avoid cyclic dependency error while importing
		if err != nil {
			color.Red.Println("An error occurred: %v", err)
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

func SetPublicClusters(publicClusters []*types.PublicClusterInfo) {
	publicClustersMap := map[string]interface{}{}

	for _, publicCluster := range publicClusters {
		publicClustersMap[publicCluster.Id] = publicCluster
	}

	viper.Set(publicClustersKey, publicClustersMap)
}

func GetPublicClusterInfo(clusterId string) *types.PublicClusterInfo {
	publicClustersMap := viper.GetStringMap(publicClustersKey)

	if publicClustersMap != nil {
		if publicCluster, ok := publicClustersMap[clusterId]; ok {
			return publicCluster.(*types.PublicClusterInfo)
		}
	}

	return nil
}

func SetClusterEnvironments(clusterEnvs []*types.ClusterEnvironment) {
	clusterEnvironmentsMap := map[string]interface{}{}

	for _, clusterEnv := range clusterEnvs {
		clusterEnvironmentsMap[clusterEnv.Id] = clusterEnv
	}

	viper.Set(publicClustersKey, clusterEnvironmentsMap)
}

func GetClusterEnvironment(envId string) *types.ClusterEnvironment {
	clusterEnvironmentsMap := viper.GetStringMap(clusterEnvironmentsKey)

	if clusterEnvironmentsMap != nil {
		if clusterEnv, ok := clusterEnvironmentsMap[envId]; ok {
			return clusterEnv.(*types.ClusterEnvironment)
		}
	}

	return nil
}
