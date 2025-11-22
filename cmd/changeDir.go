package cmd

import (
	"cli/file-manager/internal/command"
	"cli/file-manager/internal/filesystem"
)

func ChangeDirCommand(path string) error {
	fs := filesystem.NewLocalFS()
	changeDirectoryCmd := command.NewChangeDirCommand(fs)

	return changeDirectoryCmd.Execute(path)
}
