package app

import (
	fyneApp "fyne.io/fyne/v2/app"

	"pattern-command/command"
	"pattern-command/ui"
)

type App struct {
	commandManager *commandManager
	ui             *ui.UI
}

func New(title string) *App {
	app := &App{
		ui:             ui.NewUI(fyneApp.New(), title),
		commandManager: newCommandManager(),
	}

	return app.setup()
}

func (self *App) Start() error {
	self.ui.Run()
	return nil
}

func (self *App) setup() *App {
	self.ui.RegisterBtn("Create", func() {
		cmd := command.NewCreateCommand(self.ui)
		self.commandManager.do(cmd)
	})

	self.ui.RegisterBtn("Delete", func() {
		cmd := command.NewDeleteCommand(self.ui, nil)
		self.commandManager.do(cmd)
	})

	self.ui.RegisterBtn("Move", func() {
		cmd := command.NewMoveCommand(self.ui, nil)
		self.commandManager.do(cmd)
	})

	self.ui.RegisterBtn("Scale", func() {
		cmd := command.NewScaleCommand(self.ui, nil)
		self.commandManager.do(cmd)
	})

	self.ui.RegisterBtn("ReDo", func() {
		self.commandManager.redo()
	})

	self.ui.RegisterBtn("UnDo", func() {
		self.commandManager.undo()
	})

	self.ui.Setup()
	return self
}
