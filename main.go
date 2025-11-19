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

		commandMap := map[string]func() error{
			"ls": func() error {
				err := cmd.ListCommand(commandSlice[1])
				return err
			},
			"cd": func() error {
				err := cmd.ChangeDirectoryCommand(commandSlice[1])
				return err
			},
			"up": func() error {
				err := cmd.UpCommand()
				return err
			},
		}

		executor := commandMap[commandSlice[0]]

		if executor == nil {
			fmt.Println("Invalid input")
			continue
		}

		if err := executor(); err != nil {
			fmt.Printf("Error: %v\n", err)
			continue
		}

		switch cmdLine {
		case "up":
			fmt.Println("Moved up one directory")
		case "cd":
			fmt.Printf("Changed directory to: %s\n", commandSlice[1])
		}
	}
}
