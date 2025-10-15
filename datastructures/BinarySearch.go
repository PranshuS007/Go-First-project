package datastructures

// BinarySearch performs a binary search on a sorted slice of integers.
// It returns the index of the target element if found, otherwise returns -1.
// Time Complexity: O(log n)
// Space Complexity: O(1)
func BinarySearch(arr []int, target int) int {
	left := 0
	right := len(arr) - 1

	for left <= right {
		mid := left + (right-left)/2

		if arr[mid] == target {
			return mid
		}

		if arr[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return -1
}

// BinarySearchRecursive performs a recursive binary search on a sorted slice.
// It returns the index of the target element if found, otherwise returns -1.
func BinarySearchRecursive(arr []int, target int, left int, right int) int {
	if left > right {
		return -1
	}

	mid := left + (right-left)/2

	if arr[mid] == target {
		return mid
	}

	if arr[mid] < target {
		return BinarySearchRecursive(arr, target, mid+1, right)
	}

	return BinarySearchRecursive(arr, target, left, mid-1)
}

// BinarySearchFirst finds the first occurrence of target in a sorted slice.
// Returns -1 if target is not found.
func BinarySearchFirst(arr []int, target int) int {
	left := 0
	right := len(arr) - 1
	result := -1

	for left <= right {
		mid := left + (right-left)/2

		if arr[mid] == target {
			result = mid
			right = mid - 1 // Continue searching in left half
		} else if arr[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return result
}

// BinarySearchLast finds the last occurrence of target in a sorted slice.
// Returns -1 if target is not found.
func BinarySearchLast(arr []int, target int) int {
	left := 0
	right := len(arr) - 1
	result := -1

	for left <= right {
		mid := left + (right-left)/2

		if arr[mid] == target {
			result = mid
			left = mid + 1 // Continue searching in right half
		} else if arr[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return result
}
