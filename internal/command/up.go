package command

import "cli/file-manager/filesystem"

type UpCommand struct {
	FS filesystem.FileSystem
}

func NewUpCommand(fs filesystem.FileSystem) *UpCommand {
	return &UpCommand{FS: fs}
}

func (c *UpCommand) Execute() error {
	return c.FS.ChangeDirectory("..")
}
