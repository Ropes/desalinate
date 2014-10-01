package desalinate

import (
	"fmt"
	"testing"
)

func TestIO(t *testing.T) {
	raw, err := ReadFile("data.sls")
	if err != nil {
		t.Errorf("%v\n", err)
	}
	//fmt.Printf("%v\n\n", raw)

	yml, err := ReadYaml(raw)
	if err != nil {
		t.Errorf("%v\n", err)
	}
	fmt.Printf("%v\n", yml)
}
