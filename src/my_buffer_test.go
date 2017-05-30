package hello

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"io"
	"io/ioutil"
	"os"
	"sync"
	"testing"
)

func TestBuffer(t *testing.T) {
	assert := assert.New(t)

	buf := new(bytes.Buffer)
	buf.WriteString("こんにちは")
	buf.Write([]byte(" this"))
	buf.Write([]byte(" is"))
	buf.Write([]byte(" a"))
	buf.Write([]byte(" pen"))

	output := []byte{}

	ret := buf.Next(4)
	for len(ret) > 0 {
		output = append(output, ret...)
		ret = buf.Next(4)
	}

	assert.Equal("こんにちは this is a pen", string(output))
}

var bufPool = sync.Pool{
	New: func() interface{} {
		return new(bytes.Buffer)
	},
}

func myWrite(w io.Writer) {
	b := bufPool.Get().(*bytes.Buffer)
	b.Reset()

	b.WriteString("ほげ")
	b.WriteByte(' ')
	b.WriteString("ふが")
	b.WriteByte(' ')

	w.Write(b.Bytes())

	bufPool.Put(b)
}

// https://golang.org/pkg/sync/#Pool
func TestBufferPool(t *testing.T) {
	assert := assert.New(t)

	outBuf := new(bytes.Buffer)
	myWrite(outBuf)
	myWrite(outBuf)
	myWrite(outBuf)

	assert.Equal("ほげ ふが ほげ ふが ほげ ふが ", string(outBuf.Bytes()))
}

func TestWriteBufferToFile(t *testing.T) {
	assert := assert.New(t)

	f, e := ioutil.TempFile("", "test")
	defer os.Remove(f.Name())

	assert.Equal(nil, e)

	myWrite(f)
	myWrite(f)
	myWrite(f)

	assert.Equal(nil, f.Close())

	content, e := GetFileContent(f.Name())

	assert.Equal("ほげ ふが ほげ ふが ほげ ふが ", string(content))
}
