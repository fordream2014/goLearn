package demo

import (
	"fmt"
	"runtime"
)

func TestSlice001() {
	var a = []int{1, 2, 3, 4, 5}
	var r [5]int

	//注意使用数组和切片的区别
	for i, v := range a {
		if i == 0 {
			a[1] = 12
			a[2] = 13
		}
		r[i] = v
	}
	fmt.Println("r = ", r)
	fmt.Println("a = ", a)

	fmt.Println()

	var b = []int{1, 2, 3, 4, 5}
	var m [5]int

	//原因是for range 时会使用 b 的副本参与循环，len仍然是5
	for i, v := range b {
		if i == 0 {
			b = append(b, 6, 7)
		}
		m[i] = v
	}
	fmt.Println("m = ", m)
	fmt.Println("b = ", b)
}

func TestMap001() {
	var m = map[string]int{
		"A": 21,
		"B": 22,
		"C": 23,
	}
	counter := 0
	for k, v := range m {
		if counter == 0 {
			delete(m, "A")
		}
		counter++
		fmt.Println(k, v)
	}
	fmt.Println("counter is ", counter)
	//是因为map无序导致的
	//两个结果：
	//C 23
	//B 22
	//counter is  2

	//A 21
	//B 22
	//C 23
	//counter is  3
}

//多重赋值，步骤：
//1 计算等号左边的索引表达式和取址表达式，接着计算等号右边的表达式；
//2 赋值；
func Test021701() {
	i := 1
	s := []string{"A", "B", "C"}
	i, s[i-1] = 2, "Z"
	fmt.Printf("s: %v \n", s)
}

//break,continue可以在多层嵌套中指定目标层级
func Test021702() {

	outer:
		for i:=0; i<5; i++ {
			for j:=0; j<5; j++ {

				if j > 2 {
					continue outer
				}

				if i > 2 {
					break outer
				}

				fmt.Printf("%v : %v \n", i, j)
			}
		}
}

//自增、自减
// i++ 和 i-- 在Go语言中是语句，不是表达式，因此不能赋值给另外的变量。此外没有 ++i 和 --i。

func Test021703() {
	ch1 := make(chan string)
	ch2 := make(chan string)
	go func() {
		ch1 <- "hello"
	}()
	go func() {
		ch2 <- "world"
	}()
	select {
	case sh1 := <-ch1:
		fmt.Println(sh1)
	case sh2 := <-ch2:
		fmt.Println(sh2)
	}
}

func service1(ch chan string) {
	ch <- "from service1"
}

//以下代码报错
//原因：
//1 信道的默认值是 nil，不能对 nil 信道进行读写操作
//2 case 分支中如果信道是 nil，该分支就会被忽略，那么上面就变成空 select{} 语句，阻塞主协程
func Test021704() {
	var ch chan string
	go service1(ch)
	select {
		case str := <-ch:
			fmt.Println(str)
	}
}

func Test021705() {
	runtime.GOMAXPROCS(1)
	int_chan := make(chan int, 1)
	string_chan := make(chan string, 1)
	int_chan <- 1
	string_chan <- "hello"
	select {
	case value := <-int_chan:
		fmt.Println(value)
	case value := <-string_chan:
		panic(value)
	}
}

//A. 给一个 nil channel 发送数据，造成永远阻塞
//B. 从一个 nil channel 接收数据，造成永远阻塞
//C. 给一个已经关闭的 channel 发送数据，引起 panic
//D. 从一个已经关闭的 channel 接收数据，如果缓冲区中为空，则返回一个零值
func Test021706() {
	var ch chan string
	ch <- "hello"
}

const i = 100
var j = 123

// 常量不同于变量的在运行期分配内存，常量通常会被编译器在预处理阶段直接展开，作为指令数据使用，所以常量无法寻址。
func Test021707() {
	fmt.Println(&j, j)
	//fmt.Println(&i, i) //编译报错
}

type Userr struct{}
type User1 Userr
type User2 = Userr

func (i User1) m1() {
	fmt.Println("m1")
}
func (i Userr) m2() {
	fmt.Println("m2")
}

//注意使用 = 定义类型别名。因为 User2 是别名，完全等价于 User，所以 User2 具有 User 所有的方法。但是 i1.m1() 是不能执行的，因为 User1 没有定义该方法。
func Test021708() {
	var i1 User1
	var i2 User2
	i1.m1()
	i2.m2()
}

type Param map[string]interface{}

type Show struct {
	*Param
}
func Test021709() {
	show := new(Show)
	temp := make(Param)
	show.Param = &temp
	//show.Param["name"] = 1


}





















