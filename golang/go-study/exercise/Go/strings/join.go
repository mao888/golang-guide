package main

import (
	"fmt"
	"strings"
)

func main() {
	status := []string{"3", "2", "1"}

	var statusStr []string
	for _, statu := range status {
		statusStr = append(statusStr, fmt.Sprintf("'%s'", statu))
	}

	fmt.Println(strings.Join(status, ","))

	fmt.Println(strings.Join(statusStr, ","))
}
