package db

import (
	"encoding/json"
	"errors"
	"testing"
)

func TestOutputOptionsValidation(t *testing.T) {
	var tests = []struct {
		testCase string
		opts     OutputOptions
		errMsg   string
	}{
		{
			"valid options",
			OutputOptions{Extension: "mp4"},
			"",
		},
		{
			"missing extension",
			OutputOptions{Extension: ""},
			"extension is required",
		},
	}
	for _, test := range tests {
		err := test.opts.Validate()
		if err == nil {
			err = errors.New("")
		}
		if err.Error() != test.errMsg {
			t.Errorf("wrong error message\nWant %q\nGot  %q", test.errMsg, err.Error())
		}
	}
}

func TestJSONStringEquivalence(t *testing.T) {
	type A struct {
		Width  string `json:"width,omitempty" redis-hash:"width,omitempty"`
		Height string `json:"height,omitempty" redis-hash:"height,omitempty"`
	}
	type B struct {
		Width  int `json:"width,omitempty,string" redis-hash:"width,omitempty"`
		Height int `json:"height,omitempty,string" redis-hash:"height,omitempty"`
	}
	a, _ := json.Marshal(A{"1024", "768"})
	b, _ := json.Marshal(B{1024, 768})
	if string(a) != string(b) {
		t.Fatal("inputs not equal")
	}
}
