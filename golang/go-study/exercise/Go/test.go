package main

import (
	"fmt"
	"sync"
	"time"
)

type Person struct {
	Name string
}

func (p *Person) SetName() {
	p.Name = "1"
}

func main() {
	//fmt.Println(mathClass.Add(2, 3))
	//fmt.Println(mathClass.Sub(3, 2))
	//p := Person{Name: "33"}
	//p.SetName()
	//fmt.Println(p.Name)
	//for i := 0; i < 10; i++ {
	//	wg.Add(1)
	//	go func(i int) {
	//		defer wg.Done()
	//		fmt.Println(i)
	//		time.Sleep(time.Second)
	//	}(i)
	//}
	ch := make(chan int)
	var w sync.WaitGroup
	w.Add(1)
	go func() {
		fmt.Println("11111")
		<-ch
		fmt.Println("3333333")
		w.Done()
	}()
	time.Sleep(time.Second * 2)
	ch <- 1
	w.Wait()
}
