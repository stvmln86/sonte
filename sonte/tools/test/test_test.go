package test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAssertFile(t *testing.T) {
	// setup
	dire := t.TempDir()
	orig := filepath.Join(dire, t.Name())
	if err := os.WriteFile(orig, []byte("Body.\n"), 0600); err != nil {
		t.Fatal(err)
	}

	// success
	AssertFile(t, orig, "Body.\n")
}

func TestMockDire(t *testing.T) {
	// success
	dire := MockDire(t)
	assert.DirExists(t, dire)

	// confirm - file bodies
	for base, want := range MockFiles {
		orig := filepath.Join(dire, base)
		bytes, err := os.ReadFile(orig)
		if err != nil {
			t.Fatal(err)
		}

		assert.Equal(t, want, string(bytes))
	}
}

func TestMockFile(t *testing.T) {
	// success
	orig := MockFile(t, "alpha.extn")
	assert.FileExists(t, orig)

	// confirm - file body
	bytes, err := os.ReadFile(orig)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, MockFiles["alpha.extn"], string(bytes))
}
