package pkg

import (
	"testing"
)

func Test_RunCase(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
		{"demo case"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			RunCase("../examples/post_json.yaml")
		})
	}
}
