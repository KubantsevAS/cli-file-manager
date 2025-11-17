package filesystem

import "os"

type LocalFS struct{}

func NewLocalFS() *LocalFS {
	return &LocalFS{}
}

func (fs *LocalFS) ChangeDirectory(path string) error {
	return nil
}

func (fs *LocalFS) Up(dir string) error {
	return nil
}

func (fs *LocalFS) List(dir string) ([]os.DirEntry, error) {
	files, err := os.ReadDir(dir)
	if err != nil {
		return nil, err
	}
	return files, nil
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
