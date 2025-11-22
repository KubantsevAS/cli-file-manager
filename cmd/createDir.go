package cmd

import (
	"fmt"

	"cli/file-manager/internal/command"
	"cli/file-manager/internal/filesystem"
)

func CreateDirCommand(name string) error {
	if name == "" {
		return fmt.Errorf("directory name is required")
	}

	fs := filesystem.NewLocalFS()
	createDirCmd := command.NewCreateDirCommand(fs)

	if err := createDirCmd.Execute(name); err != nil {
		return fmt.Errorf("failed to create directory: %w", err)
	}

	return nil
}
