package command

import (
	"fyne.io/fyne/v2"

	"pattern-command/ui"
)

type Command interface {
	Do()
	UnDo()
	ReDo()

	SetTarget(target fyne.CanvasObject)
}

type baseCommand struct {
	*ui.UI

	target fyne.CanvasObject
}

func (self *baseCommand) SetTarget(target fyne.CanvasObject) {
	self.target = target
}
