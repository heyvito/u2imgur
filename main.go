package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/urfave/cli"

	"github.com/victorgama/u2imgur/commands"
	_ "github.com/victorgama/u2imgur/config"
)

func main() {
	app := cli.NewApp()
	app.Name = "u2imgur"
	app.Usage = "Uploads an image to imgur"
	app.Version = "0.1.0"
	app.Author = "Victor Gama <hey@vito.io>"
	app.Commands = []cli.Command{
		commands.ConfigCommand,
	}
	app.Action = func(c *cli.Context) error {
		var err error
		var link *string
		if c.NArg() == 0 {
			stat, _ := os.Stdin.Stat()
			if (stat.Mode()&os.ModeCharDevice) == 0 || (stat.Mode()&os.ModeNamedPipe) == 0 {
				link, err = commands.UploadImageFromStdin()
			} else {
				cli.ShowAppHelp(c)
				return nil
			}
		} else {
			path := strings.Join(c.Args(), " ")
			if strings.HasPrefix(path, "http") {
				link, err = commands.UploadImageFromUrl(path)
			} else {
				link, err = commands.UploadImageFromPath(path)
			}
		}
		if err != nil {
			return err
		}
		if err == nil && link == nil {
			panic("link and err equals nil.")
		}
		file := os.Stdout
		stat, err := file.Stat()
		if stat.Mode()&os.ModeNamedPipe == os.ModeNamedPipe {
			file.Write([]byte(fmt.Sprintf("%s\n", *link)))
		} else {
			file.Write([]byte(fmt.Sprintf("Upload complete. Access it through the following URL: %s\n", *link)))
		}
		return nil
	}
	app.Run(os.Args)
}
