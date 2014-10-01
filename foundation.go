package desalinate

import (
	"fmt"
	"io/ioutil"

	"gopkg.in/yaml.v1"
)

func ReadFile(path string) ([]byte, error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return []byte{}, err
	}
	return data, nil
}

func ReadYaml(str []byte, s interface{}) (interface{}, error) {
	err := yaml.Unmarshal(str, &s)
	fmt.Printf("%#v\n", s)
	if err != nil {
		return nil, err
	}
	return s, nil
}
