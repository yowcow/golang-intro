package hello

import (
	"bufio"
	"bytes"
	"encoding/csv"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReadCSV(t *testing.T) {
	f, _ := os.Open("../_data/hoge.csv")
	r := csv.NewReader(f)

	row1, err := r.Read()

	assert.Nil(t, err)
	assert.Equal(t, "hoge", row1[0])
	assert.Equal(t, "fuga", row1[1])

	row2, err := r.Read()

	assert.Nil(t, err)
	assert.Equal(t, `"foo"`, row2[0])
	assert.Equal(t, " bar", row2[1])
}

func TestWriteCSV(t *testing.T) {
	buf := bytes.Buffer{}
	bufw := bufio.NewWriter(&buf)

	w := csv.NewWriter(bufw)
	w.Comma = '\t'
	w.UseCRLF = true

	w.Write([]string{"hoge", "fuga"})
	w.Write([]string{`"foo"`, `"bar"`})

	assert.Equal(t, "", buf.String())

	w.Flush()

	assert.Equal(t, "hoge\tfuga\r\n\"\"\"foo\"\"\"\t\"\"\"bar\"\"\"\r\n", buf.String())
}
