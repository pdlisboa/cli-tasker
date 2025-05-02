package app

import (
	"github.com/urfave/cli"
)

func GetApp() *cli.App {
	app := cli.NewApp()
	app.Name = "Cli tasker"
	app.Usage = "Use this manage personal tasks"
	applyCommands(app)

	return app
}

func applyCommands(app *cli.App) {

	commands := []cli.Command{
		{
			Name:        "add",
			Description: "Add task",
			Flags: []cli.Flag{
				cli.StringFlag{Name: "name"},
			},
			Action: func() {},
		},
		{
			Name:        "list",
			Description: "List tasks",
			Action:      func() {},
		},
		{
			Name:        "done",
			Description: "Complete task",
			Flags: []cli.Flag{
				cli.IntFlag{Name: "id"},
			},
			Action: func() {},
		},
		{
			Name:        "cancel",
			Description: "Cancel  task",
			Flags: []cli.Flag{
				cli.IntFlag{Name: "id"},
			},
			Action: func() {},
		},
	}
	app.Commands = commands
}
