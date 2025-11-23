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
		currentDir, err := os.Getwd()
		if err != nil {
			panic(err.Error())
		}

		fmt.Printf("%s$ ", color.Info(currentDir))
		if !scanner.Scan() {
			break
		}

		command, args := parseCommand(strings.Fields(scanner.Text()))
		if command == "" {
			continue
		}

		commandMap := map[string]func([]string) error{
			"ls": func(args []string) error {
				dir := "."
				if len(args) > 0 {
					dir = args[0]
				}
				return cmd.ListCommand(dir)
			},
			"cd": func(args []string) error {
				if len(args) < 1 {
					return fmt.Errorf("cd: missing directory argument")
				}
				return cmd.ChangeDirCommand(args[0])
			},
			"up": func(args []string) error {
				return cmd.UpCommand()
			},
			"cat": func(args []string) error {
				if len(args) < 1 {
					return fmt.Errorf("cat: missing file argument")
				}
				return cmd.ReadCommand(args[0])
			},
			"mkdir": func(args []string) error {
				if len(args) < 1 {
					return fmt.Errorf("mkdir: missing directory name")
				}
				return cmd.CreateDirCommand(args[0])
			},
			"rm": func(args []string) error {
				if len(args) < 1 {
					return fmt.Errorf("rm: missing file name")
				}
				return cmd.DeleteCommand(args[0])
			},
			"add": func(args []string) error {
				if len(args) < 1 {
					return fmt.Errorf("add: missing file name")
				}
				return cmd.AddFileCommand(args[0])
			},
			"rn": func(args []string) error {
				if len(args) < 2 {
					return fmt.Errorf("rn: missing file path or new name")
				}
				return cmd.RenameCommand(args[0], args[1])
			},
			"cp": func(args []string) error {
				if len(args) < 2 {
					return fmt.Errorf("cp: missing file path or destination")
				}
				return cmd.CopyCommand(args[0], args[1])
			},
			"mv": func(args []string) error {
				if len(args) < 2 {
					return fmt.Errorf("mv: missing file path or destination")
				}
				return cmd.MoveCommand(args[0], args[1])
			},
			"os": func(args []string) error {
				fullArgs := append([]string{"os"}, args...)
				return cmd.OSCommand(fullArgs)
			},
		}

		executor := commandMap[command]
		if executor == nil {
			fmt.Println(color.Error("Invalid input"))
			continue
		}

		if err := executor(args); err != nil {
			fmt.Printf("%s\n", color.Error(fmt.Sprintf("Error: %v", err)))
			continue
		}

		switch command {
		case "up":
			fmt.Println("Moved up one directory")
		case "cd":
			if len(args) > 0 {
				fmt.Printf("Changed directory to: %s\n", args[0])
			}
		}
	}
}

func parseCommand(cmdLine []string) (string, []string) {
	if len(cmdLine) == 0 {
		return "", []string{}
	}
	if len(cmdLine) == 1 {
		return cmdLine[0], []string{}
	}
	return cmdLine[0], cmdLine[1:]
}
