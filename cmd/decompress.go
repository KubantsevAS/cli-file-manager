package cmd

import (
	"cli/file-manager/internal/command"
	"cli/file-manager/internal/filesystem"
)

func DecompressCommand(src, dst string) (string, error) {
	fs := filesystem.NewLocalFS()
	decompressCmd := command.NewDecompressCommand(fs)

	return decompressCmd.Execute(src, dst)
}
