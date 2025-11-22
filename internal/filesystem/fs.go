package filesystem

import (
	"io"
	"os"
)

type FileSystem interface {
	ChangeDirectory(path string) error
	List(dir string) ([]os.DirEntry, error)
	Read(path string, w io.Writer) error
	Copy(src, dst string) error
	Move(src, dst string) error
	Delete(path string) error
}
