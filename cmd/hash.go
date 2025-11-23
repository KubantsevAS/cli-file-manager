package cmd

import (
	"fmt"

	"cli/file-manager/internal/color"
	"cli/file-manager/internal/command"
	"cli/file-manager/internal/filesystem"
)

func HashCommand(path string) error {
	fs := filesystem.NewLocalFS()
	hashCmd := command.NewHashCommand(fs)

	hash, err := hashCmd.Execute(path)
	if err != nil {
		return err
	}

	fmt.Println(color.Info(hash))
	return nil
}
