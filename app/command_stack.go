package app

import (
	"pattern-command/command"

	dll "github.com/emirpasic/gods/lists/doublylinkedlist"
)

type commandStack struct {
	*dll.List
}

func newCommandStack() *commandStack {
	return &commandStack{
		List: dll.New(),
	}
}

func (self *commandStack) peek() command.Command {
	idx := self.Size() - 1

	if cmd, ok := self.Get(idx); ok {
		return cmd.(command.Command)
	}

	return nil
}

func (self *commandStack) popFront() command.Command {
	if cmd, ok := self.Get(0); ok {
		self.Remove(0)
		return cmd.(command.Command)
	}

	return nil
}

func (self *commandStack) popBack() command.Command {
	idx := self.Size() - 1

	if cmd, ok := self.Get(idx); ok {
		self.Remove(idx)
		return cmd.(command.Command)
	}

	return nil
}

func (self *commandStack) push(cmd command.Command) {
	self.Add(cmd)
}
