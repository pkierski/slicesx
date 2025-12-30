package slicesx

import "fmt"

func ExampleMinIndex() {
	slice := []int{1, 2, 3, 4, 5, -6, 7, 8, 9, 10, -6}
	minIndex := MinIndex(slice)
	fmt.Println(minIndex)
	// Output: 5
}

func ExampleMinIndexFunc() {
	slice := []string{"foo", "x", "bar", "baz", "qux", "a"}
	minIndex := MinIndexFunc(slice, func(a, b string) int {
		return len(a) - len(b)
	})
	fmt.Println(minIndex)
	// Output: 1
}

func ExampleMaxIndex() {
	slice := []int{1, 2, 3, 4, 5, -6, 7, 8, 9, 10, -6}
	maxIndex := MaxIndex(slice)
	fmt.Println(maxIndex)
	// Output: 9
}

func ExampleMaxIndexFunc() {
	slice := []string{"foo", "x", "bar", "bazbaz", "quxqux", "a"}
	maxIndex := MaxIndexFunc(slice, func(a, b string) int {
		return len(a) - len(b)
	})
	fmt.Println(maxIndex)
	// Output: 3
}

func ExampleMinMaxIndex() {
	slice := []int{1, 2, 3, 4, 5, -6, 7, 8, 9, 10, -6}
	minIndex, maxIndex := MinMaxIndex(slice)
	fmt.Println(minIndex, maxIndex)
	// Output: 5 9
}

func ExampleMinMaxIndexFunc() {
	slice := []string{"foo", "x", "bar", "bazbaz", "quxqux", "a"}
	minIndex, maxIndex := MinMaxIndexFunc(slice, func(a, b string) int {
		return len(a) - len(b)
	})
	fmt.Println(minIndex, maxIndex)
	// Output: 1 3
}
