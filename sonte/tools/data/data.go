// Package data implements data conversion and sanitisation functions.
package data

import (
	"path/filepath"
	"strings"
	"unicode"
)

// Body returns a whitespace-trimmed body with a trailing newline.
func Body(body string) string {
	return strings.TrimSpace(body) + "\n"
}

// Extn returns a lowercase file extension with a leading dot.
func Extn(extn string) string {
	extn = strings.ToLower(extn)
	extn = strings.TrimSpace(extn)
	return "." + strings.TrimPrefix(extn, ".")
}

// Name returns a lowercase alphanumeric-with-dashes file name.
func Name(name string) string {
	var runes []rune
	for _, rune := range strings.ToLower(name) {
		switch {
		case unicode.IsLetter(rune) || unicode.IsNumber(rune):
			runes = append(runes, rune)
		case unicode.IsSpace(rune) || rune == '-' || rune == '_':
			runes = append(runes, '-')
		}
	}

	return strings.Trim(string(runes), "-")
}

// Path returns a clean file path.
func Path(path string) string {
	path = strings.TrimSpace(path)
	return filepath.Clean(path)
}
