package main

import (
	"log"

	"pattern-command/app"
)

const mainWindowTitle = "commands-gui-app"

func main() {
	commandsApp := app.New(mainWindowTitle)

	if err := commandsApp.Start(); err != nil {
		log.Fatalf("app failed with error: %s", err.Error())
	}
}
