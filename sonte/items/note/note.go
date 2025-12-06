// Package note implements the Note type and methods.
package note

import (
	"os"

	"github.com/stvmln86/sonte/sonte/tools/data"
	"github.com/stvmln86/sonte/sonte/tools/file"
	"github.com/stvmln86/sonte/sonte/tools/path"
)

// Note is a single plaintext note file in a Book.
type Note struct {
	Orig string
	Mode os.FileMode
}

// New returns a new Note.
func New(orig string, mode os.FileMode) *Note {
	return &Note{Orig: orig, Mode: mode}
}

// Delete deletes the Note if it exists.
func (n *Note) Delete() error {
	return file.Delete(n.Orig)
}

// Exists returns true if the Note exists.
func (n *Note) Exists() bool {
	return file.Exists(n.Orig)
}

// Match returns true if the Note's name contains a prefix.
func (n *Note) Match(pref string) bool {
	return path.Match(n.Orig, pref)
}

// Name returns the Note's name.
func (n *Note) Name() string {
	name := path.Name(n.Orig)
	return data.Name(name)
}

// Read returns the Note's body as string.
func (n *Note) Read() (string, error) {
	body, err := file.Read(n.Orig)
	return data.Body(body), err
}

// Search returns true if the Note's body contains a substring.
func (n *Note) Search(text string) (bool, error) {
	return file.Search(n.Orig, text)
}

// Write overwrites the Note with a string.
func (n *Note) Write(body string) error {
	body = data.Body(body)
	return file.Write(n.Orig, body, n.Mode)
}
