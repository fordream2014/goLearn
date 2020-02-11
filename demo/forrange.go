package demo

import "fmt"

//day 3
func TestMapAndSlice() {
	slice := []int{0,1,2,3}
	m := make(map[int]*int)

	for k,v := range slice {
		temp := &v
		//v 声明之后使用的是同一个地址
		//fmt.Println("%v", temp)
		m[k] = temp
	}

	for k,v := range m {
		fmt.Println(k, "->", *v)
	}
	fmt.Println("修正之后.....")
	//修正
	for k,v := range slice {
		temp := v
		m[k] = &temp
	}

	for k,v := range m {
		fmt.Println(k, "->", *v)
	}
}
