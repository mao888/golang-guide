package main

import (
	"cmp"
	"fmt"
)

const insertionSortThreshold = 12

func main() {
	ints := []int{5, 2, 9, 1, 5, 6}
	QuickSort(ints)
	fmt.Println(ints)

	words := []string{"pear", "apple", "orange", "apple", "banana"}
	QuickSort(words)
	fmt.Println(words)
}

// QuickSort sorts an ordered slice in place.
func QuickSort[T cmp.Ordered](arr []T) {
	if len(arr) < 2 {
		return
	}

	quickSort[T](arr, 0, len(arr)-1)
}

// quickSort uses median-of-three pivot selection, 3-way partitioning,
// and insertion sort for small ranges.
func quickSort[T cmp.Ordered](arr []T, left, right int) {
	for left < right {
		if right-left+1 <= insertionSortThreshold {
			insertionSort[T](arr, left, right)
			return
		}

		pivotIndex := medianOfThree[T](arr, left, right)
		arr[left], arr[pivotIndex] = arr[pivotIndex], arr[left]

		lt, gt := partition[T](arr, left, right)

		if lt-left < right-gt {
			quickSort[T](arr, left, lt-1)
			left = gt + 1
		} else {
			quickSort[T](arr, gt+1, right)
			right = lt - 1
		}
	}
}

func partition[T cmp.Ordered](arr []T, left, right int) (int, int) {
	pivot := arr[left]
	lt, i, gt := left, left+1, right

	for i <= gt {
		switch {
		case arr[i] < pivot:
			arr[lt], arr[i] = arr[i], arr[lt]
			lt++
			i++
		case arr[i] > pivot:
			arr[i], arr[gt] = arr[gt], arr[i]
			gt--
		default:
			i++
		}
	}

	return lt, gt
}

func medianOfThree[T cmp.Ordered](arr []T, left, right int) int {
	mid := left + (right-left)/2

	if arr[left] > arr[mid] {
		arr[left], arr[mid] = arr[mid], arr[left]
	}
	if arr[left] > arr[right] {
		arr[left], arr[right] = arr[right], arr[left]
	}
	if arr[mid] > arr[right] {
		arr[mid], arr[right] = arr[right], arr[mid]
	}

	return mid
}

func insertionSort[T cmp.Ordered](arr []T, left, right int) {
	for i := left + 1; i <= right; i++ {
		value := arr[i]
		j := i - 1

		for j >= left && arr[j] > value {
			arr[j+1] = arr[j]
			j--
		}

		arr[j+1] = value
	}
}
