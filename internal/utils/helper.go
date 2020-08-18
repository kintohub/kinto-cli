package utils

import (
	"fmt"
	"github.com/Terry-Mao/goconf"
	"github.com/golang/protobuf/ptypes"
	"github.com/kintohub/kinto-cli/internal/config"
	"github.com/kintohub/kinto-cli/internal/types"
	"github.com/rs/zerolog/log"
	"net"
	"strings"
)

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
//If available, will return the port as it is. If not, error will be thrown.
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

func GetLocalGitUrl() string {
	conf := goconf.New()
	err := conf.Parse("./.git/config")

	if err != nil {
		TerminateWithCustomError("Not a Git Repo. Please initialize the repo with Git first")
	}
	remote := conf.Get("remote \"origin\"")

	if remote == nil {
		// In case if git ever changes their config structure.
		TerminateWithError(err)
	}

	localGitUrl, err := remote.String("url")
	if err != nil {
		TerminateWithError(err)
	}

	localGitUrl = strings.Trim(localGitUrl, "= ")
	localGitUrl = strings.Trim(localGitUrl, ".git")

	return localGitUrl
}

func CheckLogin() {
	//TODO : remove viper dependency. Using `config` creates cyclic imports.
	email := config.GetEmail()
	token := config.GetAuthToken()

	if email == "" || token == "" {
		TerminateWithCustomError("Please log-in into your account first")
	}
}
