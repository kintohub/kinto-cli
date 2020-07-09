package controller

import (
	"bufio"
	"fmt"
	"github.com/kintohub/kinto-cli-go/internal/utils"
	"github.com/spf13/viper"
	"golang.org/x/crypto/ssh/terminal"
	"os"
	"strings"
)

func (c *Controller) Init() {
	email := viper.GetString("email")
	token := viper.GetString("authToken")

	if email == "" || token == "" {
		fmt.Printf("email plz?")
		reader := bufio.NewReader(os.Stdin)

		email, err := reader.ReadString('\n')

		if err != nil {
			utils.TerminateWithError(err)
		}

		fmt.Printf("Password plz?")
		passwordBytes, err := terminal.ReadPassword(0)

		if err != nil {
			utils.TerminateWithError(err)
		}
		fmt.Printf("example user and password %s:%s",
			strings.TrimSpace(email),
			strings.TrimSpace(string(passwordBytes)),
		)

		// TODO: Login
		//c.authClient.Login()
		//c.authClient.Register()

		// TODO: Save data
		viper.Set("email", email)
	} else {
		// TODO: Let the user know they are currently logged in
	}
}

func (c *Controller) Logout() {
	// TODO: Delete auth information in viper
	panic("implement me")
}
