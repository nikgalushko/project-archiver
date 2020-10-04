package pa

import (
	"testing"
)

type storage struct{}

func (s storage) Register(title string) error {
	return nil
}

func TestDirToTitle(t *testing.T) {
	dir := t.TempDir()

	archiver := projectArchiver{
		workDir: dir,
		db:      storage{},
	}

	title, err := archiver.dirToTitle(dir)
	if err != nil {
		t.Fatal(err)
	}

	if title == "" {
		t.Fatalf("title should not be empty")
	}
}
