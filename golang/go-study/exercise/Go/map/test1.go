package main

import "fmt"

func main() {
	type Map map[string][]int
	m := make(Map)
	s := []int{1, 2}
	s = append(s, 3)
	fmt.Printf("%v\n", s) // [1 2 3]
	m["q1mi"] = s
	fmt.Println(m) // map[q1mi:[1 2 3]]

	//m["q1mi"] = make([]int, len(s))
	//copy(m["q1mi"], s)

	s = append(s[:1], s[2:]...)
	fmt.Printf("%v\n", s)         // [1 3]
	fmt.Printf("%v\n", m["q1mi"]) // [1 3 3]

	s = append(s, 4, 5, 6)
	fmt.Printf("%v\n", s)         // [1 3 4 5 6]
	fmt.Printf("%v\n", m["q1mi"]) // [1 3 3]
}
