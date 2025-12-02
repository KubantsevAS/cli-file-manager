package filesystem

import (
	"bufio"
	"compress/gzip"
	"crypto/sha256"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

type LocalFS struct{}

func NewLocalFS() *LocalFS {
	return &LocalFS{}
}

func ensureRegularFile(path string) (os.FileInfo, error) {
	info, err := os.Stat(path)
	if err != nil {
		return nil, fmt.Errorf("failed to stat source file %s: %w", path, err)
	}
	if info.IsDir() {
		return nil, fmt.Errorf("source %s is a directory, not a file", path)
	}
	return info, nil
}

func resolveDestinationPath(src, dst string) string {
	finalDst := dst
	if dstInfo, err := os.Stat(dst); err == nil && dstInfo.IsDir() {
		finalDst = filepath.Join(dst, filepath.Base(src))
	}
	return finalDst
}

func (fs *LocalFS) ChangeDir(path string) (string, error) {
	err := os.Chdir(path)
	if err != nil {
		return "", fmt.Errorf("failed to change directory to %s: %w", path, err)
	}
	return "Changed current directory", nil
}

func (fs *LocalFS) List(dir string) ([]os.DirEntry, error) {
	files, err := os.ReadDir(dir)
	if err != nil {
		return nil, err
	}
	return files, nil
}

func (fs *LocalFS) Read(path string, w io.Writer) (string, error) {
	file, err := os.Open(path)
	if err != nil {
		return "", fmt.Errorf("failed to open file %s: %w", path, err)
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	_, err = io.Copy(w, reader)
	if err != nil {
		return "", fmt.Errorf("failed to read file %s: %w", path, err)
	}

	return fmt.Sprintf("\nSuccessfully read file '%s' content", path), nil
}

func (fs *LocalFS) CreateDir(name string) (string, error) {
	err := os.Mkdir(name, 0755)
	if err != nil {
		if os.IsExist(err) {
			return "", fmt.Errorf("directory %s already exists", name)
		}
		return "", fmt.Errorf("failed to create directory %s: %w", name, err)
	}

	return fmt.Sprintf("Created new directory '%s'", name), nil
}

func (fs *LocalFS) AddFile(name string) (string, error) {
	file, err := os.Create(name)
	if err != nil {
		return "", fmt.Errorf("failed to create file %s: %w", name, err)
	}
	defer file.Close()
	return fmt.Sprintf("Created new file '%s'", name), nil
}

func (fs *LocalFS) Copy(src, dst string) (string, error) {
	srcInfo, err := ensureRegularFile(src)
	if err != nil {
		return "", err
	}

	finalDst := resolveDestinationPath(src, dst)

	srcFile, err := os.Open(src)
	if err != nil {
		return "", fmt.Errorf("failed to open source file %s: %w", src, err)
	}
	defer srcFile.Close()

	dstFile, err := os.Create(finalDst)
	if err != nil {
		return "", fmt.Errorf("failed to create destination file %s: %w", finalDst, err)
	}
	defer dstFile.Close()

	_, err = io.Copy(dstFile, srcFile)
	if err != nil {
		return "", fmt.Errorf("failed to copy file from %s to %s: %w", src, finalDst, err)
	}

	err = os.Chmod(finalDst, srcInfo.Mode())
	if err != nil {
		return "", fmt.Errorf("failed to set permissions on destination file %s: %w", finalDst, err)
	}

	err = dstFile.Sync()
	if err != nil {
		return "", fmt.Errorf("failed to sync destination file %s: %w", finalDst, err)
	}

	return fmt.Sprintf("File '%s' copied to '%s'", src, dst), nil
}

func (fs *LocalFS) Move(src, dst string) (string, error) {
	successMsg := fmt.Sprintf("File '%s' moved to '%s'", src, dst)
	finalDst := resolveDestinationPath(src, dst)

	if err := os.Rename(src, finalDst); err == nil {
		return successMsg, nil
	}

	if _, err := fs.Copy(src, finalDst); err != nil {
		return "", err
	}

	if _, err := fs.Delete(src); err != nil {
		os.Remove(finalDst)
		return "", fmt.Errorf("failed to delete source file %s after copy: %w", src, err)
	}

	return successMsg, nil
}

func (fs *LocalFS) Rename(src, newName string) (string, error) {
	srcDir := filepath.Dir(src)
	dst := filepath.Join(srcDir, newName)

	if err := os.Rename(src, dst); err != nil {
		return "", err
	}

	return fmt.Sprintf("File '%s' renamed to '%s'", src, dst), nil
}

func (fs *LocalFS) Delete(path string) (string, error) {
	err := os.Remove(path)
	if err != nil {
		return "", fmt.Errorf("failed to delete %s: %w", path, err)
	}
	return fmt.Sprintf("File '%s' successfully deleted", path), nil
}

func (fs *LocalFS) Hash(path string) (string, error) {
	file, err := os.Open(path)
	if err != nil {
		return "", fmt.Errorf("failed to open source file %s: %w", path, err)
	}
	defer file.Close()

	hash := sha256.New()

	if _, err := io.Copy(hash, file); err != nil {
		return "", fmt.Errorf("failed to read file %s: %w", path, err)
	}

	hashSum := hash.Sum(nil)
	return fmt.Sprintf("%x", hashSum), nil
}

type gzipTransform func(srcFile *os.File, dstFile *os.File) error

func (fs *LocalFS) transformGzipFile(src, dst string, transform gzipTransform, operationName, successMsg string) (string, error) {
	if _, err := ensureRegularFile(src); err != nil {
		return "", err
	}

	finalDst := resolveDestinationPath(src, dst)

	dstFile, err := os.Create(finalDst)
	if err != nil {
		return "", fmt.Errorf("failed to create destination file %s: %w", finalDst, err)
	}
	defer dstFile.Close()

	srcFile, err := os.Open(src)
	if err != nil {
		return "", fmt.Errorf("failed to open source file %s: %w", src, err)
	}

	if err := transform(srcFile, dstFile); err != nil {
		srcFile.Close()
		return "", fmt.Errorf("failed to %s file %s: %w", operationName, src, err)
	}

	if err := srcFile.Close(); err != nil {
		return "", fmt.Errorf("failed to close source file %s: %w", src, err)
	}

	if err := dstFile.Sync(); err != nil {
		return "", fmt.Errorf("failed to sync destination file %s: %w", finalDst, err)
	}

	if _, err := fs.Delete(src); err != nil {
		os.Remove(finalDst)
		return "", fmt.Errorf("failed to delete source file %s after %s: %w", src, operationName, err)
	}

	return successMsg, nil
}

func (fs *LocalFS) Compress(src, dst string) (string, error) {
	transform := func(srcFile *os.File, dstFile *os.File) error {
		gzipWriter := gzip.NewWriter(dstFile)
		defer gzipWriter.Close()

		if _, err := io.Copy(gzipWriter, srcFile); err != nil {
			return err
		}

		return gzipWriter.Close()
	}

	successMsg := fmt.Sprintf("File %s successfully compressed to %s", src, dst)
	return fs.transformGzipFile(src, dst, transform, "compress", successMsg)
}

func (fs *LocalFS) Decompress(src, dst string) (string, error) {
	transform := func(srcFile *os.File, dstFile *os.File) error {
		gzipReader, err := gzip.NewReader(srcFile)
		if err != nil {
			return fmt.Errorf("failed to create gzip reader: %w", err)
		}
		defer gzipReader.Close()

		if _, err := io.Copy(dstFile, gzipReader); err != nil {
			return err
		}

		return gzipReader.Close()
	}

	successMsg := fmt.Sprintf("File %s successfully decompressed to %s", src, dst)
	return fs.transformGzipFile(src, dst, transform, "decompress", successMsg)
}
