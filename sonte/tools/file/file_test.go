package file

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stvmln86/sonte/sonte/tools/test"
)

func TestDelete(t *testing.T) {
	// setup
	orig := test.MockFile(t, "alpha.extn")

	// success
	err := Delete(orig)
	assert.NoFileExists(t, orig)
	assert.NoError(t, err)
}

func TestExists(t *testing.T) {
	// setup
	orig := test.MockFile(t, "alpha.extn")

	// success - true
	okay := Exists(orig)
	assert.True(t, okay)

	// success - false
	okay = Exists("/nope.extn")
	assert.False(t, okay)
}

func TestMove(t *testing.T) {
	// setup
	orig := test.MockFile(t, "alpha.extn")
	dest := strings.Replace(orig, "alpha.extn", "delta.extn", 1)

	// success
	err := Move(orig, dest)
	assert.NoFileExists(t, orig)
	assert.FileExists(t, dest)
	assert.NoError(t, err)
}

func TestRead(t *testing.T) {
	// setup
	orig := test.MockFile(t, "alpha.extn")

	// success
	body, err := Read(orig)
	assert.Equal(t, "Alpha.\n", body)
	assert.NoError(t, err)
}

func TestSearch(t *testing.T) {
	// setup
	orig := test.MockFile(t, "alpha.extn")

	// success - true
	okay, err := Search(orig, "ALPH")
	assert.True(t, okay)
	assert.NoError(t, err)

	// success - false
	okay, err = Search(orig, "NOPE")
	assert.False(t, okay)
	assert.NoError(t, err)
}

func TestWrite(t *testing.T) {
	// setup
	orig := test.MockFile(t, "alpha.extn")

	// success
	err := Write(orig, "Body.\n", 0600)
	test.AssertFile(t, orig, "Body.\n")
	assert.NoError(t, err)
}
