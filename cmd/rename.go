package cmd

import (
	"cli/file-manager/internal/command"
	"cli/file-manager/internal/filesystem"
)

func RenameCommand(src, newName string) (string, error) {
	fs := filesystem.NewLocalFS()
	renameCmd := command.NewRenameCommand(fs)

	return renameCmd.Execute(src, newName)
}
