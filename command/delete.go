package command

import (
	"fyne.io/fyne/v2"

	"pattern-command/ui"
)

type DeleteCommand struct {
	baseCommand
}

func NewDeleteCommand(ui *ui.UI, obj fyne.CanvasObject) Command {
	command := DeleteCommand{baseCommand{UI: ui}}

	if obj != nil {
		command.target = obj
	} else {
		command.target = ui.GetRandomObject()
	}

	return &command
}

func (self *DeleteCommand) Do() {
	self.UI.DeleteObject(self.target)
}

func (self *DeleteCommand) UnDo() {
	self.UI.CreateObject(self.target)
}

func (self *DeleteCommand) ReDo() {
	self.Do()
}
