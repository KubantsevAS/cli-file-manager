package main

import (
	"bufio"
	"cli/file-manager/internal/color"
	"cli/file-manager/internal/router"
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {
	userName := flag.String("username", "Anonymous", "Weather format")
	flag.Parse()

	homeDir, err := os.UserHomeDir()
	if err != nil {
		warningMsg := fmt.Sprintf("Warning: failed to get home directory: %v\n", err)
		fmt.Printf("%s", color.Warning(warningMsg))
	}

	if err := os.Chdir(homeDir); err != nil {
		warningMsg := fmt.Sprintf("Warning: failed to change to home directory: %v\n", err)
		fmt.Printf("%s", color.Warning(warningMsg))
	}

	goodbyeMsg := fmt.Sprintf("\nThank you for using File Manager, %s, goodbye!", *userName)
	greetingMsg := fmt.Sprintf("Welcome to the File Manager, %s!", *userName)

	defer func() {
		fmt.Println(color.IntroOutro(goodbyeMsg))
	}()

	fmt.Println(color.IntroOutro(greetingMsg))
	fmt.Println(color.Info("Enter 'help' to get available commands info"))
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

		if command == ".exit" {
			break
		}

		executor := commandMap[command]
		if executor == nil {
			fmt.Println(color.Error("Invalid input"))
			continue
		}

		result, err := executor(args)

		if err != nil {
			fmt.Printf("%s\n", color.Error(fmt.Sprintf("Operation failed: %v", err)))
			continue
		}

		fmt.Println(color.Success(result))
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
