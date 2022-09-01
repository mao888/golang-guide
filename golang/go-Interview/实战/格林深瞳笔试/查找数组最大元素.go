package main

import "math"

func main() {

}

func findMax(nums []int) int {
	a := len(nums)

	get := func(i int) int {
		if i == -1 || i == a {
			return math.MinInt64
		}
		return nums[i]
	}

	left, right := 0, a-1
	for {
		mid := (left + right) / 2
		if get(mid-1) < get(mid) && get(mid) > get(mid+1) {
			return nums[mid]
		}
		if get(mid) < get(mid+1) {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
}
