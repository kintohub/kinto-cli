package controller

import (
	"bufio"
	"github.com/gookit/color"
	"github.com/kintohub/kinto-cli-go/internal/config"
	"github.com/kintohub/kinto-cli-go/internal/utils"
	"golang.org/x/crypto/ssh/terminal"
	"os"
	"strings"
)

func (c *Controller) Init() {

	email := config.GetString("email")
	token := config.GetString("authToken")

	if email == "" || token == "" {
		color.Red.Printf("\nPlease log-in into your account first\n")
		c.Login()

	} else {
		// TODO : Create a .kinto file and run init code
		_, err := os.Stat(".kinto")
		if os.IsNotExist(err) {
			_, err := os.Create(".kinto")
			if err != nil {
				color.Red.Println(err)
			}
			color.Green.Printf("\nRepo initialized\n")
		} else {
			color.Red.Printf("\nRepo is already initialized\n")
		}
	}
}

func (c *Controller) Register() {

	color.Gray.Printf("\nEmail address: ")
	reader := bufio.NewReader(os.Stdin)
	email, err := reader.ReadString('\n')

	if err != nil {
		utils.TerminateWithError(err)
	}

	color.Gray.Printf("Password: ")
	passwordBytes, err := terminal.ReadPassword(int(os.Stdin.Fd()))

	if err != nil {
		utils.TerminateWithError(err)
	}

	authToken, err := c.api.Register(
		strings.TrimSpace(email),
		strings.TrimSpace(string(passwordBytes)),
	)

	if err != nil {
		utils.TerminateWithError(err)
	} else {
		config.SetAuthToken(authToken)
		config.SetEmail(email)
		config.WriteConfig()
		color.Green.Printf("\nRegistered successfully with %s\n",
			strings.TrimSpace(email),
		)
	}
}

func (c *Controller) Login() {
	color.Gray.Printf("\nEmail address: ")
	reader := bufio.NewReader(os.Stdin)

	loginEmail, err := reader.ReadString('\n')

	if err != nil {
		utils.TerminateWithError(err)
	}

	color.Gray.Printf("Password: ")
	passwordBytes, err := terminal.ReadPassword(int(os.Stdin.Fd()))

	if err != nil {
		utils.TerminateWithError(err)
	}

	authToken, err := c.api.Login(
		strings.TrimSpace(loginEmail),
		strings.TrimSpace(string(passwordBytes)),
	)

	if err != nil {
		utils.TerminateWithError(err)
	} else {
		email := config.GetString("email")

		if email == loginEmail {
			color.Red.Printf("\nAlready logged in with %s\n", config.GetString("email"))
		} else {
			config.SetAuthToken(authToken)
			config.SetEmail(loginEmail)
			config.WriteConfig()
			color.Green.Printf("\nLogged in successfully with %s\n",
				strings.TrimSpace(loginEmail),
			)
		}

	}
}
