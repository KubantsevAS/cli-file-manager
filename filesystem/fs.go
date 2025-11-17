package filesystem

import "os"

type FileSystem interface {
	Up(dir string) error
	ChangeDirectory(path string) error
	List(dir string) ([]os.DirEntry, error)
	Copy(src, dst string) error
	Move(src, dst string) error
	Delete(path string) error
}
