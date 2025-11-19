package command

import "cli/file-manager/internal/filesystem"

type ChangeDirectoryCommand struct {
	FS filesystem.FileSystem
}

func NewChangeDirectoryCommand(fs filesystem.FileSystem) *ChangeDirectoryCommand {
	return &ChangeDirectoryCommand{FS: fs}
}

func (c *ChangeDirectoryCommand) Execute(path string) error {
	return c.FS.ChangeDirectory(path)
}
