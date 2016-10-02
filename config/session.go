package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"os/user"
	"path/filepath"
	"time"
)

type Session struct {
	AccessToken  string    `json:"access_token"`
	RefreshToken string    `json:"refresh_token"`
	ExpiresAt    time.Time `json:"expires_at"`
}

func (s *Session) StillValid() bool {
	return time.Now().Before(s.ExpiresAt)
}

var SessionFile string

func init() {
	usr, _ := user.Current()
	SessionFile = filepath.Join(usr.HomeDir, ".u2imgur")
}

func handleError(err error) {
	if err != nil {
		fmt.Printf("There was a problem reading your authentication data: %s\n", err)
		fmt.Printf("Deleting %s will force us to authenticate you and maybe fix the issue.\n", SessionFile)
		os.Exit(1)
	}
}

func GetSession() *Session {
	var err error
	var data Session

	if _, err := os.Stat(SessionFile); os.IsNotExist(err) {
		return nil
	}
	handleError(err)

	file, err := os.Open(SessionFile)
	handleError(err)
	defer file.Close()

	err = json.NewDecoder(file).Decode(&data)
	handleError(err)
	return &data
}

func SetSession(session *Session) {
	bytes, _ := json.Marshal(session)
	ioutil.WriteFile(SessionFile, bytes, 0755)
}
