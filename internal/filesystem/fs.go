package filesystem

import (
	"io"
	"os"
)

type FileSystem interface {
	ChangeDir(path string) (string, error)
	List(dir string) ([]os.DirEntry, error)
	Read(path string, w io.Writer) (string, error)
	CreateDir(name string) (string, error)
	AddFile(path string) (string, error)
	Copy(src, dst string) (string, error)
	Move(src, dst string) (string, error)
	Rename(src, newName string) (string, error)
	Delete(path string) (string, error)
	Hash(path string) (string, error)
	Compress(src, dst string) (string, error)
	Decompress(src, dst string) (string, error)
}
