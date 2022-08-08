/**
    @author:Hasee
    @data:2022/3/22
    @note:
**/
package main
import (
	"fmt"
	"strconv"
)

func main() {
	n, k := 0, 0
	fmt.Scan(&n, &k)
	var a = make([][]uint8, n)
	for i := range a {
		a[i] = make([]uint8, n)
	}
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			fmt.Scan(&a[i][j])
		}
	}

	var b = make([][]uint8, n*k)
	for i := range b {
		b[i] = make([]uint8, n*k)
	}

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			for c := 0; c < k; c++ {
				for d := 0; d < k; d++ {
					b[i*k+c][j*k+d] = a[i][j]
				}
			}
		}
	}
	tmp := ""
	for i := 0; i < n*k; i++ {
		for j := 0; j < n*k; j++ {
			tmp += strconv.Itoa(int(b[i][j])) + " "
		}
		tmp += "\n"
	}
	fmt.Print(tmp)
}

