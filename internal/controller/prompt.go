package controller

import (
	"github.com/AlecAivazis/survey/v2"

	"github.com/kintohub/kinto-cli-go/internal/utils"
)

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
	}

	return password
}

func TeleportPrompt(envName []string, envDetails map[string]string) string {
	selectedEnv := ""
	prompt := &survey.Select{
		Message: "Select environment:",
		Options: envName,
	}
	err := survey.AskOne(prompt, &selectedEnv)

	if err != nil {
		utils.TerminateWithCustomError("Aborted!")
	}

	return envDetails[selectedEnv]
}
