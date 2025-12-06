package path

import (
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stvmln86/sonte/sonte/tools/test"
)

func TestDire(t *testing.T) {
	// success
	dire := Dire("/dire/name.extn")
	assert.Equal(t, "/dire", dire)
}

func TestExtn(t *testing.T) {
	// success - full extension
	extn := Extn("/dire/name.extn")
	assert.Equal(t, ".extn", extn)

	// success - empty extension
	extn = Extn("/dire/name.")
	assert.Equal(t, ".", extn)

	// success - no extension
	extn = Extn("/dire/name")
	assert.Equal(t, "", extn)
}

func TestGlob(t *testing.T) {
	// setup
	dire := test.MockDire(t)

	// success
	origs := Glob(dire, ".extn")
	assert.Equal(t, []string{
		filepath.Join(dire, "alpha.extn"),
		filepath.Join(dire, "bravo.extn"),
	}, origs)
}

func TestJoin(t *testing.T) {
	// success
	dest := Join("/dire", "name", ".extn")
	assert.Equal(t, "/dire/name.extn", dest)
}

func TestMatch(t *testing.T) {
	// success - true
	okay := Match("/dire/name.extn", "NAME")
	assert.True(t, okay)

	// success - false
	okay = Match("/dire/name.extn", "NOPE")
	assert.False(t, okay)
}

func TestName(t *testing.T) {
	// success - full name
	name := Name("/dire/name.extn")
	assert.Equal(t, "name", name)

	// success - empty name
	name = Name("/dire/name.")
	assert.Equal(t, "name", name)

	// success - no name
	name = Name("/dire/name")
	assert.Equal(t, "name", name)
}
