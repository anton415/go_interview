package main

import (
	"fmt"
)

type Counter struct {
	Value int
}

func (c Counter) Print() {
	fmt.Println(c.Value)
}

// напечатает 123
func deferredMethod() {
	c := Counter{Value: 123}
	defer c.Print()
	c.Value = 456
}

// напечатает 456
func deferredClosure() {
	c := Counter{Value: 123}
	defer func() {
		c.Print()
	}()
	c.Value = 456
}

// напечатает 10 20
// 10 в момент вызова defer, 20 в момент выполнения defer
func mixed() {
	value := 10
	defer func(saved int) {
		fmt.Println(saved, value)
	}(value)
	value = 20
}

// вернет 3 2 1
func deferredOrder() (order []int) {
	for i := 1; i <= 3; i++ {
		defer func() {
			order = append(order, i)
		}()
	}
	return
}
