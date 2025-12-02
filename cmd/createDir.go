package cmd

import (
	"cli/file-manager/internal/command"
	"cli/file-manager/internal/filesystem"
)

func CreateDirCommand(name string) (string, error) {
	fs := filesystem.NewLocalFS()
	createDirCmd := command.NewCreateDirCommand(fs)

	return createDirCmd.Execute(name)
}
