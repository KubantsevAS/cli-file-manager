package command

import "cli/file-manager/internal/filesystem"

type MoveCommand struct {
	FS filesystem.FileSystem
}

func NewMoveCommand(fs filesystem.FileSystem) *MoveCommand {
	return &MoveCommand{FS: fs}
}

func (c *MoveCommand) Execute(src, dst string) (string, error) {
	return c.FS.Move(src, dst)
}
