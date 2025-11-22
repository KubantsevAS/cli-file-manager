package cmd

import (
	"cli/file-manager/internal/command"
	"cli/file-manager/internal/filesystem"
)

func MoveCommand(src, dst string) error {
	fs := filesystem.NewLocalFS()
	moveCmd := command.NewMoveCommand(fs)

	return moveCmd.Execute(src, dst)
}
