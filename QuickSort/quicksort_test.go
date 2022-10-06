package quicksort

import (
	"reflect"
	"testing"
)

type quickSortTest struct {
	arg      []int
	expected []int
}

var quickSortTests = []quickSortTest{
	{[]int{1, 4, 8, 3, 6, 9, 0}, []int{0, 1, 3, 4, 6, 8, 9}},
	{[]int{-213, 122, 103, 42, 11223, -1243, 1540, 1561, 16020, 1672, 980, 299, 52}, []int{-1243, -213, 42, 52, 103, 122, 299, 980, 1540, 1561, 1672, 11223, 16020}},
	{[]int{1}, []int{1}},
	{[]int{2, 1}, []int{1, 2}},
}

func TestQuickSort(t *testing.T) {
	for _, test := range quickSortTests {
		QuickSort(test.arg)
		if !reflect.DeepEqual(test.arg, test.expected) {
			t.Errorf("Output %v not equal to expected %v", test.arg, test.expected)
		}
	}
}
