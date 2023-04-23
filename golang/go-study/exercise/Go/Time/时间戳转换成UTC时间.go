package main

import (
	"fmt"
	"github.com/mao888/go-utils/constants"
	"time"
)

func main() {
	chuo := 1668583367
	bendi := time.Unix(int64(chuo), 0).Format(constants.TimeYMDHMM)
	lastLoginAt, err := time.Parse(constants.TimeYMDHMM, bendi)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(lastLoginAt.String())
}
