package cmd

import (
	"fmt"
	"strings"

	"cli/file-manager/internal/color"
	"cli/file-manager/internal/command"
	"cli/file-manager/internal/filesystem"
)

func ListCommand(dir string) (string, error) {
	if dir == "" {
		dir = "."
	}

	fs := filesystem.NewLocalFS()
	listCmd := command.NewListCommand(fs)

	files, err := listCmd.Execute(dir)

	if err != nil {
		return "", err
	}

	if len(files) == 0 {
		return color.Info("Directory is empty"), nil
	}

	dirContent := make([]string, len(files))

	for idx, file := range files {
		name := file.Name()

		if file.IsDir() {
			dirContent[idx] = fmt.Sprintf("📁 %s\n", color.Folder(name))
		} else {
			dirContent[idx] = color.File(name) + "\n"
		}
	}

	return strings.TrimRight(strings.Join(dirContent, ""), "\n"), nil
}
