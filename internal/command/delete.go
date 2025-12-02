package command

import "cli/file-manager/internal/filesystem"

type DeleteCommand struct {
	FS filesystem.FileSystem
}

func NewDeleteCommand(fs filesystem.FileSystem) *DeleteCommand {
	return &DeleteCommand{FS: fs}
}

func (c *DeleteCommand) Execute(path string) (string, error) {
	return c.FS.Delete(path)
}
