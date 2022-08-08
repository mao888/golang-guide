/**
    @author:Hasee
    @data:2022/4/9
    @note:
**/
package main

import (
	"fmt"
	"math/rand"
	"sync"
)

func main() {
	//fmt.Println("random: ", rand.Intn(100))
	res := Example(6, 10)
	for i := 0; i < len(res); i++ {
		fmt.Println(i+1, res[i])
	}
}

/**
len: 长度
max: 个数
随机生成长度为len的max个纯数字随机数
*/
func Example(length int, max int) []string {
	// Seeding with the same value results in the same random sequence each run.
	// For different numbers, seed with a different value, such as
	// time.Now().UnixNano(), which yields a constantly-changing number.
	//rand.Seed(42)

	digitNumber := []string{
		"1",
		"2",
		"3",
		"4",
		"5",
		"6",
		"7",
		"8",
		"9",
		"0",
	}
	// 用make创建map

	set := New()
	for set.Len() < max {
		ranNumber := ""
		for j := 1; j < length; j++ {
			ranNumber += digitNumber[rand.Intn(len(digitNumber))]
		}
		if !set.Has(ranNumber) {
			set.Add(ranNumber)
		}
	}

	return set.List()

}

/**
构造set类型
*/
type Set struct {
	m map[string]bool
	sync.RWMutex
}

func New() *Set {
	return &Set{
		m: map[string]bool{},
	}
}

func (s *Set) Add(item string) {
	s.Lock()
	defer s.Unlock()
	s.m[item] = true
}

func (s *Set) Remove(item string) {
	s.Lock()
	s.Unlock()
	delete(s.m, item)
}

func (s *Set) Has(item string) bool {
	s.RLock()
	defer s.RUnlock()
	_, ok := s.m[item]
	return ok
}

func (s *Set) Len() int {
	return len(s.List())
}

func (s *Set) Clear() {
	s.Lock()
	defer s.Unlock()
	s.m = map[string]bool{}
}

func (s *Set) IsEmpty() bool {
	if s.Len() == 0 {
		return true
	}
	return false
}

func (s *Set) List() []string {
	s.RLock()
	defer s.RUnlock()
	list := []string{}
	for item := range s.m {
		list = append(list, item)
	}
	return list
}
