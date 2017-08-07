package hello

import (
	"encoding/json"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

func GetFileContent(filename string) ([]byte, error) {
	return ioutil.ReadFile(filename)
}

func DecodeYAML(in []byte, out interface{}) error {
	return yaml.Unmarshal(in, out)
}

func EncodeYAML(in interface{}) ([]byte, error) {
	return yaml.Marshal(in)
}

func DecodeJSON(in []byte, out interface{}) error {
	return json.Unmarshal(in, out)
}

func EncodeJSON(in interface{}) ([]byte, error) {
	return json.Marshal(in)
}
