/**
    @author: edy
    @since: 2022/8/2
    @desc: //TODO
**/
package main

import (
	"sync"
)

type singleton struct{}

var instance *singleton
var once sync.Once

func GetInstance() *singleton {
	once.Do(func() {
		instance = &singleton{}
	})
	return instance
}

func main() {
	GetInstance()
}
