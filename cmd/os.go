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

	flagHandlers := map[string]func() error{
		"homedir": func() error {
			sysInfo := systeminfo.NewLocalSystem()
			homeDirCmd := command.NewHomeDirCommand(sysInfo)
			homeDir, err := homeDirCmd.Execute()
			if err != nil {
				return fmt.Errorf("failed to get home directory: %w", err)
			}
			fmt.Println(color.Info(homeDir))
			return nil
		},
		"EOL": func() error {
			// TODO: Implement EOL command
			return fmt.Errorf("os: --EOL not yet implemented")
		},
		"cpus": func() error {
			// TODO: Implement CPUs command
			return fmt.Errorf("os: --cpus not yet implemented")
		},
		"username": func() error {
			// TODO: Implement username command
			return fmt.Errorf("os: --username not yet implemented")
		},
		"architecture": func() error {
			// TODO: Implement architecture command
			return fmt.Errorf("os: --architecture not yet implemented")
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
