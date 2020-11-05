package controller

import (
	"fmt"
	"github.com/kintohub/kinto-cli/internal/config"
	"github.com/kintohub/kinto-cli/internal/utils"
	"strings"
)

//Set MasterHost for CLI or reset it to default production host.
func (c *Controller) Init(masterHost string) {

	if masterHost == "Default" || masterHost == "default" {
		config.SetMasterHost(config.DefaultMasterHost)
		config.Save()
	} else {
		config.SetMasterHost(masterHost)
		config.Save()
	}

	//	TODO: To be handled in the future once .kinto files are ready
	/*	utils.CheckLogin()
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
	*/
}

func (c *Controller) Login() {
	utils.StopSpinner()
	loginEmail := EmailPrompt()
	passwordBytes := PasswordPrompt()

	authToken, err := c.api.Login(
		strings.TrimSpace(loginEmail),
		strings.TrimSpace(passwordBytes),
	)

	if err != nil {
		utils.TerminateWithError(err)
		return
	}
	email := config.GetEmail()

	if email == loginEmail {
		utils.WarningMessage(fmt.Sprintf("Already logged in with %s",
			config.GetEmail()))
	} else {
		config.SetAuthToken(authToken)
		config.SetEmail(loginEmail)
		config.Save()
		utils.SuccessMessage(fmt.Sprintf("Logged in successfully with %s",
			strings.TrimSpace(loginEmail)),
		)
	}

}
