package selection_sort

import "testing"

func TestAsc(t *testing.T) {
	unsorted := []int{11, 7, 33, 9, 12, 54, 3, 27, 41, 99, 2, 87, 0}
	sorted := []int{0, 2, 3, 7, 9, 11, 12, 27, 33, 41, 54, 87, 99}
	Asc(unsorted)
	for i, v := range unsorted {
		if sorted[i] != v {
			t.Error("Unexpected ordering in selection sort result")
		}
	}
}

func TestDesc(t *testing.T) {
	unsorted := []int{11, 7, 33, 9, 12, 54, 0, 3, 27, 41, 99, 2, 87}
	sorted := []int{99, 87, 54, 41, 33, 27, 12, 11, 9, 7, 3, 2, 0}
	Desc(unsorted)
	for i, v := range unsorted {
		if sorted[i] != v {
			t.Error("Unexpected ordering in selection sort result")
		}
	}
}
