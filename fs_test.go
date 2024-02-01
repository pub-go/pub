package main

import (
	"io/fs"
	"path/filepath"
	"testing"
)

func TestFS(t *testing.T) {
	t.Logf("path/../abc/def --> %s", filepath.Join("path", "../abc", "def"))
	fs.WalkDir(views, ".", func(path string, d fs.DirEntry, err error) error {
		t.Logf("path=%v, d=%v, err=%+v", path, d, err)
		return err
	})
}
