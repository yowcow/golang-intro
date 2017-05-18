package hello

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"unicode/utf8"
)

func partialString(s string, from, to int) string {
	return s[from:to]
}

func TestPartialString(t *testing.T) {
	str := "Apple Banana Orange"
	p1 := partialString(str, 0, 5)
	p2 := partialString(str, 6, 6+6)
	if p1 != "Apple" {
		t.Error("Expected Apple but got", p1)
	}
	if p2 != "Banana" {
		t.Error("Expected Banana but got", p2)
	}
}

type MyRuneItem struct {
	R rune
	S int
}

func TestStringByteRune(t *testing.T) {
	assert := assert.New(t)

	s := "foo 日本語"
	b := []byte(s)

	assert.Equal(13, len(s))
	assert.Equal(13, len(b))
	assert.Equal(7, utf8.RuneCount(b))
	assert.Equal(7, utf8.RuneCountInString(s))

	stringToRuneMapArray := func(s string) []MyRuneItem {
		runes := []MyRuneItem{}

		for len(s) > 0 {
			r, size := utf8.DecodeRuneInString(s)
			runes = append(runes, MyRuneItem{r, size})
			s = s[size:]
		}

		return runes
	}

	res := stringToRuneMapArray(s)

	assert.Equal(res[0].R, 'f')
	assert.Equal(res[0].S, 1)
	assert.Equal(res[1].R, 'o')
	assert.Equal(res[1].S, 1)
	assert.Equal(res[2].R, 'o')
	assert.Equal(res[2].S, 1)
	assert.Equal(res[3].R, ' ')
	assert.Equal(res[3].S, 1)
	assert.Equal(res[4].R, '日')
	assert.Equal(res[4].S, 3)
	assert.Equal(res[5].R, '本')
	assert.Equal(res[5].S, 3)
	assert.Equal(res[6].R, '語')
	assert.Equal(res[6].S, 3)
}
