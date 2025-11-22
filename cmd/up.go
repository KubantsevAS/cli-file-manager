package cmd

import (
	"cli/file-manager/internal/command"
	"cli/file-manager/internal/filesystem"
)

func UpCommand() error {
	fs := filesystem.NewLocalFS()
	upCmd := command.NewUpCommand(fs)

	return upCmd.Execute()
}
