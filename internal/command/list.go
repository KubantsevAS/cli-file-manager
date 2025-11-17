package command

import (
	"os"
	"sort"

	"cli/file-manager/filesystem"
)

type ListCommand struct {
	FS filesystem.FileSystem
}

func NewListCommand(fs filesystem.FileSystem) *ListCommand {
	return &ListCommand{FS: fs}
}

func (c *ListCommand) Execute(dir string) ([]os.DirEntry, error) {
	files, err := c.FS.List(dir)
	if err != nil {
		return nil, err
	}

	sort.Slice(files, func(i, j int) bool {
		isDirI := files[i].IsDir()
		isDirJ := files[j].IsDir()

		if isDirI != isDirJ {
			return isDirI
		}

		return files[i].Name() < files[j].Name()
	})

	return files, nil
}
