# slicesx

[![Go Reference](https://pkg.go.dev/badge/github.com/pkierski/slicesx.svg)](https://pkg.go.dev/github.com/pkierski/slicesx)
[![rcard](https://goreportcard.com/badge/github.com/pkierski/slicesx)](https://goreportcard.com/report/github.com/pkierski/slicesx)

This package is an enhancement for the standard [slices](https://pkg.go.dev/slices) Go package.

This is a set of generic functions I missed in `slices`:

## Installation

```bash
go get github.com/pkierski/slicesx
```

## Functions

### Looking for index of min/max element

The standard `slices.Min` and `slices.Max` functions return the minimal/maximal element, but sometimes you need the index of the element in the slice, not the value itself.

#### `MinIndex` / `MinIndexFunc`

Returns the index of the minimum element in the slice. Returns `-1` if the slice is empty.

```go
slice := []int{1, 2, 3, 4, 5, -6, 7, 8, 9, 10}
minIndex := MinIndex(slice) // returns 5
```

#### `MaxIndex` / `MaxIndexFunc`

Returns the index of the maximum element in the slice. Returns `-1` if the slice is empty.

```go
slice := []int{1, 2, 3, 4, 5, -6, 7, 8, 9, 10}
maxIndex := MaxIndex(slice) // returns 9
```

#### `MinMaxIndex` / `MinMaxIndexFunc`

Returns the indices of both the minimum and maximum elements in a single pass. Returns `(-1, -1)` if the slice is empty.

```go
slice := []int{1, 2, 3, 4, 5, -6, 7, 8, 9, 10}
minIndex, maxIndex := MinMaxIndex(slice) // returns 5, 9
```

### Partitioning

Split a slice into two parts based on a predicate function.

#### `PartitionInPlace`

Partitions elements by predicate in place. **Relative order of elements is not preserved.** Returns two slices which are partitions of the original slice.

- Complexity: O(n)
- Space complexity: O(1)

```go
slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
evenSlice, oddSlice := PartitionInPlace(slice, func(i int) bool {
    return i%2 == 0
})
// slice is modified in place
```

#### `PartitionInPlaceStable`

Partitions elements by predicate in place while **preserving the relative order** of elements within each partition.

- Complexity: O(n²)
- Space complexity: O(1)

```go
slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
evenSlice, oddSlice := PartitionInPlaceStable(slice, func(i int) bool {
    return i%2 == 0
})
// slice is modified in place, but order is preserved
```

#### `PartitionStable`

Partitions elements by predicate while **preserving the relative order** of elements within each partition. The original slice is **not modified** - returns two new slices.

- Complexity: O(n)
- Space complexity: O(n)

```go
slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
evenSlice, oddSlice := PartitionStable(slice, func(i int) bool {
    return i%2 == 0
})
// slice remains unchanged
```

### Selection

#### `Select`

Selects elements from a slice that satisfy a predicate. Returns a new slice containing only the elements that match the predicate.

```go
slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
evenSlice := Select(slice, func(i int) bool {
    return i%2 == 0
})
// returns [2 4 6 8 10]
```