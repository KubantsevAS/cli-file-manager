package command

import "cli/file-manager/internal/systeminfo"

type CPUsCommand struct {
	SystemInfo systeminfo.SystemInfo
}

func NewCPUsCommand(si systeminfo.SystemInfo) *CPUsCommand {
	return &CPUsCommand{SystemInfo: si}
}

func (c *CPUsCommand) Execute() (string, error) {
	return c.SystemInfo.CPUs()
}
