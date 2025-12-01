package cmd

import (
	"flag"
	"fmt"

	"cli/file-manager/internal/color"
	"cli/file-manager/internal/command"
	"cli/file-manager/internal/systeminfo"
)

func OSCommand(args []string) error {
	fs := flag.NewFlagSet("os", flag.ContinueOnError)

	fs.Bool("homedir", false, "Get home directory")
	fs.Bool("EOL", false, "Get EOL (default system End-Of-Line)")
	fs.Bool("cpus", false, "Get host machine CPUs info")
	fs.Bool("username", false, "Get current system user name")
	fs.Bool("architecture", false, "Get CPU architecture")

	err := fs.Parse(args[1:])
	if err != nil {
		return fmt.Errorf("os: %w", err)
	}

	sysInfo := systeminfo.NewLocalSystem()

	flagHandlers := map[string]func() error{
		"homedir": func() error {
			return executeCommand(command.NewHomeDirCommand(sysInfo))
		},
		"EOL": func() error {
			return executeCommand(command.NewEOLCommand(sysInfo))
		},
		"cpus": func() error {
			return executeCommand(command.NewCPUsCommand(sysInfo))
		},
		"username": func() error {
			return executeCommand(command.NewUsernameCommand(sysInfo))
		},
		"architecture": func() error {
			return executeCommand(command.NewArchitectureCommand(sysInfo))
		},
	}

	var handler func() error
	fs.Visit(func(f *flag.Flag) {
		if h, ok := flagHandlers[f.Name]; ok {
			handler = h
		}
	})

	if handler != nil {
		return handler()
	}

	return fmt.Errorf("os: no flag specified")
}

type ExecutableCmd interface {
	Execute() (string, error)
}

func executeCommand(cmd ExecutableCmd) error {
	info, err := cmd.Execute()
	if err != nil {
		return err
	}
	fmt.Print(color.Info(info))
	return nil
}
