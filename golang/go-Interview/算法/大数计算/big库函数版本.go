/**
    @author:HuChao
    @data:2022/9/24
    @note:
**/
package main

import (
	"fmt"
	"math/big"
)

func main() {
	var a, b string
	fmt.Scan(&a)
	fmt.Scan(&b)
	t1, _ := new(big.Int).SetString(a, 16)
	t2, _ := new(big.Int).SetString(b, 16)
	dummy := t1.Add(t1, t2)
	fmt.Print(dummy)
}
