package main

import "fmt"

func main() {
	m := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
		"five":  5,
		"six":   6,
		"seven": 7,
		"eight": 8,
	}
	//for k, v := range m {
	//	delete(m, "two")
	//	m["four"] = 4
	//	m["nine"] = 9
	//	m["ten"] = 10
	//	fmt.Printf("%v: %v\n", k, v)
	//}
	//one: 1
	//four: 4
	//three: 3
	//five: 5
	//six: 6
	//seven: 7
	//eight: 8

	for i := 0; i < len(m); i++ {
		delete(m, "two")
		m["four"] = 4
		m["nine"] = 9
		m["ten"] = 10
		fmt.Printf("%v: %v\n", i, m)
	}

	m2 := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
	}

	for i := 0; i < len(m2); i++ {
		fmt.Println("len(m):", len(m2))
		delete(m2, "two")
		m2["four"] = 4
		m2["five"] = 5 // 添加更多键，强制 rehash
		fmt.Println("len(m):", len(m2))
	}

}
