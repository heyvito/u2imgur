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
		var result string
		if stat.Mode()&os.ModeNamedPipe == os.ModeNamedPipe {
			result = fmt.Sprintf("%s\n", *link)
		} else {
			result = fmt.Sprintf("Upload complete. The image is available at: %s\n", *link)
		}
		file.Write([]byte(result))
		return nil
	}
	app.Run(os.Args)
}
