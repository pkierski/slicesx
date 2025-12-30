package slicesx

import "fmt"

func ExampleSelect() {
	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	evenSlice := Select(slice, func(i int) bool {
		return i%2 == 0
	})
	fmt.Println(evenSlice)
	// Output: [2 4 6 8 10]
}
