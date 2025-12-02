package color

const (
	reset      = "\033[0m"
	green      = "\033[32m"
	yellow     = "\033[33m"
	white      = "\033[37m"
	boldRed    = "\033[1;31m"
	boldGreen  = "\033[1;32m"
	boldYellow = "\033[1;33m"
	boldBlue   = "\033[1;34m"
	boldCyan   = "\033[1;36m"
)

func Folder(s string) string {
	return yellow + s + reset
}

func File(s string) string {
	return white + s + reset
}

func Success(s string) string {
	return green + s + reset
}

func Error(s string) string {
	return boldRed + s + reset
}

func IntroOutro(s string) string {
	return boldGreen + s + reset
}

func Warning(s string) string {
	return boldYellow + s + reset
}

func Path(s string) string {
	return boldBlue + s + reset
}

func Info(s string) string {
	return boldCyan + s + reset
}
