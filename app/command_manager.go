package app

import (
	"fyne.io/fyne/v2"

	"pattern-command/command"
)

type cmdState int

const (
	cmdDone cmdState = iota
	cmdUndone
)

const maxStackSize = 3

type commandManager struct {
	stack        commandStack
	currenctObj  fyne.CanvasObject
	peekCmdState cmdState
}

func newCommandManager() *commandManager {
	return &commandManager{
		stack: *newCommandStack(),
	}
}

func (self *commandManager) do(cmd command.Command) {
	if self.stack.Size() == maxStackSize {
		self.stack.popFront()
	}

	self.stack.push(cmd)
	cmd.Do()
}

func (self *commandManager) undo() {
	if self.stack.Empty() {
		return
	}

	if self.peekCmdState == cmdUndone {
		self.stack.popFront()
		self.peekCmdState = cmdDone
	}

	if cmd := self.stack.peek(); cmd != nil {
		cmd.UnDo()
		self.peekCmdState = cmdUndone
	}
}

func (self *commandManager) redo() {
	if self.stack.Empty() {
		return
	}

	switch self.peekCmdState {
	case cmdUndone:
		self.stack.peek().ReDo()
	case cmdDone:
		cmd := self.stack.peek()
		cmd.SetTarget(self.currenctObj)
	}

	if self.stack.Size() > maxStackSize {
		self.stack.popFront()
	}
}
