package main

import (
	"cli-tasker/app"
	"log"
	"os"
)

func main() {
	app := app.GetApp()

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
		panic(err)
	}
}
