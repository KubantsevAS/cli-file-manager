package cmd

import (
	"fmt"

	"cli/file-manager/filesystem"
	"cli/file-manager/internal/command"
)

func ListCommand(dir string) error {
	fs := filesystem.NewLocalFS()
	listCmd := command.NewListCommand(fs)

	files, err := listCmd.Execute(dir)

	if err != nil {
		return fmt.Errorf("failed to list directory: %w", err)
	}

	if len(files) == 0 {
		fmt.Println("Directory is empty")
		return nil
	}

	for _, file := range files {
		if file.IsDir() {
			fmt.Printf("./%s\n", file.Name())
		} else {
			fmt.Println(file.Name())
		}
	}

	return nil
}
