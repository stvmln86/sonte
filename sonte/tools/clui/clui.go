// Package clui implements command-line user interface functions.
package clui

import (
	"fmt"
	"os"
	"strings"
)

// Env returns the value of an existing environment variable.
func Env(name string) (string, error) {
	data, okay := os.LookupEnv(name)
	data = strings.TrimSpace(data)

	switch {
	case !okay:
		return "", fmt.Errorf("cannot find variable %q - does not exist", name)
	case data == "":
		return "", fmt.Errorf("cannot find variable %q - is blank", name)
	default:
		return data, nil
	}
}

// Split returns a command name and argument slice from an argument slice.
func Split(elems []string) (string, []string) {
	switch len(elems) {
	case 0:
		return "", nil
	case 1:
		return strings.ToLower(elems[0]), nil
	default:
		return strings.ToLower(elems[0]), elems[1:]
	}
}
