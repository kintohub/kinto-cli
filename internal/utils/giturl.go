package utils

import (
	"github.com/Terry-Mao/goconf"
	_ "github.com/Terry-Mao/goconf"
	"strings"
	_ "strings"
)

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
