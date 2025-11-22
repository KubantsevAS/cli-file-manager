package command

import "cli/file-manager/internal/filesystem"

type RenameCommand struct {
	FS filesystem.FileSystem
}

func NewRenameCommand(fs filesystem.FileSystem) *RenameCommand {
	return &RenameCommand{FS: fs}
}

func (c *RenameCommand) Execute(src, newName string) error {
	return c.FS.Rename(src, newName)
}
