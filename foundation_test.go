package desalinate

import (
	"fmt"
	"testing"
)

type S struct {
	Env    string
	User   string
	Master string
}

func TestIO(t *testing.T) {
	raw, err := ReadFile("resources/mini.yaml")
	if err != nil {
		t.Errorf("%v\n", err)
	}
	//fmt.Printf("%v\n\n", raw)

	s := S{}
	yml, err := ReadYaml(raw, &s)
	if err != nil {
		t.Errorf("%v\n", err)
	}
	fmt.Printf("%v\n", yml)
}
