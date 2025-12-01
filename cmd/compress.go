package cmd

import (
	"cli/file-manager/internal/command"
	"cli/file-manager/internal/filesystem"
)

func CompressCommand(src, dst string) (string, error) {
	fs := filesystem.NewLocalFS()
	compressCmd := command.NewCompressCommand(fs)

	return compressCmd.Execute(src, dst)
}
