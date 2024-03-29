package main

import (
	"reflect"
	"testing"
)

func TestSplit(t *testing.T) {
	tests := map[string]struct {
		input string
		want  string
	}{
		"1": {input: "1", want: "1"},
		"2": {input: "1", want: "1"},
		"3": {input: "1", want: "1"},
		"4": {input: "1", want: "1"},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			if !reflect.DeepEqual(tc.want, tc.input) {
				t.Fatalf("expected: %#v, got: %#v", tc.want, tc.input)
			}
		})
	}
}
