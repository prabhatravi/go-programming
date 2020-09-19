package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/mitchellh/go-homedir"
	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	info(app)
	flags(app)
	commands(app)

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

// add meta info to cli
func info(app *cli.App) {
	app.Name = "Daily Standup Helper CLI"
	app.Usage = "Reports git history"
	app.Author = "github.com/prabhatravi"
	app.Version = "1.0.0"
}

// add flags to cli
func flags(app *cli.App) {
	dir, _ := homedir.Dir()

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "user, u",
			Value: "",
			Usage: "git user name",
		},
		cli.StringFlag{
			Name:  "dir, d",
			Value: dir,
			Usage: "parent directory to start recursively searching for *.git files",
		},
		cli.StringFlag{
			Name:  "after, a",
			Value: time.Now().Add(-24 * time.Hour).Format("2006-01-02T15:04:05"),
			Usage: "when to start looking at commit history",
		},
	}
}

// add command to cli
func commands(app *cli.App) {
	app.Action = func(c *cli.Context) error {
		dir := c.String("dir")
		after := c.String("after")

		user := c.String("user")
		if len(user) == 0 {
			return fmt.Errorf("no 'user' flag value provided")
		}

		// 		err := runClient(dir, user, after)
		// 		if err != nil {
		// 			return err
		// 		}
		return nil
	}
}
