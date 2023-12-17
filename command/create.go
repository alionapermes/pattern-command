package command

import "pattern-command/ui"

type CreateCommand struct {
	baseCommand
}

func NewCreateCommand(ui *ui.UI) Command {
	return &CreateCommand{baseCommand{
		UI:     ui,
		target: nil,
	}}
}

func (self *CreateCommand) Do() {
	self.target = self.UI.CreateRandomObject()
}

func (self *CreateCommand) UnDo() {
	self.UI.DeleteObject(self.target)
}

func (self *CreateCommand) ReDo() {
	self.Do()
}
