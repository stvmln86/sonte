// Package path implements file path manipulation functions.
package path

import (
	"path/filepath"
	"sort"
	"strings"
)

// Dire returns a file path's parent directory.
func Dire(orig string) string {
	return filepath.Dir(orig)
}

// Extn returns a file path's extension with a leading dot.
func Extn(orig string) string {
	base := filepath.Base(orig)
	if clip := strings.Index(base, "."); clip != -1 {
		return base[clip:]
	}

	return ""
}

// Glob returns all file paths in a directory matching an extension.
func Glob(dire, extn string) []string {
	glob := filepath.Join(dire, "*"+extn)
	origs, _ := filepath.Glob(glob)
	sort.Strings(origs)
	return origs
}

// Join returns a file path from a directory, name and extension.
func Join(dire, name, extn string) string {
	return filepath.Join(dire, name+extn)
}

// Match returns true if a file path's name contains a prefix.
func Match(orig, pref string) bool {
	name := Name(orig)
	name = strings.ToLower(name)
	pref = strings.ToLower(pref)
	return strings.HasPrefix(name, pref)
}

// Name returns a file path's name.
func Name(orig string) string {
	base := filepath.Base(orig)
	if clip := strings.Index(base, "."); clip != -1 {
		return base[:clip]
	}

	return base
}

// Reextn returns a file path with a different extension.
func Reextn(orig, extn string) string {
	dire := Dire(orig)
	name := Name(orig)
	return Join(dire, name, extn)
}

// Rename returns a file path with a different name.
func Rename(orig, name string) string {
	dire := Dire(orig)
	extn := Extn(orig)
	return Join(dire, name, extn)
}
