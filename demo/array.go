package demo

import "fmt"

func TestArray() {
	arr := [5]int{1,2,3,4,5}
	for k,v := range arr {
		if k + 1 < len(arr) {
			arr[k+1] = arr[k+1] + 1
		}
		fmt.Printf("%v-%v ", v, arr[k])
	}
	fmt.Println()
	//初始化切片
	slice := make([]int, 5, 8)
	fmt.Println(slice)

	//空切片
	empty := []int{}
	fmt.Println(empty)

	//扩容
	base := make([]int, 10)
	fmt.Printf("base len: %v, capacity: %v \n", len(base), cap(base))
	base = append(base, 1)
	fmt.Printf("base len: %v, capacity: %v \n", len(base), cap(base))

	base2 := make([]int, 80)
	fmt.Printf("base len: %v, capacity: %v \n", len(base2), cap(base2))
	base2 = append(base2, 1)
	fmt.Printf("base len: %v, capacity: %v \n", len(base2), cap(base2))
}
