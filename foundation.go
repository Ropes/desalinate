package desalinate

import (
	"io/ioutil"

	"gopkg.in/yaml.v1"
)

type S struct{}

func ReadFile(path string) ([]byte, error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return []byte{}, err
	}
	return data, nil
}

func ReadYaml(str []byte) (interface{}, error) {
	s := S{}
	err := yaml.Unmarshal(str, &s)
	if err != nil {
		return nil, err
	}
	return s, nil
}
