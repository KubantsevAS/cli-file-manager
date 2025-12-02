package cmd

import (
	"os"

	"cli/file-manager/internal/command"
	"cli/file-manager/internal/filesystem"
)

func ReadCommand(path string) (string, error) {
	fs := filesystem.NewLocalFS()
	readCmd := command.NewReadCommand(fs)

	return readCmd.Execute(path, os.Stdout)
}
