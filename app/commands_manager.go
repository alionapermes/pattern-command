package app

import (
	"command-gui/command"
	"command-gui/object"

	"github.com/emirpasic/gods/stacks"
	lls "github.com/emirpasic/gods/stacks/linkedliststack"
)

const maxStackSize = 10

type commandManager struct {
	stack       stacks.Stack
	currenctObj object.Object
}

func newCommandManager() *commandManager {
	return &commandManager{
		stack: lls.New(),
	}
}

func (self *commandManager) do(cmd command.Command, obj object.Object) error {
	cmd.Execute(obj)

	self.currenctObj = obj
	self.stack.Push(cmd)

	return nil
}

func (self *commandManager) redo() error {
	if cmd, ok := self.stack.Peek(); ok {
		cmd.(command.Command).Execute(self.currenctObj)
	}

	return nil
}

func (self *commandManager) undo() error {
	self.stack.Pop()

	return nil
}

func (self *commandManager) recallStack() {
	for _, cmd := range self.stack.Values() {
		cmd.(command.Command).Execute(self.currenctObj)
	}
}
