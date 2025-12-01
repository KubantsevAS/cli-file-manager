package command

import "cli/file-manager/internal/filesystem"

type DecompressCommand struct {
	FS filesystem.FileSystem
}

func NewDecompressCommand(fs filesystem.FileSystem) *DecompressCommand {
	return &DecompressCommand{FS: fs}
}

func (c *DecompressCommand) Execute(src, dst string) (string, error) {
	return c.FS.Decompress(src, dst)
}
