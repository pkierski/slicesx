package slicesx

// PartitionInPlace partitions elements by predicate in place.
//
// Relative order of elements is not preserved.
// Returns two slices which are the partitions of the original slice.
//
// Complexity: O(n)
// Space complexity: O(1)
func PartitionInPlace[S ~[]E, E any](s S, predicate func(E) bool) (S, S) {
	left := 0
	right := len(s) - 1

	for left < right {
		for left < right && predicate(s[left]) {
			left++
		}
		for left < right && !predicate(s[right]) {
			right--
		}
		if left < right {
			s[left], s[right] = s[right], s[left]
		}
	}

	return s[:left], s[left:]
}

// PartitionInPlaceStable partitions elements by predicate in place while preserving
// the relative order of elements within each partition.
//
// Returns two slices which are the partitions of the original slice.
//
// Complexity: O(n²)
// Space complexity: O(1)
func PartitionInPlaceStable[S ~[]E, E any](s S, predicate func(E) bool) (S, S) {
	if len(s) == 0 {
		return s, s
	}

	// Find the first element that should be on the right (doesn't match predicate)
	firstRight := 0
	for firstRight < len(s) && predicate(s[firstRight]) {
		firstRight++
	}

	if firstRight == len(s) {
		// All elements match the predicate
		return s, s[len(s):]
	}

	// Process remaining elements
	for i := firstRight + 1; i < len(s); i++ {
		if predicate(s[i]) {
			// Rotate elements from firstRight to i to bring s[i] to firstRight
			// This preserves the relative order of all elements
			rotateRight(s[firstRight : i+1])
			firstRight++
		}
	}

	return s[:firstRight], s[firstRight:]
}

// rotateRight rotates the slice right by one position
// This brings the last element to the first position
func rotateRight[S ~[]E, E any](s S) {
	if len(s) <= 1 {
		return
	}
	last := s[len(s)-1]
	copy(s[1:], s[:len(s)-1])
	s[0] = last
}

// PartitionStable partitions elements by predicate while preserving
// the relative order of elements within each partition.
//
// Returns two new slices which are the partitions of the original slice.
// The original slice is not modified.
//
// Complexity: O(n)
// Space complexity: O(n)
func PartitionStable[S ~[]E, E any](s S, predicate func(E) bool) (S, S) {
	return Select(s, predicate), Select(s, func(e E) bool { return !predicate(e) })
}
