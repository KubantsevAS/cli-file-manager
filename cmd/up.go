package cmd

import (
	"fmt"

	"cli/file-manager/internal/command"
	"cli/file-manager/internal/filesystem"
)

func UpCommand() error {
	fs := filesystem.NewLocalFS()
	upCmd := command.NewUpCommand(fs)

	if err := upCmd.Execute(); err != nil {
		return fmt.Errorf("failed to go up directory: %w", err)
	}

	return nil
}
