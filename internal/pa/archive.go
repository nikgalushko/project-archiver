package pa

import (
	"errors"
	"os"
	"path/filepath"
)

type Target string

const (
	Local  Target = "local"
	Github Target = "github"
)

func (pa projectArchiver) Archive(dir string, target Target) error {
	switch target {
	case Local:
		return pa.local(dir)
	case Github:
		return errors.New("unimplemented")
	default:
		return errors.New("unknow target")
	}
}

func (pa projectArchiver) local(dir string) error {
	// make zip
	_, err := pa.makeZip(dir)
	if err != nil {
		return err
	}
	// move to local dir
	// store tags to local db

	return nil
}

func (pa projectArchiver) makeZip(dir string) (string, error) {
	return "", nil
}

func (pa projectArchiver) dirToTitle(path string) (string, error) {
	stat, err := os.Stat(path)
	if err != nil {
		return "", err
	}

	if !stat.IsDir() {
		return "", errors.New("is not a directory")
	}

	abs, err := filepath.Abs(path)
	if err != nil {
		return "", err
	}

	_, file := filepath.Split(abs)

	err = pa.db.Register(file)
	if err != nil {
		return "", err
	}

	return file, nil
}
