package cmd

import (
	"cli/file-manager/internal/command"
)

func HelpCommand() (string, error) {
	helpCmd := command.NewHelpCommand()
	return helpCmd.Execute()
}
