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
	userName := "Anonymous"

	goodbyeMsg := fmt.Sprintf("\nThank you for using File Manager, %s, goodbye!", userName)
	greetingMsg := fmt.Sprintf("Welcome to the File Manager, %s!", userName)

	defer func() {
		fmt.Println(color.Success(goodbyeMsg))
	}()

	fmt.Println(color.Success(greetingMsg))
	scanner := bufio.NewScanner(os.Stdin)
	commandMap := router.BuildCommandMap()

	for {
		currentDir, err := os.Getwd()
		if err != nil {
			panic(err.Error())
		}

		fmt.Printf("\n%s$ ", color.Path(currentDir))
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

		result, err := executor(args)

		if err != nil {
			fmt.Printf("%s\n", color.Error(fmt.Sprintf("Error: %v", err)))
			continue
		}

		fmt.Println(color.CommandExecuted(result))
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
