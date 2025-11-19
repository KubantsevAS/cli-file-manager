package cmd

import (
	"fmt"

	"cli/file-manager/filesystem"
	"cli/file-manager/internal/command"
)

func UpCommand() error {
	fs := filesystem.NewLocalFS()
	upCmd := command.NewUpCommand(fs)

	if err := upCmd.Execute(); err != nil {
		return fmt.Errorf("failed to go up directory: %w", err)
	}

	return nil
}
