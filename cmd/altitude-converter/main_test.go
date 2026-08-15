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

func TestReadFileRejectsCSVWithoutDataRows(t *testing.T) {
	path := filepath.Join(t.TempDir(), "input.csv")
	if err := os.WriteFile(path, []byte("value,unit\n"), 0o600); err != nil {
		t.Fatal(err)
	}
	if _, err := readFile(path, "m"); err == nil {
		t.Fatal("CSV without data rows unexpectedly succeeded")
	}
}
