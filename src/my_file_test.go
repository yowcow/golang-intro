package hello

import (
	"github.com/stretchr/testify/assert"
	"io"
	"io/ioutil"
	"os"
	"testing"
)

func TestGetFileContent(t *testing.T) {
	assert := assert.New(t)

	b, e := GetFileContent("../data/hoge.txt")

	expected := `hoge: 123
fuga: fuga
props:
  foo: 111
  bar: 222
items:
  - id: 1
    name: foo
  - id: 2
    name: bar
`
	assert.Nil(e)
	assert.Equal([]byte(expected), b)
}

type MyDataItem struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type MyDataProps struct {
	Foo int    `json:"foo"`
	Bar string `json:"bar"`
}

type MyData struct {
	Hoge        int           `json:"hoge"`
	Fuga        string        `json:"fuga"`
	Items       []*MyDataItem `json:"items"`
	Props       *MyDataProps  `json:"props"`
	NonExisting string        `json:"nonexisting"`
}

func TestDecodeYAML(t *testing.T) {
	assert := assert.New(t)

	var data MyData

	b, err := GetFileContent("../data/hoge.txt")
	err = DecodeYAML(b, &data)

	assert.Nil(err)
	assert.Equal(123, data.Hoge)
	assert.Equal("fuga", data.Fuga)
	assert.Equal("", data.NonExisting)

	assert.Equal(1, data.Items[0].ID)
	assert.Equal("foo", data.Items[0].Name)
	assert.Equal(2, data.Items[1].ID)
	assert.Equal("bar", data.Items[1].Name)

	assert.Equal(111, data.Props.Foo)
	assert.Equal("222", data.Props.Bar)
}

func TestEncodeYAML(t *testing.T) {
	assert := assert.New(t)

	dataItems := []*MyDataItem{
		&MyDataItem{1, "foo"},
		&MyDataItem{2, "bar"},
	}
	dataProps := &MyDataProps{111, "222"}
	data := MyData{
		Hoge:  123,
		Fuga:  "fuga",
		Items: dataItems,
		Props: dataProps,
	}

	b, e := EncodeYAML(&data)

	expected := `hoge: 123
fuga: fuga
items:
- id: 1
  name: foo
- id: 2
  name: bar
props:
  foo: 111
  bar: "222"
nonexisting: ""
`
	assert.Nil(e)
	assert.Equal(expected, string(b))
}

func TestDecodeJSON(t *testing.T) {
	assert := assert.New(t)

	data := MyData{}
	b := []byte(`{
	"hoge": 123,
	"fuga": "fuga",
	"props": {
		"foo": 111,
		"bar": "222"
	},
	"items": [
		{"id": 1, "name": "foo"},
		{"id": 2, "name": "bar"}
	]
	}`)

	err := DecodeJSON(b, &data)

	assert.Nil(err)
	assert.Equal(123, data.Hoge)
	assert.Equal("fuga", data.Fuga)
	assert.Equal("", data.NonExisting)

	assert.Equal(1, data.Items[0].ID)
	assert.Equal("foo", data.Items[0].Name)
	assert.Equal(2, data.Items[1].ID)
	assert.Equal("bar", data.Items[1].Name)

	assert.Equal(111, data.Props.Foo)
	assert.Equal("222", data.Props.Bar)
}

func TestEncodeJSON(t *testing.T) {
	assert := assert.New(t)

	data := MyData{
		Hoge: 123,
		Fuga: "fuga",
		Items: []*MyDataItem{
			&MyDataItem{1, "foo"},
			&MyDataItem{2, "bar"},
		},
		Props: &MyDataProps{111, "222"},
	}

	b, e := EncodeJSON(&data)

	expected := `{"hoge":123,"fuga":"fuga","items":[{"id":1,"name":"foo"},{"id":2,"name":"bar"}],"props":{"foo":111,"bar":"222"},"nonexisting":""}`
	assert.Nil(e)
	assert.Equal(expected, string(b))
}

func TestFileCopy(t *testing.T) {
	assert := assert.New(t)

	dir, e := ioutil.TempDir("", "hoge")
	assert.Nil(e)

	defer os.RemoveAll(dir)

	fout, e := os.Create(dir + "/test.data")
	assert.Nil(e)

	fin, e := os.Open("../data/hoge.txt")
	assert.Nil(e)

	io.Copy(fout, fin)

	fout.Close()
	fin.Close()

	actual, _ := ioutil.ReadFile(dir + "/test.data")
	expected, _ := ioutil.ReadFile("../data/hoge.txt")

	assert.Equal(expected, actual)
}
