package command

import (
	"math/rand"

	"fyne.io/fyne/v2"

	"pattern-command/ui"
)

type ScaleCommand struct {
	baseCommand

	scale float32
}

func NewScaleCommand(ui *ui.UI, obj fyne.CanvasObject) Command {
	command := ScaleCommand{
		baseCommand: baseCommand{UI: ui},
		scale:       1.0 + rand.Float32(),
	}

	if obj != nil {
		command.target = obj
	} else {
		command.target = ui.GetRandomObject()
	}

	return &command
}

func (self *ScaleCommand) Do() {
	self.UI.ScaleObject(self.target, self.scale)
}

func (self *ScaleCommand) UnDo() {
	self.UI.ScaleObject(self.target, 1/self.scale)
}

func (self *ScaleCommand) ReDo() {
	self.Do()
}
