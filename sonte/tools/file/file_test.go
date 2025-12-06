package file

import (
	"path/filepath"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stvmln86/sonte/sonte/tools/test"
)

func TestCreate(t *testing.T) {
	// setup
	dire := t.TempDir()
	dest := filepath.Join(dire, "name.extn")

	// success
	err := Create(dest, "Body.\n", 0600)
	test.AssertFile(t, dest, "Body.\n")
	assert.NoError(t, err)

	// failure - already exists
	err = Create(dest, "Body.\n", 0600)
	assert.EqualError(t, err, `cannot create file "name.extn" - already exists`)
}

func TestDelete(t *testing.T) {
	// setup
	orig := test.MockFile(t, "alpha.extn")

	// success
	err := Delete(orig)
	assert.NoFileExists(t, orig)
	assert.NoError(t, err)

	// failure - does not exist
	err = Delete(orig)
	assert.EqualError(t, err, `cannot delete file "alpha.extn" - does not exist`)
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

func TestRead(t *testing.T) {
	// setup
	orig := test.MockFile(t, "alpha.extn")

	// success
	body, err := Read(orig)
	assert.Equal(t, "Alpha.\n", body)
	assert.NoError(t, err)

	// failure - does not exist
	body, err = Read("/nope.extn")
	assert.Empty(t, body)
	assert.EqualError(t, err, `cannot read file "nope.extn" - does not exist`)
}

func TestReextn(t *testing.T) {
	// setup
	orig := test.MockFile(t, "alpha.extn")
	dest := strings.Replace(orig, "alpha.extn", "alpha.test", 1)

	// success
	err := Reextn(orig, ".test")
	assert.NoFileExists(t, orig)
	assert.FileExists(t, dest)
	assert.NoError(t, err)

	// failure - does not exist
	err = Reextn("/nope.extn", ".test")
	assert.EqualError(t, err, `cannot move file "nope.extn" - does not exist`)

	// failure - destination exists
	err = Reextn(dest, ".test")
	assert.EqualError(t, err, `cannot move file "alpha.test" - destination exists`)
}

func TestRename(t *testing.T) {
	// setup
	orig := test.MockFile(t, "alpha.extn")
	dest := strings.Replace(orig, "alpha.extn", "test.extn", 1)

	// success
	err := Rename(orig, "test")
	assert.NoFileExists(t, orig)
	assert.FileExists(t, dest)
	assert.NoError(t, err)

	// failure - does not exist
	err = Rename("/nope.extn", "test")
	assert.EqualError(t, err, `cannot move file "nope.extn" - does not exist`)

	// failure - destination exists
	err = Rename(dest, "test")
	assert.EqualError(t, err, `cannot move file "test.extn" - destination exists`)
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

	// failure - does not exist
	okay, err = Search("/nope.extn", "NOPE")
	assert.False(t, okay)
	assert.EqualError(t, err, `cannot search file "nope.extn" - does not exist`)
}

func TestUpdate(t *testing.T) {
	// setup
	orig := test.MockFile(t, "alpha.extn")

	// success
	err := Update(orig, "Body.\n", 0600)
	test.AssertFile(t, orig, "Body.\n")
	assert.NoError(t, err)

	// failure - does not exist
	err = Update("/nope.extn", "Body.\n", 0600)
	assert.EqualError(t, err, `cannot update file "nope.extn" - does not exist`)
}
