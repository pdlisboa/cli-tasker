package app

import (
	"cli-tasker/task"

	"github.com/urfave/cli"
)

func GetApp() *cli.App {
	app := cli.NewApp()
	app.Name = "Cli tasker"
	app.Usage = "Manage personal tasks"
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
			Action: func(ctx *cli.Context) {
				task.AddTask(ctx.String("name"))

			},
		},
		{
			Name:        "list",
			Description: "List tasks",
			Action:      func(_ *cli.Context) { task.ListTasks() },
		},
		{
			Name:        "done",
			Description: "Complete task",
			Flags: []cli.Flag{
				cli.UintFlag{Name: "id"},
			},
			Action: func(ctx *cli.Context) { task.CompleteTask(ctx.Uint("id")) },
		},
		{
			Name:        "cancel",
			Description: "Cancel  task",
			Flags: []cli.Flag{
				cli.UintFlag{Name: "id"},
			},
			Action: func(ctx *cli.Context) { task.CancelTask(ctx.Uint("id")) },
		},
	}
	app.Commands = commands
}
