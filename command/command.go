package command

import "command-gui/object"

type Command interface {
	Execute(object.Object)
}
