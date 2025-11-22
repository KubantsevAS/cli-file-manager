package cmd

import (
	"cli/file-manager/internal/command"
	"cli/file-manager/internal/filesystem"
)

func AddFileCommand(name string) error {
	fs := filesystem.NewLocalFS()
	addFileCmd := command.NewAddFileCommand(fs)

	return addFileCmd.Execute(name)
}
