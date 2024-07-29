package main

import "fmt"

//type TreeNode struct {
//	Val   int
//	Left  *TreeNode
//	Right *TreeNode
//}

//type ListNode struct {
//	Val  int
//	Next *ListNode
//}

func main() {
	a := twoSum([]int{2, 7, 11, 15}, 9)
	fmt.Printf("%v", a)
}

// 给定一个整数数组 nums 和一个整数目标值 target，请你在该数组中找出 和为目标值 target  的那 两个 整数，并返回它们的数组下标。
func twoSum(nums []int, target int) []int {
	hashTable := map[int]int{}
	for i, num := range nums {
		if preIndex, ok := hashTable[target-num]; ok {
			return []int{preIndex, i}
		}
		hashTable[num] = i
	}
	return nil
}
