package hello

import (
	"bytes"
	"log"
	"regexp"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLogFormatting(t *testing.T) {
	buf := bytes.Buffer{}
	logger := log.New(&buf, "[hoge] ", log.Lshortfile)
	logger.Println("ほげほげ has happened")

	// [hoge] my_log_test.go:14: ほげほげ has happened
	re := regexp.MustCompile("\\A\\[hoge\\] my_log_test.go:\\d+: ほげほげ has happened\n\\z")

	assert.True(t, re.Match(buf.Bytes()))
}
