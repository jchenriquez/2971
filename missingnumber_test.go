package missingnumber

import (
	"bufio"
	"encoding/json"
	"io"
	"os"
	"testing"
)

type Test struct {
	Input  []int `json:"input"`
	Output int   `json:"output"`
}

func TestMissingNumber(t *testing.T) {
	f, err := os.Open("tests.json")

	if err != nil {
		t.Error(err)
		return
	}

	defer f.Close()

	reader := bufio.NewReader(f)
	decoder := json.NewDecoder(reader)
	var tests map[string]Test

	for {
		err = decoder.Decode(&tests)

		if err == nil {
			for name, test := range tests {
				t.Run(name, func(k *testing.T) {
					missing := MissingNumber(test.Input)

					if test.Output != missing {
						k.Errorf("missing number returned %d", missing)
					}
				})
			}
		} else if err == io.EOF {
			break
		} else {
			t.Error(err)
			break
		}
	}

}
