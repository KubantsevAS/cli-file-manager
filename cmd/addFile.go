package cmd

import (
	"cli/file-manager/internal/command"
	"cli/file-manager/internal/filesystem"
	"fmt"
)

func AddFileCommand(name string) error {
	if name == "" {
		return fmt.Errorf("file name is required")
	}

	fs := filesystem.NewLocalFS()
	addFileCmd := command.NewAddFileCommand(fs)

	return addFileCmd.Execute(name)
}
