package slicesx

import (
	"fmt"
)

func ExamplePartitionInPlace() {
	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	evenSlice, oddSlice := PartitionInPlace(slice, func(i int) bool {
		return i%2 == 0
	})
	fmt.Println(slice)
	fmt.Println(evenSlice)
	fmt.Println(oddSlice)
	// Note: PartitionInPlace doesn't preserve original relative order.

	// Output:
	// [10 2 8 4 6 5 7 3 9 1]
	// [10 2 8 4 6]
	// [5 7 3 9 1]
}

func ExamplePartitionInPlaceStable() {
	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	evenSlice, oddSlice := PartitionInPlaceStable(slice, func(i int) bool {
		return i%2 == 0
	})
	fmt.Println(slice)
	fmt.Println(evenSlice)
	fmt.Println(oddSlice)
	// Note: PartitionInPlaceStable preserves original relative order.

	// Output:
	// [2 4 6 8 10 1 3 5 7 9]
	// [2 4 6 8 10]
	// [1 3 5 7 9]
}

func ExamplePartitionStable() {
	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	evenSlice, oddSlice := PartitionStable(slice, func(i int) bool {
		return i%2 == 0
	})
	fmt.Println(slice)
	fmt.Println(evenSlice)
	fmt.Println(oddSlice)
	// Note: PartitionStable creates new slices, so original slice is preserved.

	// Output:
	// [1 2 3 4 5 6 7 8 9 10]
	// [2 4 6 8 10]
	// [1 3 5 7 9]
}
