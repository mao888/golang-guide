package main

import (
	"fmt"
)

// getMaxXorSubset 计算所有非空子集的最大异或和
func getMaxXorSubset(arr []int) int {
	// 创建一个长度为64的数组basis，用于存储每个位上的基
	basis := make([]int, 64)

	// 遍历数组中的每一个数
	for _, num := range arr {
		// 检查num的每一位，从高到低
		for i := 63; i >= 0; i-- {
			// 如果num在第i位上没有1，继续检查下一位
			if (num>>i)&1 == 0 {
				continue
			}
			// 如果basis[i]是0，说明该位置还没有基向量，把num放入basis[i]
			if basis[i] == 0 {
				basis[i] = num
				break
			}
			// 如果basis[i]已经存在，则将num与basis[i]异或，继续检查下一位
			num ^= basis[i]
		}
	}

	// 初始化maxXor为0，用于存储最大异或和
	maxXor := 0
	// 遍历basis中的每一个基向量
	for i := 63; i >= 0; i-- {
		// 尝试将basis[i]加入到maxXor中，如果能得到更大的值则更新maxXor
		if (maxXor ^ basis[i]) > maxXor {
			maxXor ^= basis[i]
		}
	}
	return maxXor
}

func main() {
	// 示例数组
	arr := []int{1, 2, 3, 4}
	// 计算并打印所有非空子集的最大异或和
	fmt.Println("所有非空子集的最大异或和是:", getMaxXorSubset(arr)) // 输出: 7
}
