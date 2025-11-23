package systeminfo

import (
	"fmt"
	"os"
	"os/user"
	"runtime"
	"strings"

	"github.com/shirou/gopsutil/v3/cpu"
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
	var cpuData strings.Builder
	cpuInfo, err := cpu.Info()
	if err != nil {
		return "", fmt.Errorf("failed to get current user: %w", err)
	}

	numCPU := runtime.NumCPU()
	cpuData.WriteString(fmt.Sprintf("Overall amount of CPUs: %d\n", numCPU))

	switch {
	case len(cpuInfo) == 1:
		for i := range numCPU {
			info := cpuInfo[0]
			cpuData.WriteString(getCPUInfo(info, i))
		}
	case len(cpuInfo) > 1:
		for i, info := range cpuInfo {
			cpuData.WriteString(getCPUInfo(info, i))
		}
	default:
		for i := 0; i < numCPU; i++ {
			cpuData.WriteString(fmt.Sprintf("CPU %d: (model information not available)\n", i+1))
		}
	}

	return cpuData.String(), nil
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

func getCPUInfo(info cpu.InfoStat, number int) string {
	clockGHz := float64(info.Mhz) / 1000.0
	return fmt.Sprintf("CPU %d: Model=%s, Clock Rate=%.2f GHz\n", number+1, info.ModelName, clockGHz)
}
