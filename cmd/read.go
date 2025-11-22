package cmd

import (
	"fmt"
	"os"

	"cli/file-manager/internal/command"
	"cli/file-manager/internal/filesystem"
)

func ReadCommand(path string) error {
	if path == "" {
		return fmt.Errorf("file path is required")
	}

	fs := filesystem.NewLocalFS()
	readCmd := command.NewReadCommand(fs)

	if err := readCmd.Execute(path, os.Stdout); err != nil {
		return fmt.Errorf("failed to read file: %w", err)
	}

	return nil
}
