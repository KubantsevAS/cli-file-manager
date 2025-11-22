package cmd

import (
	"cli/file-manager/internal/command"
	"cli/file-manager/internal/filesystem"
)

func CopyCommand(src, dst string) error {
	fs := filesystem.NewLocalFS()
	copyCmd := command.NewCopyCommand(fs)

	return copyCmd.Execute(src, dst)
}
