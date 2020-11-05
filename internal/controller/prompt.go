package controller

import (
	"github.com/AlecAivazis/survey/v2"
	"github.com/kintohub/kinto-cli/internal/api"

	"github.com/kintohub/kinto-cli/internal/utils"
)

// Contains different types of prompts for the UX.

func EmailPrompt() string {
	email := ""
	prompt := &survey.Input{
		Message: "Email address:",
	}
	err := survey.AskOne(prompt, &email, survey.WithValidator(survey.Required),
		survey.WithIcons(func(icons *survey.IconSet) {
			icons.Question.Text = "?"
			icons.Question.Format = "green"
		}))

	if err != nil {
		utils.TerminateWithCustomError("Aborted!")
		return ""
	}

	return email
}

func PasswordPrompt() string {
	password := ""
	prompt := &survey.Password{
		Message: "Password:",
	}
	err := survey.AskOne(prompt, &password, survey.WithValidator(survey.Required),
		survey.WithIcons(func(icons *survey.IconSet) {
			icons.Question.Text = "?"
			icons.Question.Format = "green"
		}))

	if err != nil {
		utils.TerminateWithCustomError("Aborted!")
		return ""
	}

	return password
}

// Selection prompt, to be used in screens requiring selection of single entry from multiple options.
func SelectionPrompt(envDetails []api.EnvDetails) (string, string) {
	var envNames []string
	var selectedEnv int

	for _, i := range envDetails {
		envNames = append(envNames, i.EnvName)
	}

	prompt := &survey.Select{
		Message: "Select environment:",
		Options: envNames,
	}
	err := survey.AskOne(prompt, &selectedEnv)

	if err != nil {
		utils.TerminateWithCustomError("Aborted!")
		return "", ""
	}

	return envDetails[selectedEnv].EnvId, envDetails[selectedEnv].ClusterId
}
