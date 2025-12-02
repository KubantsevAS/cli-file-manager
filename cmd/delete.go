package cmd

import (
	"cli/file-manager/internal/command"
	"cli/file-manager/internal/filesystem"
)

func DeleteCommand(path string) (string, error) {
	fs := filesystem.NewLocalFS()
	deleteCmd := command.NewDeleteCommand(fs)

	return deleteCmd.Execute(path)
}
