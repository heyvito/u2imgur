package commands

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/urfave/cli"

	"github.com/victorgama/u2imgur/config"
	"github.com/victorgama/u2imgur/remote"
)

var ConfigCommand = cli.Command{
	Name:  "config",
	Usage: "configures u2imgur to work with your imgur account",
	Action: func(c *cli.Context) error {
		currentConfig := config.GetSession()
		if currentConfig != nil {
			fmt.Println("Hey! Looks like you're already authenticated, but I'm not preventing you from reauthenticating.")
			fmt.Println("")
		}
		fmt.Println("Hey there! Let's go trough the authentication procedure, so you can start uploading things asap, shall we?")
		fmt.Println("")
		fmt.Println("Please visit the following URL: ")
		fmt.Printf("https://api.imgur.com/oauth2/authorize?client_id=%s&response_type=pin\n", config.EnvironData.ID)
		fmt.Println("")
		fmt.Println("Authorize the application, copy the token imgur will provide you and...")
		fmt.Print("...paste it here: ")

		voided := false
	readPin:
		if voided {
			fmt.Println("Looks like something went awry. Shall we try again?")
			fmt.Print("Please provide the token provided by the imgur server: ")
		}
		voided = true
		reader := bufio.NewReader(os.Stdin)
		pin, _ := reader.ReadString('\n')
		pin = strings.Trim(pin, "\n")
		fmt.Println("")
		fmt.Printf("Just a second... ")
		session, err := remote.GetTokenFromPin(pin)
		if err != nil {
			fmt.Printf("%s\n", err)
			goto readPin
		}
		config.SetSession(session)
		fmt.Println("OK.")
		fmt.Printf("We just stored your credentials to %s. You're now ready to upload gifs!", config.SessionFile)
		return nil
	},
}
