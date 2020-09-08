package utils

import (
	"fmt"
	"github.com/Terry-Mao/goconf"
	"github.com/golang/protobuf/ptypes"
	"github.com/kintohub/kinto-cli/internal/config"
	"github.com/kintohub/kinto-cli/internal/types"
	"github.com/rs/zerolog/log"
	"net"
	"os"
	"os/signal"
	"strconv"
	"strings"
	"syscall"
)

//Gets the latest successful release from any service
func GetLatestSuccessfulRelease(releases map[string]*types.Release) *types.Release {
	if releases == nil || len(releases) == 0 {
		return nil
	}

	var latestRelease *types.Release
	for _, release := range releases {
		if release.Status.State == types.Status_SUCCESS {
			if latestRelease == nil {
				latestRelease = release
				continue
			}

			latestCreatedAt, err := ptypes.Timestamp(latestRelease.CreatedAt)

			if err != nil {
				log.Error().Err(err).Msgf(
					"cannot parse timestamp %v to time for release %v",
					latestRelease.CreatedAt, latestRelease)
				continue
			}

			releaseCreateAt, err := ptypes.Timestamp(release.CreatedAt)

			if err != nil {
				log.Error().Err(err).Msgf(
					"cannot parse timestamp %v to time for release %v",
					release.CreatedAt, release)
				continue
			}

			if releaseCreateAt.After(latestCreatedAt) {
				latestRelease = release
			}
		}
	}

	return latestRelease
}

//CheckPort takes a port number and checks if its available.
//If available, will return the port as it is. If not, it will terminate the CLI with error.
func CheckPort(port int) int {

	address := fmt.Sprintf(":%d", port)
	connection, err := net.Listen("tcp", address)
	if err != nil {
		TerminateWithCustomError(
			fmt.Sprintf("Port %d is already in use. Please free it first!", port))
	} else {
		_ = connection.Close()
	}
	return port
}

// Check if Local Git Repo exists
func CheckLocalGitOrDie() {

	conf := goconf.New()
	err := conf.Parse("./.git/config")
	if err != nil {
		TerminateWithCustomError("Not a Git Repo. Please initialize the repo with Git first")
	}

}

//Compare passed URL with local git repo url
func CompareGitUrl(remoteGitUrl string) bool {
	conf := goconf.New()
	_ = conf.Parse("./.git/config")
	remote := conf.Get("remote \"origin\"")
	if remote == nil {
		// In case if git ever changes their config structure.
		TerminateWithCustomError("Cannot parse Git config")
	}
	localGitUrl, err := remote.String("url")
	if err != nil {
		TerminateWithError(err)
	}
	localGitUrl = strings.Trim(localGitUrl, "= ")

	if strings.Replace(remoteGitUrl, ".git", "", -1) == strings.Replace(localGitUrl, ".git", "", -1) {
		return true
	}
	return false
}

//Set default ports for services that are to be passed to chisel.
//Special ports are specified for catalogs since the ports for them are not in buildconfig.
//if the service has either of the given names occurring in the service name, it will return the specified port
//otherwise will fetch the port from buildconfig and return it.
func GetBlockPort(block *types.Block) int {

	if strings.Contains(block.Name, "redis") {
		return config.RedisPort
	} else if strings.Contains(block.Name, "postgres") {
		return config.PostgresPort
	} else if strings.Contains(block.Name, "mongodb") {
		return config.MongoPort
	} else if strings.Contains(block.Name, "minio") {
		return config.MinioPort
	} else if strings.Contains(block.Name, "mysql") {
		return config.MysqlPort
	} else {
		resp := GetLatestSuccessfulRelease(block.Releases).RunConfig.Port
		port, err := strconv.Atoi(resp)
		if err != nil {
			TerminateWithError(err)
		}
		return port
	}
}

//checks if user is logged in or not.
func CheckLogin() {
	email := config.GetEmail()
	token := config.GetAuthToken()

	if email == "" || token == "" {
		TerminateWithCustomError("Please log-in into your account first")
	}
}

func CloseCli() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		TerminateWithCustomError("Aborted!")
	}()
}
