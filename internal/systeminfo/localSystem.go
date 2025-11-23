package systeminfo

import (
	"fmt"
	"os"
	"os/user"
	"runtime"
	"strings"

	"golang.org/x/sys/cpu"
)

type LocalSystem struct{}

func NewLocalSystem() *LocalSystem {
	return &LocalSystem{}
}

func (s *LocalSystem) HomeDir() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("failed to get home directory: %w", err)
	}
	return homeDir, nil
}

func (s *LocalSystem) EOL() (string, error) {
	if runtime.GOOS == "windows" {
		return "\\r\\n", nil
	}
	return "\\n", nil
}

func (s *LocalSystem) CPUs() (string, error) {
	var cpuInfo strings.Builder
	cpuInfo.WriteString(fmt.Sprintf("Overall amount of CPUs: %d\n", runtime.NumCPU()))
	cpuInfo.WriteString(fmt.Sprintf("Operating System: %s\n", runtime.GOOS))
	cpuInfo.WriteString(fmt.Sprintf("Supports AVX: %t\n", cpu.X86.HasAVX))
	cpuInfo.WriteString(fmt.Sprintf("Supports SSE4.2: %t", cpu.X86.HasSSE42))

	return cpuInfo.String(), nil
}

func (s *LocalSystem) Username() (string, error) {
	currentUser, err := user.Current()
	if err != nil {
		return "", fmt.Errorf("failed to get current user: %w", err)
	}
	return currentUser.Username, nil
}

func (s *LocalSystem) Architecture() (string, error) {
	return fmt.Sprintf("Architecture: %s", runtime.GOARCH), nil
}
