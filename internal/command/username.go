package command

import "cli/file-manager/internal/systeminfo"

type UsernameCommand struct {
	SystemInfo systeminfo.SystemInfo
}

func NewUsernameCommand(si systeminfo.SystemInfo) *UsernameCommand {
	return &UsernameCommand{SystemInfo: si}
}

func (c *UsernameCommand) Execute() (string, error) {
	return c.SystemInfo.Username()
}
