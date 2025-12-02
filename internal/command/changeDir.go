package command

import "cli/file-manager/internal/filesystem"

type ChangeDirCommand struct {
	FS filesystem.FileSystem
}

func NewChangeDirCommand(fs filesystem.FileSystem) *ChangeDirCommand {
	return &ChangeDirCommand{FS: fs}
}

func (c *ChangeDirCommand) Execute(path string) (string, error) {
	return c.FS.ChangeDir(path)
}
