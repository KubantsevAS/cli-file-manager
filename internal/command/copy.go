package command

import "cli/file-manager/internal/filesystem"

type CopyCommand struct {
	FS filesystem.FileSystem
}

func NewCopyCommand(fs filesystem.FileSystem) *CopyCommand {
	return &CopyCommand{FS: fs}
}

func (c *CopyCommand) Execute(src, dst string) error {
	return c.FS.Copy(src, dst)
}
