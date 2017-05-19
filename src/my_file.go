package hello

import (
	"encoding/json"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

func GetFileContent(filename string) ([]byte, error) {
	return ioutil.ReadFile(filename)
}

func DecodeYaml(in []byte, out interface{}) error {
	return yaml.Unmarshal(in, out)
}

func EncodeYaml(in interface{}) ([]byte, error) {
	return yaml.Marshal(in)
}

func DecodeJson(in []byte, out interface{}) error {
	return json.Unmarshal(in, out)
}

func EncodeJson(in interface{}) ([]byte, error) {
	return json.Marshal(in)
}
