package filesystem

import (
	"io"
	"os"
)

type FileSystem interface {
	ChangeDir(path string) error
	List(dir string) ([]os.DirEntry, error)
	Read(path string, w io.Writer) error
	CreateDir(name string) error
	AddFile(path string) error
	Copy(src, dst string) error
	Move(src, dst string) error
	Rename(src, newName string) error
	Delete(path string) error
}
