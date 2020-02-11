package demo

import (
	"fmt"
	"errors"
)

func defer_func() {
	//栈 先进后出
	defer func() {
		if p:=recover(); p!=nil {
			fmt.Println(p)
		}
		fmt.Println("1")
	}()
	defer func() { fmt.Println("2")}()
	defer func() { fmt.Println("3")}()

	panic(errors.New("触发异常"))
}

func multi_defer() {
	defer fmt.Println("first defer")
	for i := 0; i < 3; i++ {
		defer fmt.Printf("defer in for [%d]\n", i)
	}
	defer fmt.Println("last defer")
}

func TestDefer() {
	//defer_func()
	multi_defer()
}
