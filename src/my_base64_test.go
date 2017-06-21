package hello

import (
	"bytes"
	"encoding/base64"
	"io"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStreamingEncode(t *testing.T) {
	buf := &bytes.Buffer{}
	w := base64.NewEncoder(base64.StdEncoding, buf)
	input := []string{"0", "0", "0", "0", "0", "0"}

	for _, v := range input {
		w.Write([]byte(v))
	}

	assert.Nil(t, w.Close())
	assert.Equal(t, "MDAwMDAw", buf.String())
}

func TestStreamingDecode(t *testing.T) {
	buf := &bytes.Buffer{}
	r := base64.NewDecoder(base64.StdEncoding, buf)
	input := []string{"M", "D", "A", "w", "M", "D", "A", "w"}

	for _, v := range input {
		buf.Write([]byte(v))
	}

	outbuf := &bytes.Buffer{}
	io.Copy(outbuf, r)

	assert.Equal(t, "000000", outbuf.String())
}

func TestURLSafeEncode(t *testing.T) {
	buf := &bytes.Buffer{}
	w := base64.NewEncoder(base64.URLEncoding, buf)
	w.Write([]byte("\xfb"))
	assert.Nil(t, w.Close())

	assert.Equal(t, "-w==", buf.String())
}

func TestURLSafeDecode(t *testing.T) {
	buf := &bytes.Buffer{}
	r := base64.NewDecoder(base64.URLEncoding, buf)
	buf.Write([]byte("-w=="))

	outbuf := &bytes.Buffer{}
	io.Copy(outbuf, r)

	assert.Equal(t, "\xfb", outbuf.String())
}
