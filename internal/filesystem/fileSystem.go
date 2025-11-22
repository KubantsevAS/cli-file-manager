package filesystem

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

type LocalFS struct{}

func NewLocalFS() *LocalFS {
	return &LocalFS{}
}

func (fs *LocalFS) ChangeDir(path string) error {
	err := os.Chdir(path)
	if err != nil {
		return fmt.Errorf("failed to change directory to %s: %w", path, err)
	}
	return nil
}

func (fs *LocalFS) List(dir string) ([]os.DirEntry, error) {
	files, err := os.ReadDir(dir)
	if err != nil {
		return nil, err
	}
	return files, nil
}

func (fs *LocalFS) Read(path string, w io.Writer) error {
	file, err := os.Open(path)
	if err != nil {
		return fmt.Errorf("failed to open file %s: %w", path, err)
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	_, err = io.Copy(w, reader)
	if err != nil {
		return fmt.Errorf("failed to read file %s: %w", path, err)
	}

	return nil
}

func (fs *LocalFS) AddFile(path string) error {
	return nil
}

func (fs *LocalFS) CreateDir(name string) error {
	err := os.Mkdir(name, 0755)
	if err != nil {
		if os.IsExist(err) {
			return fmt.Errorf("directory %s already exists", name)
		}
		return fmt.Errorf("failed to create directory %s: %w", name, err)
	}

	return nil
}

func (fs *LocalFS) Copy(src, dst string) error {
	return nil
}

func (fs *LocalFS) Move(src, dst string) error {
	return nil
}

func (fs *LocalFS) Delete(path string) error {
	err := os.Remove(path)
	if err != nil {
		return fmt.Errorf("failed to delete %s: %w", path, err)
	}
	return nil
}
