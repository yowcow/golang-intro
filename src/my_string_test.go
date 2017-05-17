package hello

import "testing"

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
