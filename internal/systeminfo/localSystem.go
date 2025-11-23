package systeminfo

import (
	"fmt"
	"os"
	"runtime"
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
	// TODO: Implement EOL
	return "", fmt.Errorf("EOL not yet implemented")
}

func (s *LocalSystem) CPUs() (string, error) {
	// TODO: Implement CPUs
	return "", fmt.Errorf("CPUs not yet implemented")
}

func (s *LocalSystem) Username() (string, error) {
	// TODO: Implement Username
	return "", fmt.Errorf("Username not yet implemented")
}

func (s *LocalSystem) Architecture() (string, error) {
	return runtime.GOARCH, nil
}
