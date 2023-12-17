package command

import (
	"math/rand"

	"fyne.io/fyne/v2"

	"pattern-command/ui"
)

type MoveCommand struct {
	baseCommand

	movement fyne.Vector2
}

func NewMoveCommand(ui *ui.UI, obj fyne.CanvasObject) Command {
	xMove := rand.Float32() * 10
	yMove := xMove

	command := MoveCommand{
		baseCommand: baseCommand{UI: ui},
		movement:    fyne.NewPos(xMove, yMove),
	}

	if obj != nil {
		command.target = obj
	} else {
		command.target = ui.GetRandomObject()
	}

	return &command
}

func (self *MoveCommand) Do() {
	self.target.Position().Add(self.movement)
}

func (self *MoveCommand) UnDo() {
	self.target.Position().Subtract(self.movement)
}

func (self *MoveCommand) ReDo() {
	self.Do()
}
