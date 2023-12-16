package app

import (
	"fyne.io/fyne/v2"
	fyneApp "fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
)

type App struct {
	mainWindow     fyne.Window
	commandManager *commandManager
}

func New(title string) *App {
	app := &App{
		mainWindow:     fyneApp.New().NewWindow(title),
		commandManager: newCommandManager(),
	}

	return app.setup()
}

func (self *App) Start() error {
	self.mainWindow.SetFullScreen(true)

	self.mainWindow.ShowAndRun()
	return nil
}

func (self *App) setup() *App {
	self.mainWindow.SetContent(widget.NewLabel("Hello World!"))

	return self
}
