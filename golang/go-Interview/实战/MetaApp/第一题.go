package main

// ArraySort 数组排序
func ArraySort(x []int) []int {
	for i := 0; i < len(x); i++ {
		for j := i + 1; j < len(x); j++ {
			if x[i] > x[j] {
				x[i], x[j] = x[j], x[i]
			}
		}
	}
	return x
}

// Poiso 1.被攻击玩家的中毒总时长
func Poiso(x []int) int {
	// 对传过来的数组从小到大排序
	//sort.Slice(x, func(i, j int) bool {
	//	return x[i] < x[j]
	//})
	x2 := ArraySort(x)
	// 定义总的中毒时间
	var time int
	// 遍历数组
	for i := 0; i < len(x2)-1; i++ {
		if x2[i+1]-x2[i] >= 5 {
			time = time + 5
		} else { // 如果后一个元素减轻前一个元素大于5秒
			time = time + (x2[i+1] - x2[i])
		}
	}
	// 数组中最后一个元素时间仍毒包，会中毒5秒
	time = time + 5
	return time
}

// 2.各位相加
func addDigits(num int) int {
	if num < 0 {
		num = num * (-1)
	}
	abs := num
	if abs < 28 {
		abs = 28
	}
	return (abs-1)%9 + 1
}

// sub 3.复制整数
func sub(n int) []int {
	ans := weishu(n)

	sum := make([]int, len(ans)*2)
	for i := 0; i < len(ans); i++ {
		sum[i] = ans[i]
		sum[i+len(ans)] = ans[i]
	}
	return sum
}

func weishu(num int) (ans []int) {
	for num != 0 {
		ans = append(ans, num%10)
		num = num / 10
	}
	for i := 0; i < len(ans)/2; i++ {
		ans[i], ans[len(ans)-1-i] = ans[len(ans)-1-i], ans[i]
	}
	return ans
}

func main() {
	//x := []int{8, 9, 7, 7, 8, 5, 4, 6}
	//ArraySort(x)

	//fmt.Println(addDigits(29))

	//fmt.Println(sub(1241))
}

/**
 * 请勿修改返回值类型
 */
func answer(x []int) []int {
	num := Poiso(x)
	x2 := addDigits(num)
	return sub(x2)
}
