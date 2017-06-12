package hello

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSimpleMap(t *testing.T) {
	assert := assert.New(t)

	m := map[string]int{"apple": 1, "banana": 2}

	m["orange"] = 3

	assert.Equal(1, m["apple"])
	assert.Equal(2, m["banana"])
	assert.Equal(3, m["orange"])
}

func TestArrayOfMaps(t *testing.T) {
	assert := assert.New(t)

	a := []map[string]int{}

	a = append(a, map[string]int{"apple": 1})
	a = append(a, map[string]int{"banana": 2})

	assert.Equal(map[string]int{"apple": 1}, a[0])
	assert.Equal(map[string]int{"banana": 2}, a[1])
}

func TestMapOfMaps(t *testing.T) {
	assert := assert.New(t)

	m := map[string]map[string]int{}

	m["apple"] = map[string]int{"apple": 1}
	m["banana"] = map[string]int{"banana": 2}

	assert.Equal(1, m["apple"]["apple"])
	assert.Equal(2, m["banana"]["banana"])
}

func TestMapKeyLength(t *testing.T) {
	assert := assert.New(t)

	data := map[string]interface{}{}
	data["a"] = 1
	data["b"] = "2"
	data["c"] = 3.0

	assert.Equal(3, len(data))
	assert.Equal(data["a"], 1)
	assert.Equal(data["b"], "2")
	assert.Equal(data["c"], float64(3))
}
