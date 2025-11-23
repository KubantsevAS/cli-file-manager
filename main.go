package main

import (
	"bufio"
	"cli/file-manager/internal/color"
	"cli/file-manager/internal/router"
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
	commandMap := router.BuildCommandMap()

	for {
		currentDir, err := os.Getwd()
		if err != nil {
			panic(err.Error())
		}

		fmt.Printf("\n%s$ ", color.Info(currentDir))
		if !scanner.Scan() {
			break
		}

		command, args := parseCommand(strings.Fields(scanner.Text()))
		if command == "" {
			continue
		}

		executor := commandMap[command]
		if executor == nil {
			fmt.Println(color.Error("Invalid input"))
			continue
		}

		if err := executor(args); err != nil {
			fmt.Printf("%s\n", color.Error(fmt.Sprintf("Error: %v", err)))
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
