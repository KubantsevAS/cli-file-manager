package cmd

import (
	"cli/file-manager/internal/command"
	"cli/file-manager/internal/filesystem"
	"fmt"
)

func DeleteCommand(path string) error {
	fs := filesystem.NewLocalFS()
	deleteCmd := command.NewDeleteCommand(fs)

	if err := deleteCmd.Execute(path); err != nil {
		return fmt.Errorf("failed to delete content: %w", err)
	}

	return nil
}
