package filesystem

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"path/filepath"
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

func (fs *LocalFS) AddFile(name string) error {
	file, err := os.Create(name)
	if err != nil {
		return fmt.Errorf("failed to create file %s: %w", name, err)
	}
	defer file.Close()
	return nil
}

func (fs *LocalFS) Copy(src, dst string) error {
	srcInfo, err := os.Stat(src)
	if err != nil {
		return fmt.Errorf("failed to stat source file %s: %w", src, err)
	}
	if srcInfo.IsDir() {
		return fmt.Errorf("source %s is a directory, not a file", src)
	}

	finalDst := dst
	if dstInfo, err := os.Stat(dst); err == nil && dstInfo.IsDir() {
		finalDst = filepath.Join(dst, filepath.Base(src))
	}

	srcFile, err := os.Open(src)
	if err != nil {
		return fmt.Errorf("failed to open source file %s: %w", src, err)
	}
	defer srcFile.Close()

	dstFile, err := os.Create(finalDst)
	if err != nil {
		return fmt.Errorf("failed to create destination file %s: %w", finalDst, err)
	}
	defer dstFile.Close()

	_, err = io.Copy(dstFile, srcFile)
	if err != nil {
		return fmt.Errorf("failed to copy file from %s to %s: %w", src, finalDst, err)
	}

	err = os.Chmod(finalDst, srcInfo.Mode())
	if err != nil {
		return fmt.Errorf("failed to set permissions on destination file %s: %w", finalDst, err)
	}

	err = dstFile.Sync()
	if err != nil {
		return fmt.Errorf("failed to sync destination file %s: %w", finalDst, err)
	}

	return nil
}

func (fs *LocalFS) Move(src, dst string) error {
	finalDst := dst
	if dstInfo, err := os.Stat(dst); err == nil && dstInfo.IsDir() {
		finalDst = filepath.Join(dst, filepath.Base(src))
	}

	if err := os.Rename(src, finalDst); err == nil {
		return nil
	}

	if err := fs.Copy(src, finalDst); err != nil {
		return err
	}

	if err := fs.Delete(src); err != nil {
		os.Remove(finalDst)
		return fmt.Errorf("failed to delete source file %s after copy: %w", src, err)
	}

	return nil
}

func (fs *LocalFS) Rename(src, newName string) error {
	srcDir := filepath.Dir(src)
	dst := filepath.Join(srcDir, newName)
	return os.Rename(src, dst)
}

func (fs *LocalFS) Delete(path string) error {
	err := os.Remove(path)
	if err != nil {
		return fmt.Errorf("failed to delete %s: %w", path, err)
	}
	return nil
}
