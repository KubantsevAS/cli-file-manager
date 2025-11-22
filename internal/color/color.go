package color

const (
	reset  = "\033[0m"
	red    = "\033[31m"
	green  = "\033[32m"
	yellow = "\033[33m"
	blue   = "\033[34m"
	purple = "\033[35m"
	cyan   = "\033[36m"
	white  = "\033[37m"
	gray   = "\033[90m"

	bold       = "\033[1m"
	boldRed    = "\033[1;31m"
	boldGreen  = "\033[1;32m"
	boldYellow = "\033[1;33m"
	boldBlue   = "\033[1;34m"
	boldPurple = "\033[1;35m"
	boldCyan   = "\033[1;36m"
	boldWhite  = "\033[1;37m"
)

func Folder(s string) string {
	return yellow + s + reset
}

func Error(s string) string {
	return boldRed + s + reset
}

func Success(s string) string {
	return boldGreen + s + reset
}

func Warning(s string) string {
	return boldYellow + s + reset
}

func Info(s string) string {
	return boldCyan + s + reset
}
