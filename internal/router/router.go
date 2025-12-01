package router

import (
	"cli/file-manager/cmd"
	"fmt"
)

type CommandExecutor func([]string) (string, error)

func BuildCommandMap() map[string]CommandExecutor {
	return map[string]CommandExecutor{
		"ls": func(args []string) (string, error) {
			dir := "."
			if len(args) > 0 {
				dir = args[0]
			}
			return cmd.ListCommand(dir)
		},
		"cd": func(args []string) (string, error) {
			if len(args) < 1 {
				return "", fmt.Errorf("cd: missing directory argument")
			}
			return cmd.ChangeDirCommand(args[0])
		},
		"up": func(args []string) (string, error) {
			return cmd.UpCommand()
		},
		"cat": func(args []string) (string, error) {
			if len(args) < 1 {
				return "", fmt.Errorf("cat: missing file argument")
			}
			return cmd.ReadCommand(args[0])
		},
		"mkdir": func(args []string) (string, error) {
			if len(args) < 1 {
				return "", fmt.Errorf("mkdir: missing directory name")
			}
			return cmd.CreateDirCommand(args[0])
		},
		"rm": func(args []string) (string, error) {
			if len(args) < 1 {
				return "", fmt.Errorf("rm: missing file name")
			}
			return cmd.DeleteCommand(args[0])
		},
		"add": func(args []string) (string, error) {
			if len(args) < 1 {
				return "", fmt.Errorf("add: missing file name")
			}
			return cmd.AddFileCommand(args[0])
		},
		"rn": func(args []string) (string, error) {
			if len(args) < 2 {
				return "", fmt.Errorf("rn: missing file path or new name")
			}
			return cmd.RenameCommand(args[0], args[1])
		},
		"cp": func(args []string) (string, error) {
			if len(args) < 2 {
				return "", fmt.Errorf("cp: missing file path or destination")
			}
			return cmd.CopyCommand(args[0], args[1])
		},
		"mv": func(args []string) (string, error) {
			if len(args) < 2 {
				return "", fmt.Errorf("mv: missing file path or destination")
			}
			return cmd.MoveCommand(args[0], args[1])
		},
		"os": func(args []string) (string, error) {
			return "", cmd.OSCommand(append([]string{"os"}, args...))
		},
		"hash": func(args []string) (string, error) {
			if len(args) < 1 {
				return "", fmt.Errorf("hash: missing file name")
			}
			return cmd.HashCommand(args[0])
		},
		"compress": func(args []string) (string, error) {
			if len(args) < 2 {
				return "", fmt.Errorf("cp: missing file path or destination")
			}
			return cmd.CompressCommand(args[0], args[1])
		},
	}
}
