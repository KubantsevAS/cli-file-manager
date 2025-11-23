package command

import "cli/file-manager/internal/systeminfo"

type HomeDirCommand struct {
	SystemInfo systeminfo.SystemInfo
}

func NewHomeDirCommand(si systeminfo.SystemInfo) *HomeDirCommand {
	return &HomeDirCommand{SystemInfo: si}
}

func (c *HomeDirCommand) Execute() (string, error) {
	return c.SystemInfo.HomeDir()
}
