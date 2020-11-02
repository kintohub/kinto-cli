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
	"time"
)

//Gets the latest successful release from any service
func GetLatestSuccessfulRelease(releases map[string]*types.Release) *types.Release {
	if releases == nil || len(releases) == 0 {
		return nil
	}

	var latestRelease *types.Release
	for _, release := range releases {
		// filter release by only successfully deployed and with valid deployment type (exclude SUSPEND and UNDEPLOY)
		// NOT_SET is included as well for backward compatibility
		if release.Status.State == types.Status_SUCCESS &&
			(release.Type == types.Release_ROLLBACK ||
				release.Type == types.Release_DEPLOY ||
				release.Type == types.Release_NOT_SET) {
			if latestRelease == nil {
				latestRelease = release
				continue
			}

			latestCreatedAt, err := ptypes.Timestamp(latestRelease.CreatedAt)

			if err != nil {
				log.Error().Err(err).Msgf(
					"cannot parse timestamp %v to time for release %v", latestRelease.CreatedAt, latestRelease)
				continue
			}

			releaseCreateAt, err := ptypes.Timestamp(release.CreatedAt)

			if err != nil {
				log.Error().Err(err).Msgf(
					"cannot parse timestamp %v to time for release %v", release.CreatedAt, release)
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
func CheckIfPortAvailable(port int) int {
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

func CheckTeleportStatus(port int) bool {
	time.Sleep(1 * time.Second)
	address := fmt.Sprintf("0.0.0.0:%d", port)
	connection, err := net.Listen("tcp", address)
	if err != nil {
		return false
	} else {
		_ = connection.Close()
		return true
	}

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
func GetBlockPort(blockName string, release *types.Release) int {

	if strings.Contains(blockName, "redis") {
		return config.RedisPort
	} else if strings.Contains(blockName, "postgres") {
		return config.PostgresPort
	} else if strings.Contains(blockName, "mongodb") {
		return config.MongoPort
	} else if strings.Contains(blockName, "minio") {
		return config.MinioPort
	} else if strings.Contains(blockName, "mysql") {
		return config.MysqlPort
	} else {
		port, err := strconv.Atoi(release.RunConfig.Port)
		if err != nil {
			TerminateWithError(err)
		}
		return port
	}
}

func CanPortForwardToRelease(release *types.Release) bool {
	if release.RunConfig != nil &&
		(release.RunConfig.Type == types.Block_BACKEND_API ||
			release.RunConfig.Type == types.Block_WEB_APP ||
			release.RunConfig.Type == types.Block_CATALOG) {
		return true
	} else {
		return false
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
