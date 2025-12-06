// Package file implements file system handling functions.
package file

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/stvmln86/sonte/sonte/tools/path"
)

// Create creates a new file with a body string.
func Create(dest, body string, mode os.FileMode) error {
	if Exists(dest) {
		base := filepath.Base(dest)
		return fmt.Errorf("cannot create file %q - already exists", base)
	}

	if err := os.WriteFile(dest, []byte(body), mode); err != nil {
		base := filepath.Base(dest)
		return fmt.Errorf("cannot create file %q - %w", base, err)
	}

	return nil
}

// Delete deletes an existing file.
func Delete(orig string) error {
	if !Exists(orig) {
		base := filepath.Base(orig)
		return fmt.Errorf("cannot delete file %q - does not exist", base)
	}

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

// Read returns an existing file's body as a string.
func Read(orig string) (string, error) {
	if !Exists(orig) {
		base := filepath.Base(orig)
		return "", fmt.Errorf("cannot read file %q - does not exist", base)
	}

	bytes, err := os.ReadFile(orig)
	if err != nil {
		base := filepath.Base(orig)
		return "", fmt.Errorf("cannot read file %q - %w", base, err)
	}

	return string(bytes), nil
}

// Reextn moves an existing file to a new extension.
func Reextn(orig, extn string) error {
	if !Exists(orig) {
		base := filepath.Base(orig)
		return fmt.Errorf("cannot move file %q - does not exist", base)
	}

	dest := path.Reextn(orig, extn)
	if Exists(dest) {
		base := filepath.Base(orig)
		return fmt.Errorf("cannot move file %q - destination exists", base)
	}

	if err := os.Rename(orig, dest); err != nil {
		base := filepath.Base(orig)
		return fmt.Errorf("cannot move file %q - %w", base, err)
	}

	return nil
}

// Rename moves an existing file to a new name.
func Rename(orig, name string) error {
	if !Exists(orig) {
		base := filepath.Base(orig)
		return fmt.Errorf("cannot move file %q - does not exist", base)
	}

	dest := path.Rename(orig, name)
	if Exists(dest) {
		base := filepath.Base(orig)
		return fmt.Errorf("cannot move file %q - destination exists", base)
	}

	if err := os.Rename(orig, dest); err != nil {
		base := filepath.Base(orig)
		return fmt.Errorf("cannot move file %q - %w", base, err)
	}

	return nil
}

// Search returns true if a file's body contains a substring.
func Search(orig, text string) (bool, error) {
	if !Exists(orig) {
		base := filepath.Base(orig)
		return false, fmt.Errorf("cannot search file %q - does not exist", base)
	}

	bytes, err := os.ReadFile(orig)
	if err != nil {
		base := filepath.Base(orig)
		return false, fmt.Errorf("cannot search file %q - %w", base, err)
	}

	text = strings.ToLower(text)
	body := strings.ToLower(string(bytes))
	return strings.Contains(body, text), nil
}

// Update overwrites an existing file with a string.
func Update(orig, body string, mode os.FileMode) error {
	if !Exists(orig) {
		base := filepath.Base(orig)
		return fmt.Errorf("cannot update file %q - does not exist", base)
	}

	if err := os.WriteFile(orig, []byte(body), mode); err != nil {
		base := filepath.Base(orig)
		return fmt.Errorf("cannot update file %q - %w", base, err)
	}

	return nil
}
