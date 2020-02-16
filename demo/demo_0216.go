package demo

import (
	"fmt"
	"time"
)

//var a[]int  不会分配内存，优先使用这种使用方式
//a := []int{}

//函数参数为interface{}时可以接收任何类型的参数，包括用户自定义类型等，即使是接收指针类型也用interface{},
//而不是*interface{}

//golang的字符串类型不能赋值nil，也不能跟nil比较

//map的输出是无序的
func TestMap() {
	var arr map[int]string = make(map[int]string)
	arr[0] = "zero"
	arr[1] = "one"

	for k,v := range arr {
		fmt.Println(k, v)
	}
	//输出
	//0 zero
	//1 one
	//或者
	// 1 one
	//0 zero

}

//基于类型创建的方法必须定义在同一个包内
//下面的代码compilation error
//func (i int) PrintInt() {
//	fmt.Println(i)
//}

//解决方法：
type Myint int
func (i Myint) PrintInt() {
	fmt.Println(i)
}
func TestOtherName() {
	var i Myint = 1
	i.PrintInt()
}

const (
	a = iota
	b = iota
)

const (
	name = "name"
	c = iota
	d = iota
)

//iota是golang中的常量计数器，只能在常量表达式中使用
//iota在const关键字出现时将被重置为0，const中每增一行常量声明将使iota计数一次。
func TestConst() {
	fmt.Println(a, b, c, d)
}

//当使用fmt.Printf()、fmt.Print()、fmt.Println()会自动使用String()方法，实现字符串的打印。

//对于类似X = Y的赋值操作，必须知道X的地址，才能够将Y的值赋给X，
// 但是go中的map的value本身是不可寻址的。
//因此，类似如下的操作，是编译错误的
type Math struct {
	x, y int
}

var m = map[string]Math{
	"foo": Math{2, 3},
}
//
//func main() {
//	m["foo"].x = 4
//	fmt.Println(m["foo"].x)
//}
//解决方法如下：

var n = map[string]*Math{
	"foo": &Math{2,3},
}
func TestMapUpdate() {
	//1 使用临时变量
	tmp := m["foo"]
	tmp.x = 4
	m["foo"] = tmp
	fmt.Println(m["foo"].x)

	//2 修改数据结构
	n["foo"].x = 4
	fmt.Println(n["foo"].x)

	//Map是对底层数据的引用，编写代码的过程中，会涉及到Map拷贝、函数间传递Map等。
	//跟Slice类似，Map指向的底层数据是不会发生copy的。
	m := map[string]int{
		"January":1,
		"February":2,
		"March":3,
	}
	month := m
	delete(month,"February")
	fmt.Println(m)
	fmt.Println(month)

	//如果想拷贝一个Map，如下所示
	month = map[string]int{}
	m = map[string]int{
		"January":1,
		"February":2,
		"March":3,
	}
	for key,value := range m{
		month[key] = value
	}
	delete(month,"February")
	fmt.Println(m)
	fmt.Println(month)
}

//golang中不同类型不能比较，数组长度是数组类型的一部分，不同类型的数组不能比较
//切片，map，函数不能比较

var pp *int

func foo() (*int, error) {
	var i int = 5
	return &i, nil
}

func bar() {
	//use p
	fmt.Println(*pp)
}

//runtime error
//问题出在:=，对于使用:=定义的变量，如果新变量与同名已定义的变量不在同一个作用域中
//那么Go会新定义这个变量。
func TestLocal() {
	pp, err := foo()
	if err != nil {
		fmt.Println(err)
		return
	}
	bar() //使用的是全局变量
	fmt.Println(*pp) //使用的是局部变量pp
}

func TestSliceRand() {
	v := []int{1, 2, 3}
	for i := range v {
		v = append(v, i)
	}
	fmt.Println(v)
}

func TestArrayRand() {

	var m = [...]int{1, 2, 3}

	for i, v := range m {
		go func() {
			fmt.Println(i, v)
		}()
	}

	time.Sleep(time.Second * 3)
}

//len函数可以获取切片、数组、map、channel的大小





