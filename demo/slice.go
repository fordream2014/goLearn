package demo

import "fmt"

func Assign1(s []int) {
	s = []int{6, 6, 6}
}

func Reverse0(s [5]int) {
	for i, j := 0, len(s)-1; i < j; i++ {
		j = len(s) - (i + 1)
		s[i], s[j] = s[j], s[i]
	}
}

func Reverse1(s []int) {
	for i, j := 0, len(s)-1; i < j; i++ {
		j = len(s) - (i + 1)
		s[i], s[j] = s[j], s[i]
	}
}

func Reverse2(s []int) {
	s = append(s, 999)
	//fmt.Println(s)
	for i, j := 0, len(s)-1; i < j; i++ {
		j = len(s) - (i + 1)
		s[i], s[j] = s[j], s[i]
		//fmt.Println(s)
	}
	//fmt.Println(s)
}

func Reverse3(s []int) {
	s = append(s, 999, 1000, 1001)
	for i, j := 0, len(s)-1; i < j; i++ {
		j = len(s) - (i + 1)
		s[i], s[j] = s[j], s[i]
	}
	fmt.Println(s)
}
//slice这种数据结构便于使用和管理数据集合，可以理解为是一种“动态数组”，slice也是围绕动态数组的概念来构建的。既然是动态数组，那么slice是如何扩容的呢？
//
//请记住以下两条规则：
//
//如果切片的容量小于1024个元素，那么扩容的时候slice的cap就翻番，乘以2；一旦元素个数超过1024个元素，增长因子就变成1.25，即每次增加原来容量的四分之一。
//如果扩容之后，还没有触及原数组的容量，那么，切片中的指针指向的位置，就还是原数组，如果扩容之后，超过了原数组的容量，那么，Go就会开辟一块新的内存，把原来的值拷贝过来，
// 这种情况丝毫不会影响到原数组。
func TestSlice() {
	var a []int
	fmt.Println(len(a), cap(a)) //0
	a = append(a, 1)
	fmt.Println(len(a), cap(a)) //1
	a = append(a, 1)
	fmt.Println(len(a), cap(a)) //2
	a = append(a, 1)
	fmt.Println(len(a), cap(a)) //3
	a = append(a, 1)
	fmt.Println(len(a), cap(a)) //4
	a = append(a, 1)
	fmt.Println(len(a), cap(a)) //3

	//s := []int{1, 2, 3, 4, 5, 6}
	//Assign1(s)
	//fmt.Println(s) // (1)
	//
	//array := [5]int{1, 2, 3, 4, 5}
	//Reverse0(array)
	//fmt.Println(array) // (2)
	//
	//s = []int{1, 2, 3}
	//Reverse2(s)
	//fmt.Println(s) // (3)

	//var a []int
	//for i := 1; i <= 3; i++ {
	//	a = append(a, i)
	//}
	//Reverse2(a)
	//fmt.Println(a) // (4)
	//

	//var b []int
	//for i := 1; i <= 3; i++ {
	//	b = append(b, i)
	//}
	//Reverse3(b)
	//fmt.Println(b) // (5)
	//

	//c := [3]int{1, 2, 3}
	//d := c
	//c[0] = 999
	//fmt.Println(d) // (6)
}

