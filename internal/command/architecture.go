package command

import "cli/file-manager/internal/systeminfo"

type ArchitectureCommand struct {
	SystemInfo systeminfo.SystemInfo
}

func NewArchitectureCommand(si systeminfo.SystemInfo) *ArchitectureCommand {
	return &ArchitectureCommand{SystemInfo: si}
}

func (c *ArchitectureCommand) Execute() (string, error) {
	return c.SystemInfo.Architecture()
}
