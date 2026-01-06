package slicesx

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func ExamplePartitionInPlace() {
	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	evenSlice, oddSlice := PartitionInPlace(slice, func(i int) bool {
		return i%2 == 0
	})
	fmt.Println(evenSlice)
	fmt.Println(oddSlice)
	// Note: PartitionInPlace doesn't preserve original relative order.
	// Also, the order of the elements in the slices is not guaranteed.

	// Output:
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

func isEven(i int) bool {
	return i%2 == 0
}

func TestPartitionInPlace(t *testing.T) {
	t.Run("empty slice", func(t *testing.T) {
		slice := []int{}
		evenSlice, oddSlice := PartitionInPlace(slice, isEven)
		assert.Len(t, evenSlice, 0)
		assert.Len(t, oddSlice, 0)
	})
	t.Run("all elements match predicate", func(t *testing.T) {
		slice := []int{2, 4, 6, 8, 10}
		evenSlice, oddSlice := PartitionInPlace(slice, isEven)
		assert.ElementsMatch(t, slice, evenSlice)
		assert.Len(t, oddSlice, 0)
	})
	t.Run("all elements do not match predicate", func(t *testing.T) {
		slice := []int{1, 3, 5, 7, 9}
		evenSlice, oddSlice := PartitionInPlace(slice, isEven)
		assert.ElementsMatch(t, slice, oddSlice)
		assert.Len(t, evenSlice, 0)
	})
	t.Run("mixed elements", func(t *testing.T) {
		slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
		evenSlice, oddSlice := PartitionInPlace(slice, isEven)
		assert.ElementsMatch(t, []int{2, 4, 6, 8, 10}, evenSlice)
		assert.ElementsMatch(t, []int{1, 3, 5, 7, 9}, oddSlice)
	})
}

func TestPartitionInPlaceStable(t *testing.T) {
	t.Run("empty slice", func(t *testing.T) {
		slice := []int{}
		evenSlice, oddSlice := PartitionInPlaceStable(slice, isEven)
		assert.Len(t, evenSlice, 0)
		assert.Len(t, oddSlice, 0)
	})
	t.Run("all elements match predicate", func(t *testing.T) {
		slice := []int{2, 4, 6, 8, 10}
		evenSlice, oddSlice := PartitionInPlaceStable(slice, isEven)
		assert.Equal(t, slice, evenSlice)
		assert.Len(t, oddSlice, 0)
	})
	t.Run("all elements do not match predicate", func(t *testing.T) {
		slice := []int{1, 3, 5, 7, 9}
		evenSlice, oddSlice := PartitionInPlaceStable(slice, isEven)
		assert.Equal(t, slice, oddSlice)
		assert.Len(t, evenSlice, 0)
	})
	t.Run("mixed elements", func(t *testing.T) {
		slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
		evenSlice, oddSlice := PartitionInPlaceStable(slice, isEven)
		assert.Equal(t, []int{2, 4, 6, 8, 10}, evenSlice)
		assert.Equal(t, []int{1, 3, 5, 7, 9}, oddSlice)
	})
}

func TestPartitionStable(t *testing.T) {
	t.Run("empty slice", func(t *testing.T) {
		slice := []int{}
		evenSlice, oddSlice := PartitionStable(slice, isEven)
		assert.Len(t, evenSlice, 0)
		assert.Len(t, oddSlice, 0)
	})
	t.Run("all elements match predicate", func(t *testing.T) {
		slice := []int{2, 4, 6, 8, 10}
		evenSlice, oddSlice := PartitionStable(slice, isEven)
		assert.Equal(t, slice, evenSlice)
		assert.Len(t, oddSlice, 0)
	})
	t.Run("all elements do not match predicate", func(t *testing.T) {
		slice := []int{1, 3, 5, 7, 9}
		evenSlice, oddSlice := PartitionStable(slice, isEven)
		assert.Equal(t, slice, oddSlice)
		assert.Len(t, evenSlice, 0)
	})
	t.Run("mixed elements", func(t *testing.T) {
		slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
		evenSlice, oddSlice := PartitionStable(slice, isEven)
		assert.Equal(t, []int{2, 4, 6, 8, 10}, evenSlice)
		assert.Equal(t, []int{1, 3, 5, 7, 9}, oddSlice)
	})
}

