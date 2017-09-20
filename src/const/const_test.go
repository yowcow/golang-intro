package myconst

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	Level1 = 1 + iota
	Level2
	Level3
)

const (
	Mode1 = 1 << iota
	Mode2
	Mode3
)

func TestLevels(t *testing.T) {
	assert.Equal(t, 1, Level1)
	assert.Equal(t, 2, Level2)
	assert.Equal(t, 3, Level3)
}

func TestModes(t *testing.T) {
	assert.Equal(t, 1, Mode1)
	assert.Equal(t, 2, Mode2)
	assert.Equal(t, 4, Mode3)

	assert.Equal(t, 0, Mode1&Mode2)
	assert.Equal(t, 3, Mode1|Mode2)
	assert.Equal(t, 7, Mode1|Mode2|Mode3)

	assert.Equal(t, "111", fmt.Sprintf("%b", Mode1|Mode2|Mode3))
}
