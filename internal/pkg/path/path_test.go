package path

import (
	"reflect"
	"testing"
)

func TestParse(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name string
		args args
		want Path
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Parse(tt.args.path); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Parse() = %v, want %v", got, tt.want)
			}
		})
	}
}
