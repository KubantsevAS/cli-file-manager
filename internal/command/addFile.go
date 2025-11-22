package command

import "cli/file-manager/internal/filesystem"

type addFileCommand struct {
	FS filesystem.FileSystem
}

func NewAddFileCommand(fs filesystem.FileSystem) *addFileCommand {
	return &addFileCommand{FS: fs}
}

func (c *addFileCommand) Execute(name string) error {
	return c.FS.AddFile(name)
}
