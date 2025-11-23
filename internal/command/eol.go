package command

import "cli/file-manager/internal/systeminfo"

type EOLCommand struct {
	SystemInfo systeminfo.SystemInfo
}

func NewEOLCommand(si systeminfo.SystemInfo) *EOLCommand {
	return &EOLCommand{SystemInfo: si}
}

func (c *EOLCommand) Execute() (string, error) {
	return c.SystemInfo.EOL()
}
