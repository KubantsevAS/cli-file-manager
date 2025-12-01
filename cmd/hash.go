package cmd

import (
	"fmt"

	"cli/file-manager/internal/color"
	"cli/file-manager/internal/command"
	"cli/file-manager/internal/filesystem"
)

func HashCommand(path string) (string, error) {
	fs := filesystem.NewLocalFS()
	hashCmd := command.NewHashCommand(fs)

	hash, err := hashCmd.Execute(path)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("File '%s' hash info:\n%s", path, color.Info(hash)), nil
}
