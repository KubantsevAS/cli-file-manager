package command

import (
	"cli/file-manager/internal/filesystem"
)

type CreateDirCommand struct {
	FS filesystem.FileSystem
}

func NewCreateDirCommand(fs filesystem.FileSystem) *CreateDirCommand {
	return &CreateDirCommand{FS: fs}
}

func (c *CreateDirCommand) Execute(name string) (string, error) {
	return c.FS.CreateDir(name)
}
