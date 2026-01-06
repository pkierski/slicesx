package slicesx

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func ExampleSelect() {
	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	evenSlice := Select(slice, isEven)
	fmt.Println(evenSlice)
	// Output: [2 4 6 8 10]
}

func TestSelect(t *testing.T) {
	t.Run("empty slice", func(t *testing.T) {
		slice := []int{}
		evenSlice := Select(slice, isEven)
		assert.Len(t, evenSlice, 0)
	})
	t.Run("all elements match predicate", func(t *testing.T) {
		slice := []int{2, 4, 6, 8, 10}
		evenSlice := Select(slice, isEven)
		assert.Equal(t, slice, evenSlice)
	})
	t.Run("all elements do not match predicate", func(t *testing.T) {
		slice := []int{1, 3, 5, 7, 9}
		evenSlice := Select(slice, isEven)
		assert.Len(t, evenSlice, 0)
	})
	t.Run("mixed elements", func(t *testing.T) {
		slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
		evenSlice := Select(slice, isEven)
		assert.Equal(t, []int{2, 4, 6, 8, 10}, evenSlice)
	})
}
