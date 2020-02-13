package demo

import "fmt"
/*
1 结构体只能比较是否相等，不能比较大小
2 相同类型的结构体才能比较，结构体是否相等不但与属性类型有关，还与属性顺序相关
3 如果 struct 的所有成员都可以比较，则该 struct 就可以通过 == 或 != 进行比较是否相等，比较时逐个项进行比较，如果每一项都相等，则两个结构体才相等，否则不相等；
4 那什么是可比较的呢，常见的有 bool、数值型、字符、指针、数组等，像切片、map、函数等是不能比较的
 */
func TestStruct() {
	sn1 := struct {
		age int
		name string
	}{age: 11, name: "qq"}
	sn2 := struct {
		age  int
		name string
	}{age: 11, name: "qq"}

	if sn1 == sn2 {
		fmt.Println("sn1 == sn2")
	}
	//sm1 := struct {
	//	age int
	//	m   map[string]string
	//}{age: 11, m: map[string]string{"a": "1"}}
	//sm2 := struct {
	//	age int
	//	m   map[string]string
	//}{age: 11, m: map[string]string{"a": "1"}}
	//if sm1 == sm2 {
	//	fmt.Println("sm1 == sm2")
	//}
}

//

type MyInt1 int		//定义了另一种类型
type MyInt2 = int //类型别名

func TestAlias() {
	var i int =0
	var i1 MyInt1 = MyInt1(i)
	var i2 MyInt2 = i
	fmt.Println(i1,i2)
}

// iota只能在常量的表达式中使用

const (
	x = iota //0
    _
    y  //2
    z = "zz"
    k //zz
    p = iota //5
)

//iota 在下一行增长，而不是立即取得它的引用。
const (
	Apple, Banana = iota + 1, iota + 2
	Cherimoya, Durian
	Elderberry, Fig
)

func TestIota() {
	fmt.Println(x, y, z, k, p)
	fmt.Println(Apple, Banana, Cherimoya, Durian, Elderberry, Fig)
}

//nil 只能赋值给指针、chan、func、interface、map 或 slice 类型的变量。

//【init使用】
//init() 函数是用于程序执行前做包的初始化的函数，比如初始化包里的变量等;
//一个包可以出线多个 init() 函数,一个源文件也可以包含多个 init() 函数；
//同一个包中多个 init() 函数的执行顺序没有明确定义，但是不同包的init函数是根据包导入的依赖关系决定的（看下图）;
//init() 函数在代码中不能被显示调用、不能被引用（赋值给函数变量），否则出现编译错误;
//一个包被引用多次，如 A import B,C import B,A import C，B 被引用多次，但 B 包只会初始化一次；
//引入包，不可出现死循坏。即 A import B,B import A，这种情况编译失败；

type Shape interface {
	Area() float32
}

type Rect struct {
	width  float32
	height float32
}

func (r Rect) Area() float32 {
	return r.width * r.height
}

func TestInterface() {
	var s Shape
	s = Rect{5.0, 4.0}
	r := Rect{5.0, 4.0}
	fmt.Printf("type of s is %T\n", s)
	fmt.Printf("value of s is %v\n", s)
	fmt.Println("area of rectange s", s.Area())
	fmt.Println("s == r is", s == r)  //动态值的比较
}

//【接口】https://mp.weixin.qq.com/s/eDdrHwg0g7kLutDs-ejNpw
//当且仅当动态值和动态类型都为 nil 时，接口类型值才为 nil。
//类型选择的语法形如：i.(type)，其中 i 是接口，type 是固定关键字，需要注意的是，只有接口类型才可以使用类型选择。

func hello(num ...int) {
	fmt.Printf("%T \n", num)
	num[0] = 18
	fmt.Println(num)
}

func TestFunchello() {
	i := []int{1,2,3,4}
	hello(i...)
	fmt.Println(i)
}

//Go 提供了将切片传入可变参数函数的语法糖：直接在切片后加上 … 后缀。这样，切片将直接传入函数，不会再创建新的切片。
func change(s ...int) {
	s = append(s,3)
	fmt.Println(s)
}

func TestVariableFunc() {
	slice := make([]int,5,5)
	slice[0] = 1
	slice[1] = 2
	change(slice...)
	fmt.Println(slice)
	change(slice[0:2]...)
	fmt.Println(slice)

	fmt.Println()
	s := make([]int, 5,5)
	s[0] = 1
	s[1] = 2
	fmt.Println(s)

	ss := append(s[0:2], 3)
	fmt.Println(ss)
	fmt.Println(s)

	fmt.Println()
	//以上操作类似于
	sss := s[0:2]
	fmt.Println(len(sss), cap(sss))
	sss = append(sss, 3)
	fmt.Println(sss)
	fmt.Println(s)

}

type person struct {
	name string
}

func TestMap1111() {
	var m map[person]int
	p := person{"mike"}
	fmt.Println(m[p])
}

func TestAdd() {
	//a := [2]int{5, 6}
	//b := [3]int{5, 6}
	//if a == b {
	//	fmt.Println("equal")
	//} else {
	//	fmt.Println("not equal")
	//}
	// --以上结果为 编译出错，数组大小不同，不是同一类型

	//c := 4
	//d := 3.2
	//同类型的才能比较大小
	//if c > d {
	//
	//}

	arr := []int{1,2,3,4,5,6}
	fmt.Println(len(arr), cap(arr))

	slice := arr[1:3]
	fmt.Println(slice, len(slice), cap(slice))

	slice2 := arr[1:3:4]
	fmt.Println(slice2, len(slice2), cap(slice2))
}






