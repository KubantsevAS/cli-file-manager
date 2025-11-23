package main

import (
	"bufio"
	"cli/file-manager/cmd"
	"cli/file-manager/internal/color"
	"fmt"
	"os"
	"strings"
)

func main() {
	defer func() {
		fmt.Println(color.Success("\nThank you for using File Manager, Username, goodbye!"))
	}()

	fmt.Println(color.Success("Welcome to the File Manager, Username!"))
	scanner := bufio.NewScanner(os.Stdin)

	for {
		var cmdLine string
		currentDir, err := os.Getwd()

		if err != nil {
			panic(err.Error())
		}

		fmt.Printf("%s$ ", color.Info(currentDir))
		scanner.Scan()
		cmdLine = scanner.Text()

		commandSlice := strings.Fields(cmdLine)

		if len(commandSlice) == 0 {
			continue
		}

		commandMap := map[string]func() error{
			"ls": func() error {
				dir := ""
				if len(commandSlice) > 1 {
					dir = commandSlice[1]
				}
				return cmd.ListCommand(dir)
			},
			"cd": func() error {
				if len(commandSlice) < 2 {
					return fmt.Errorf("cd: missing directory argument")
				}
				return cmd.ChangeDirCommand(commandSlice[1])
			},
			"up": func() error {
				return cmd.UpCommand()
			},
			"cat": func() error {
				if len(commandSlice) < 2 {
					return fmt.Errorf("cat: missing file argument")
				}
				return cmd.ReadCommand(commandSlice[1])
			},
			"mkdir": func() error {
				if len(commandSlice) < 2 {
					return fmt.Errorf("mkdir: missing directory name")
				}
				return cmd.CreateDirCommand(commandSlice[1])
			},
			"rm": func() error {
				if len(commandSlice) < 2 {
					return fmt.Errorf("rm: missing file name")
				}
				return cmd.DeleteCommand(commandSlice[1])
			},
			"add": func() error {
				if len(commandSlice) < 2 {
					return fmt.Errorf("add: missing file name")
				}
				return cmd.AddFileCommand(commandSlice[1])
			},
			"rn": func() error {
				if len(commandSlice) < 3 {
					return fmt.Errorf("rn: missing file path or new name")
				}
				return cmd.RenameCommand(commandSlice[1], commandSlice[2])
			},
			"cp": func() error {
				if len(commandSlice) < 3 {
					return fmt.Errorf("cp: missing file path or destination")
				}
				return cmd.CopyCommand(commandSlice[1], commandSlice[2])
			},
			"mv": func() error {
				if len(commandSlice) < 3 {
					return fmt.Errorf("mv: missing file path or destination")
				}
				return cmd.MoveCommand(commandSlice[1], commandSlice[2])
			},
			"os": func() error {
				return cmd.OSCommand(commandSlice)
			},
		}

		executor := commandMap[commandSlice[0]]

		if executor == nil {
			fmt.Println(color.Error("Invalid input"))
			continue
		}

		if err := executor(); err != nil {
			fmt.Printf("%s\n", color.Error(fmt.Sprintf("Error: %v", err)))
			continue
		}

		switch commandSlice[0] {
		case "up":
			fmt.Println("Moved up one directory")
		case "cd":
			if len(commandSlice) > 1 {
				fmt.Printf("Changed directory to: %s\n", commandSlice[1])
			}
		}
	}
}
