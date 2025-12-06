// Package file implements file system handling functions.
package file

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// Delete deletes an existing file.
func Delete(orig string) error {
	if err := os.Remove(orig); err != nil {
		base := filepath.Base(orig)
		return fmt.Errorf("cannot delete file %q - %w", base, err)
	}

	return nil
}

// Exists returns true if a file exists.
func Exists(orig string) bool {
	_, err := os.Stat(orig)
	return !errors.Is(err, os.ErrNotExist)
}

// Move moves an existing file to a new path.
func Move(orig, dest string) error {
	if err := os.Rename(orig, dest); err != nil {
		base := filepath.Base(orig)
		return fmt.Errorf("cannot move file %q - %w", base, err)
	}

	return nil
}

// Read returns an existing file's body as a string.
func Read(orig string) (string, error) {
	bytes, err := os.ReadFile(orig)
	if err != nil {
		base := filepath.Base(orig)
		return "", fmt.Errorf("cannot read file %q - %w", base, err)
	}

	return string(bytes), nil
}

// Search returns true if a file's body contains a substring.
func Search(orig, text string) (bool, error) {
	bytes, err := os.ReadFile(orig)
	if err != nil {
		base := filepath.Base(orig)
		return false, fmt.Errorf("cannot search file %q - %w", base, err)
	}

	text = strings.ToLower(text)
	body := strings.ToLower(string(bytes))
	return strings.Contains(body, text), nil
}

// Write writes a new or existing file with a string.
func Update(orig, body string, mode os.FileMode) error {
	if err := os.WriteFile(orig, []byte(body), mode); err != nil {
		base := filepath.Base(orig)
		return fmt.Errorf("cannot update file %q - %w", base, err)
	}

	return nil
}
