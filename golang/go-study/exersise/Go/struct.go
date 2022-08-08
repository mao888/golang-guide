/**
    @author:Hasee
    @data:2022/6/16
    @note:
**/
package main

import "fmt"

type name struct {
	name int
	kk   func(int2 int) int
}

func (a *name) ff() {
	a.name = 1
	a.kk = func(int2 int) int {
		return 1
	}

	fmt.Printf("%d", a.kk(1))

}

func main() {
	a := name{}
	a.ff()
}
