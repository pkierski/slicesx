package slicesx

import (
	"cmp"

	"golang.org/x/exp/constraints"
)

// MinIndex returns the index of the minimum element in the slice.
func MinIndex[S ~[]E, E constraints.Ordered](s S) int {
	return MinIndexFunc(s, cmp.Compare)
}

// MinIndexFunc returns the index of the minimum element in the slice.
func MinIndexFunc[S ~[]E, E any](s S, cmp func(E, E) int) int {
	if len(s) == 0 {
		return -1
	}
	minIndex := 0
	for i := range s {
		if cmp(s[i], s[minIndex]) < 0 {
			minIndex = i
		}
	}
	return minIndex
}

// MaxIndex returns the index of the maximum element in the slice.
func MaxIndex[S ~[]E, E constraints.Ordered](s S) int {
	return MaxIndexFunc(s, cmp.Compare)
}

// MaxIndexFunc returns the index of the maximum element in the slice.
func MaxIndexFunc[S ~[]E, E any](s S, cmp func(E, E) int) int {
	if len(s) == 0 {
		return -1
	}
	maxIndex := 0
	for i := range s {
		if cmp(s[i], s[maxIndex]) > 0 {
			maxIndex = i
		}
	}
	return maxIndex
}

// MinMaxIndex returns the indices of the minimum and maximum elements in the slice.
func MinMaxIndex[S ~[]E, E constraints.Ordered](s S) (int, int) {
	return MinMaxIndexFunc(s, cmp.Compare)
}

// MinMaxIndexFunc returns the indices of the minimum and maximum elements in the slice.
func MinMaxIndexFunc[S ~[]E, E any](s S, cmp func(E, E) int) (int, int) {
	if len(s) == 0 {
		return -1, -1
	}
	minIndex := 0
	maxIndex := 0
	for i := range s {
		if cmp(s[i], s[minIndex]) < 0 {
			minIndex = i
		}
		if cmp(s[i], s[maxIndex]) > 0 {
			maxIndex = i
		}
	}
	return minIndex, maxIndex
}
