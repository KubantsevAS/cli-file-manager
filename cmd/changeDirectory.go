package cmd

import (
	"fmt"

	"cli/file-manager/internal/command"
	"cli/file-manager/internal/filesystem"
)

func ChangeDirectoryCommand(path string) error {
	fs := filesystem.NewLocalFS()
	changeDirectoryCmd := command.NewChangeDirectoryCommand(fs)

	if err := changeDirectoryCmd.Execute(path); err != nil {
		return fmt.Errorf("failed to change directory: %w", err)
	}

	return nil
}
