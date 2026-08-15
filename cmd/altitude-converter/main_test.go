package main

import (
	"os"
	"path/filepath"
	"testing"
)

func TestReadFileRejectsUnknownTargetForEmptyTextFile(t *testing.T) {
	path := filepath.Join(t.TempDir(), "empty.txt")
	if err := os.WriteFile(path, nil, 0o600); err != nil {
		t.Fatal(err)
	}
	if _, err := readFile(path, "unknown"); err == nil {
		t.Fatal("unknown target unit unexpectedly succeeded for empty input")
	}
}
