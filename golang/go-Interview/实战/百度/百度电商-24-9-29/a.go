package main

import (
	"fmt"
	"math"
)

//找零36，现有硬币面额1、2、5、10、
//20（各面值硬币量足够）：最少需要
//多少枚硬币，该硬币组合是？

// 找零问题：给定面额为1、2、5、10、20的硬币，找零36最少需要多少枚硬币
func coinChange(coins []int, amount int) (int, []int) {
	// 建立一个长度为 amount+1 的数组，用于存储每个金额所需的最少硬币数
	dp := make([]int, amount+1)
	// 初始化 dp 数组，将未被计算过的金额设为 math.MaxInt32
	for i := 1; i <= amount; i++ {
		dp[i] = math.MaxInt32
	}

	// 用于记录每个金额所使用的硬币
	coinUsed := make([][]int, amount+1)

	// 动态规划过程
	for i := 1; i <= amount; i++ {
		for _, coin := range coins {
			if i-coin >= 0 && dp[i-coin]+1 < dp[i] {
				dp[i] = dp[i-coin] + 1
				// 复制前一个状态的硬币使用记录，并添加当前硬币
				coinUsed[i] = append(coinUsed[i-coin], coin)
			}
		}
	}

	// 如果 dp[amount] 仍然是 math.MaxInt32，说明无法找零
	if dp[amount] == math.MaxInt32 {
		return -1, nil
	}

	return dp[amount], coinUsed[amount]
}

func main() {
	coins := []int{1, 2, 5, 10, 20}
	amount := 36

	minCoins, coinCombination := coinChange(coins, amount)
	if minCoins == -1 {
		fmt.Println("无法找零")
	} else {
		fmt.Printf("最少需要 %d 枚硬币\n", minCoins)
		fmt.Printf("硬币组合是: %v\n", coinCombination)
	}
}
