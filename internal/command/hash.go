package command

import (
	"cli/file-manager/internal/filesystem"
)

type HashCommand struct {
	FS filesystem.FileSystem
}

func NewHashCommand(fs filesystem.FileSystem) *HashCommand {
	return &HashCommand{FS: fs}
}

func (c *HashCommand) Execute(path string) (string, error) {
	return c.FS.Hash(path)
}
