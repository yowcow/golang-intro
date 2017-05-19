package hello

import (
	"github.com/stretchr/testify/assert"
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
	Id   int    `json:"id"`
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

func TestDecodeYaml(t *testing.T) {
	assert := assert.New(t)

	var data MyData

	b, err := GetFileContent("../data/hoge.txt")
	err = DecodeYaml(b, &data)

	assert.Nil(err)
	assert.Equal(123, data.Hoge)
	assert.Equal("fuga", data.Fuga)
	assert.Equal("", data.NonExisting)

	assert.Equal(1, data.Items[0].Id)
	assert.Equal("foo", data.Items[0].Name)
	assert.Equal(2, data.Items[1].Id)
	assert.Equal("bar", data.Items[1].Name)

	assert.Equal(111, data.Props.Foo)
	assert.Equal("222", data.Props.Bar)
}

func TestEncodeYaml(t *testing.T) {
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

	b, e := EncodeYaml(&data)

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

func TestDecodeJson(t *testing.T) {
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

	err := DecodeJson(b, &data)

	assert.Nil(err)
	assert.Equal(123, data.Hoge)
	assert.Equal("fuga", data.Fuga)
	assert.Equal("", data.NonExisting)

	assert.Equal(1, data.Items[0].Id)
	assert.Equal("foo", data.Items[0].Name)
	assert.Equal(2, data.Items[1].Id)
	assert.Equal("bar", data.Items[1].Name)

	assert.Equal(111, data.Props.Foo)
	assert.Equal("222", data.Props.Bar)
}

func TestEncodeJson(t *testing.T) {
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

	b, e := EncodeJson(&data)

	expected := `{"hoge":123,"fuga":"fuga","items":[{"id":1,"name":"foo"},{"id":2,"name":"bar"}],"props":{"foo":111,"bar":"222"},"nonexisting":""}`
	assert.Nil(e)
	assert.Equal(expected, string(b))
}
