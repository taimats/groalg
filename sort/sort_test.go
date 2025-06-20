package sort

import "testing"

func TestQuickSort(t *testing.T) {
	tests := []struct {
		name string
		data []int
		want []int
	}{
		{
			"正常系:01",
			[]int{4, 6, 1, 5, 3, 8, 7, 2},
			sequence(8),
		},
		{
			"正常系:02",
			[]int{1, 2, 3, 4, 5, 6, 7},
			sequence(7),
		},
		{
			"正常系:03",
			[]int{1, 2},
			sequence(2),
		},
		{
			"正常系:04",
			[]int{2, 2, 1, 1, -1, -1, -4, -4, -10, 10},
			[]int{-10, -4, -4, -1, -1, 1, 1, 2, 2, 10},
		},
		{
			"異常系:01:処理なし",
			[]int{4},
			[]int{4},
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			quickSort(tt.data)
			if !equalSlices(tt.data, tt.want) {
				t.Errorf("not equal: (got: %+v, want: %+v)", tt.data, tt.want)
			}
		})
	}
}

func equalSlices(got []int, want []int) bool {
	for i, num := range got {
		if num != want[i] {
			return false
		}
	}
	return true
}

func sequence(max int) []int {
	seq := make([]int, 0, max)
	for i := range max {
		seq = append(seq, i+1)
	}
	return seq
}
