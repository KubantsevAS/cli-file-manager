package command

import "cli/file-manager/internal/filesystem"

type UpCommand struct {
	FS filesystem.FileSystem
}

func NewUpCommand(fs filesystem.FileSystem) *UpCommand {
	return &UpCommand{FS: fs}
}

func (c *UpCommand) Execute() (string, error) {
	return c.FS.ChangeDir("..")
}
