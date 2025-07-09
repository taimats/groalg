package compress_test

import (
	"slices"
	"testing"

	"github.com/taimats/groalg/compress"
)

func TestVBEncode(t *testing.T) {
	tests := []struct {
		name  string
		input int64
		want  []byte
	}{
		{"1byte", 5, []byte{133}},
		{"2bytes", 130, []byte{1, 130}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got := compress.VBEncode(int(tt.input))
			if !slices.Equal(got, tt.want) {
				t.Errorf("Not Equal: (got: %08b, want: %v)", got, tt.want)
			}
		})
	}
}

func TestVBDecode(t *testing.T) {
	tests := []struct {
		name  string
		input []byte
		want  int
	}{
		{"1byte", []byte{133}, 5},
		{"2bytes", []byte{1, 130}, 130},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got := compress.VBDecode(tt.input)
			if got != tt.want {
				t.Errorf("Not Equal: (got: %16b, want: %16b)", got, tt.want)
			}
		})
	}
}

func TestCompressNums(t *testing.T) {
	input := []int{3, 5, 20, 21, 23, 76, 77, 78}
	got := compress.CompressNums(input)
	t.Log("before:", input)
	t.Log("after:", got)
}
