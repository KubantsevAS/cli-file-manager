package command

import (
	"io"

	"cli/file-manager/internal/filesystem"
)

type ReadCommand struct {
	FS filesystem.FileSystem
}

func NewReadCommand(fs filesystem.FileSystem) *ReadCommand {
	return &ReadCommand{FS: fs}
}

func (c *ReadCommand) Execute(path string, w io.Writer) (string, error) {
	return c.FS.Read(path, w)
}
