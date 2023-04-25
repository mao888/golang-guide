/**
    @author:Hasee
    @data:2022/7/14
    @note:
**/
package main

import (
	"fmt"
	"os"
)

func main() {
	dir, _ := os.UserConfigDir()
	fmt.Println(dir)
}
