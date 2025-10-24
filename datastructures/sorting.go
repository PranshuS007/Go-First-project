package datastructures

// QuickSort sorts a slice of integers using the QuickSort algorithm.
// It sorts the slice in-place and returns the sorted slice.
// Time Complexity: O(n log n) average case, O(n²) worst case
// Space Complexity: O(log n) due to recursion stack
func QuickSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}
	quickSortHelper(arr, 0, len(arr)-1)
	return arr
}

// quickSortHelper is a helper function that performs the recursive QuickSort.
func quickSortHelper(arr []int, low int, high int) {
	if low < high {
		// Partition the array and get the pivot index
		pivotIndex := partition(arr, low, high)

		// Recursively sort elements before and after partition
		quickSortHelper(arr, low, pivotIndex-1)
		quickSortHelper(arr, pivotIndex+1, high)
	}
}

// partition rearranges the array so that elements smaller than the pivot
// are on the left and elements greater than the pivot are on the right.
// Returns the final position of the pivot element.
func partition(arr []int, low int, high int) int {
	// Choose the rightmost element as pivot
	pivot := arr[high]
	i := low - 1

	for j := low; j < high; j++ {
		if arr[j] <= pivot {
			i++
			// Swap arr[i] and arr[j]
			arr[i], arr[j] = arr[j], arr[i]
		}
	}

	// Swap arr[i+1] and arr[high] (pivot)
	arr[i+1], arr[high] = arr[high], arr[i+1]
	return i + 1
}

// QuickSortDescending sorts a slice of integers in descending order using QuickSort.
// It sorts the slice in-place and returns the sorted slice.
func QuickSortDescending(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}
	quickSortDescendingHelper(arr, 0, len(arr)-1)
	return arr
}

// quickSortDescendingHelper is a helper function for descending QuickSort.
func quickSortDescendingHelper(arr []int, low int, high int) {
	if low < high {
		pivotIndex := partitionDescending(arr, low, high)
		quickSortDescendingHelper(arr, low, pivotIndex-1)
		quickSortDescendingHelper(arr, pivotIndex+1, high)
	}
}

// partitionDescending partitions the array for descending order sorting.
func partitionDescending(arr []int, low int, high int) int {
	pivot := arr[high]
	i := low - 1

	for j := low; j < high; j++ {
		if arr[j] >= pivot {
			i++
			arr[i], arr[j] = arr[j], arr[i]
		}
	}

	arr[i+1], arr[high] = arr[high], arr[i+1]
	return i + 1
}

// MergeSort sorts a slice of integers using the MergeSort algorithm.
// It creates a new sorted slice and returns it without modifying the original.
// Time Complexity: O(n log n)
// Space Complexity: O(n)
func MergeSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	// Create a copy to avoid modifying the original array
	result := make([]int, len(arr))
	copy(result, arr)

	mergeSortHelper(result, 0, len(result)-1)
	return result
}

// mergeSortHelper is a helper function that performs the recursive MergeSort.
func mergeSortHelper(arr []int, left int, right int) {
	if left < right {
		// Find the middle point
		mid := left + (right-left)/2

		// Sort first and second halves
		mergeSortHelper(arr, left, mid)
		mergeSortHelper(arr, mid+1, right)

		// Merge the sorted halves
		merge(arr, left, mid, right)
	}
}

// merge combines two sorted subarrays into a single sorted subarray.
func merge(arr []int, left int, mid int, right int) {
	// Calculate sizes of two subarrays
	n1 := mid - left + 1
	n2 := right - mid

	// Create temporary arrays
	leftArr := make([]int, n1)
	rightArr := make([]int, n2)

	// Copy data to temporary arrays
	for i := 0; i < n1; i++ {
		leftArr[i] = arr[left+i]
	}
	for j := 0; j < n2; j++ {
		rightArr[j] = arr[mid+1+j]
	}

	// Merge the temporary arrays back into arr[left..right]
	i := 0 // Initial index of first subarray
	j := 0 // Initial index of second subarray
	k := left // Initial index of merged subarray

	for i < n1 && j < n2 {
		if leftArr[i] <= rightArr[j] {
			arr[k] = leftArr[i]
			i++
		} else {
			arr[k] = rightArr[j]
			j++
		}
		k++
	}

	// Copy remaining elements of leftArr[], if any
	for i < n1 {
		arr[k] = leftArr[i]
		i++
		k++
	}

	// Copy remaining elements of rightArr[], if any
	for j < n2 {
		arr[k] = rightArr[j]
		j++
		k++
	}
}

// MergeSortDescending sorts a slice of integers in descending order using MergeSort.
// It creates a new sorted slice and returns it without modifying the original.
func MergeSortDescending(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	result := make([]int, len(arr))
	copy(result, arr)

	mergeSortDescendingHelper(result, 0, len(result)-1)
	return result
}

// mergeSortDescendingHelper is a helper function for descending MergeSort.
func mergeSortDescendingHelper(arr []int, left int, right int) {
	if left < right {
		mid := left + (right-left)/2
		mergeSortDescendingHelper(arr, left, mid)
		mergeSortDescendingHelper(arr, mid+1, right)
		mergeDescending(arr, left, mid, right)
	}
}

// mergeDescending combines two sorted subarrays in descending order.
func mergeDescending(arr []int, left int, mid int, right int) {
	n1 := mid - left + 1
	n2 := right - mid

	leftArr := make([]int, n1)
	rightArr := make([]int, n2)

	for i := 0; i < n1; i++ {
		leftArr[i] = arr[left+i]
	}
	for j := 0; j < n2; j++ {
		rightArr[j] = arr[mid+1+j]
	}

	i := 0
	j := 0
	k := left

	for i < n1 && j < n2 {
		if leftArr[i] >= rightArr[j] {
			arr[k] = leftArr[i]
			i++
		} else {
			arr[k] = rightArr[j]
			j++
		}
		k++
	}

	for i < n1 {
		arr[k] = leftArr[i]
		i++
		k++
	}

	for j < n2 {
		arr[k] = rightArr[j]
		j++
		k++
	}
}

// MergeSortInPlace sorts a slice of integers in-place using MergeSort.
// This modifies the original slice.
func MergeSortInPlace(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}
	mergeSortHelper(arr, 0, len(arr)-1)
	return arr
}
