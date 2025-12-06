package clui

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEnv(t *testing.T) {
	// setup
	os.Clearenv()
	os.Setenv("NAME", "\tdata\n")
	os.Setenv("BLANK", "\t\n")

	// success
	data, err := Env("NAME")
	assert.Equal(t, "data", data)
	assert.NoError(t, err)

	// failure - does not exist
	data, err = Env("NOPE")
	assert.Empty(t, data)
	assert.EqualError(t, err, `cannot find variable "NOPE" - does not exist`)

	// failure - is blank
	data, err = Env("BLANK")
	assert.Empty(t, data)
	assert.EqualError(t, err, `cannot find variable "BLANK" - is blank`)
}

func TestSplit(t *testing.T) {
	// success - zero arguments
	name, elems := Split(nil)
	assert.Empty(t, name)
	assert.Nil(t, elems)

	// success - one argument
	name, elems = Split([]string{"name"})
	assert.Equal(t, "name", name)
	assert.Nil(t, elems)

	// success - multiple arguments
	name, elems = Split([]string{"name", "argument"})
	assert.Equal(t, "name", name)
	assert.Equal(t, []string{"argument"}, elems)
}
