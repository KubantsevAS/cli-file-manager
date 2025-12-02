package command

import "cli/file-manager/internal/filesystem"

type CompressCommand struct {
	FS filesystem.FileSystem
}

func NewCompressCommand(fs filesystem.FileSystem) *CompressCommand {
	return &CompressCommand{FS: fs}
}

func (c *CompressCommand) Execute(src, dst string) (string, error) {
	return c.FS.Compress(src, dst)
}
