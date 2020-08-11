package controller

import (
	"fmt"
	"github.com/kintohub/kinto-cli-go/internal/config"
	"github.com/kintohub/kinto-cli-go/internal/utils"
	"os"
	"strings"
)

func (c *Controller) Init() {

	email := config.GetString("emailKey")
	token := config.GetString("authTokenKey")

	if email == "" || token == "" {
		utils.WarningMessage("Please log-in into your account first")
		c.Login()

	} else {
		// TODO : Create a .kinto file and run init code
		_, err := os.Stat(".kinto")
		if os.IsNotExist(err) {
			_, err := os.Create(".kinto")
			if err != nil {
				utils.TerminateWithError(err)
			}
			utils.SuccessMessage("Repo initialized")
		} else {
			utils.WarningMessage("Repo is already initialized")
		}
	}
}

func (c *Controller) Login() {
	utils.StopSpinner()
	loginEmail := utils.EmailPrompt()
	passwordBytes := utils.PasswordPrompt()

	authToken, err := c.api.Login(
		strings.TrimSpace(loginEmail),
		strings.TrimSpace(passwordBytes),
	)

	if err != nil {
		utils.TerminateWithError(err)
	} else {
		email := config.GetString("emailKey")

		if email == loginEmail {
			utils.WarningMessage(fmt.Sprintf("Already logged in with %s",
				config.GetString("emailKey")))
		} else {
			config.SetAuthToken(authToken)
			config.SetEmail(loginEmail)
			config.Save()
			utils.SuccessMessage(fmt.Sprintf("Logged in successfully with %s",
				strings.TrimSpace(loginEmail)),
			)
		}

	}
}
