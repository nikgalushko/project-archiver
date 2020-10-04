package pa

import (
	"os"
	"path/filepath"
)

type store interface {
	Register(title string) error
}

type projectArchiver struct {
	workDir string
	db      store
}

func New() (projectArchiver, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return projectArchiver{}, err
	}

	dir := filepath.Join(home, ".projects_archive")

	_, err = os.Stat(dir)
	if os.IsNotExist(err) {
		err = os.Mkdir(dir, os.ModeDir)
		if err != nil {
			return projectArchiver{}, err
		}
	}

	return projectArchiver{workDir: dir}, nil
}
