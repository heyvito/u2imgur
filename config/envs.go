package config

import (
	"fmt"
	"os"
)

type EnvData struct {
	ID     string
	Secret string
}

var EnvironData = EnvData{}

func init() {
	var (
		id     = os.Getenv("U2IMGUR_CLIENT_ID")
		secret = os.Getenv("U2IMGUR_CLIENT_SECRET")
	)

	if id == "" || secret == "" {
		fmt.Println("In order to use u2imgur, both U2IMGUR_CLIENT_ID and U2IMGUR_CLIENT_ID should be present in your environment variables.")
		os.Exit(1)
	}

	EnvironData.ID = id
	EnvironData.Secret = secret
}
