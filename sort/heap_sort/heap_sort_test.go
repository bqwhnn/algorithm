package algorithm

import (
	"fmt"
	"testing"
)

func TestHeapSort(t *testing.T) {
	nums := []int{}

	fmt.Println("before heap sort", nums)
	heapSort(nums)
	fmt.Println("after heap sort", nums)
}