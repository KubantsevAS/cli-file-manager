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
	return os.Chdir(path)
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
		return fmt.Errorf("failed to open file: %w", err)
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	_, err = io.Copy(w, reader)
	if err != nil {
		return fmt.Errorf("error reading file: %w", err)
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
			return fmt.Errorf("directory already exists: %s", name)
		}
		return fmt.Errorf("failed to create directory: %w", err)
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
	return nil
}
