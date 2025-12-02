package command

import (
	"fmt"
	"sort"
	"strings"
)

type HelpCommand struct{}

func NewHelpCommand() *HelpCommand {
	return &HelpCommand{}
}

type CommandInfo struct {
	Name        string
	Description string
	Usage       string
	Category    string
}

func (c *HelpCommand) Execute() (string, error) {
	commands := []CommandInfo{
		{"cd", "Change directory", "cd <path>", "Navigation"},
		{"up", "Go up one directory", "up", "Navigation"},
		{"ls", "List files and directories", "ls [directory]", "Navigation"},

		{"cat", "Read and display file content", "cat <file>", "File Operations"},
		{"add", "Create empty file", "add <filename>", "File Operations"},
		{"mkdir", "Create directory", "mkdir <dirname>", "File Operations"},
		{"cp", "Copy file", "cp <source> <destination>", "File Operations"},
		{"mv", "Move file", "mv <source> <destination>", "File Operations"},
		{"rn", "Rename file", "rn <file> <newname>", "File Operations"},
		{"rm", "Delete file", "rm <file>", "File Operations"},

		{"compress", "Compress file using gzip", "compress <file> <destination>", "Compression"},
		{"decompress", "Decompress gzip file", "decompress <file> <destination>", "Compression"},

		{"hash", "Calculate SHA256 hash of file", "hash <file>", "Utilities"},
		{".exit", "Exit the program", ".exit", "Utilities"},

		{"os", "Get system information", "os --<flag>", "System Info"},
	}

	categories := make(map[string][]CommandInfo)
	for _, cmd := range commands {
		categories[cmd.Category] = append(categories[cmd.Category], cmd)
	}

	categoryOrder := []string{"Navigation", "File Operations", "Compression", "Utilities", "System Info"}
	var result strings.Builder

	for _, category := range categoryOrder {
		if cmds, ok := categories[category]; ok {
			sort.Slice(cmds, func(i, j int) bool {
				return cmds[i].Name < cmds[j].Name
			})

			result.WriteString(fmt.Sprintf("\n%s:\n", category))
			for _, cmd := range cmds {
				result.WriteString(fmt.Sprintf("  %-12s %s\n", cmd.Name, cmd.Description))
				result.WriteString(fmt.Sprintf("    Usage: %s\n", cmd.Usage))
			}
		}
	}

	return strings.TrimSpace(result.String()), nil
}
