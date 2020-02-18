package demo

import (
	"fmt"
	"strconv"
)

//cap()函数适用于数组、切片、channel
//字典，删除不存在的key，不会报错；获取不存在的key，返回零值

//%d 输出十进制数字， +表示输出数值的符号
func TestPrintValue() {
	a := -1
	b := +5
	fmt.Printf("%+d, %d, %+d, %d", a, a, b, b) // -1, -1, +5, 5
}

type Speaker interface {
	speak()
}

type User struct {
	Name string
	Email string
}

func (u *User) speak() {
	fmt.Println("I am user speak")
}

type Admin struct {
	User
	Level string
}

func gotospeak(speaker Speaker) {
	speaker.speak()
}

//struct嵌入类型

//通过嵌入，内部类型的属性、方法，可以为外部类型所有，就好像是外部类型自己的一样。此外，外部类型还可以定义自己的属性和方法，
// 甚至可以定义与内部相同的方法，这样内部类型的方法就会被“屏蔽”。

//假设外部结构体类型是 S，内部类型是 T，则关于内部类型的方法提升如下规则：
//T 嵌入 S，外部类型 S 可以通过值类型或指针类型调用内部类型 T 的值方法；
//T 嵌入 S，外部类型 S 只能通过指针类型调用内部类型 T 的指针方法；
//*T 嵌入 S，外部类型 S 可以通过值类型和指针类型调用内部类型 T 的值方法和指针方法；
//上面的三条规则可以总结成一句话：不管是 T 嵌入 S，还是 *T 嵌入 S，外部类型 S 唯独通过值类型
// 不能调用内部类型 T 的指针方法外，其他情况下内部类型 T 的方法都可以获得提升，即可被外部类型 S 访问 。

func TestSpeak() {

	admin := &Admin{
		User: User{
			Name: "xugaung",
			Email: "xuguang5@sina.com",
		},
	}

	gotospeak(admin)
}

//-----------------
type People struct{}

func (p *People) ShowA() {
	fmt.Println("showA")
	p.ShowB()
}
func (p *People) ShowB() {
	fmt.Println("showB")
}

type Teacher struct {
	People
}

func (t *Teacher) ShowB() {
	fmt.Println("teacher showB")
}

func TestSuper() {
	t := Teacher{}
	t.ShowA()
}

//--------
func helloworld(a int) {
	fmt.Println(a)
}

//helloworld() 函数的参数在执行 defer 语句的时候会保存一份副本
func TestDeferRand() {
	i := 5
	defer helloworld(i)
	i += 10
}

//字符串为只读
func TestString() {
	str := "hello"
	//str[0] = 'x' //compilation error
	fmt.Println(str)
}

func TestNil() {
	var s1 []int
	var s2 = []int{}
	if s2 == nil {
		fmt.Println("yes nil")
	}else{
		fmt.Println("no nil")
	}

	if s1 == nil {
		fmt.Println("yes nil")
	}else{
		fmt.Println("no nil")
	}

	i := 65
	fmt.Println(strconv.Itoa(i)) //65
	fmt.Println(string(i)) //A
}

//-------------
//increaseA()int 函数的返回值没有被提前声名，其值来自于其他变量的赋值，而defer中修改的也是其他变量，而非返回值本身，因此函数退出时返回值并没有被改变。
//increaseB()(r int) 函数的返回值被提前声名，也就意味着defer中是可以调用到真实返回值的，因此defer在return赋值返回值 i 之后，再一次地修改了 i 的值，
// 最终函数退出后的返回值才会是defer修改过的值。
func increaseA() int {
	var i int
	defer func() {
		i++
		fmt.Println("defer " + strconv.Itoa(i))
	}()
	fmt.Println("after defer " + strconv.Itoa(i))
	return i
}

func increaseB() (r int) {
	defer func() {
		r++
	}()
	return r
}

func TestIncrease() {
	fmt.Println(increaseA())
	fmt.Println(increaseB())
}

//-----
func f1() (r int) { //1
	defer func() {
		r++
	}()
	return 0
}

func f2() (r int) { //5
	t := 5
	defer func() {
		t = t + 5
	}()
	return t
}

func f3() (r int) { //1
	defer func(r int) {
		r = r + 5
	}(r)
	return 1
}
func TestDeferRand1() {
	f1r := f1()
	f2r := f2()
	f3r := f3()
	fmt.Println(f1r, f2r, f3r)
}

//------------
type Person struct {
	age int
}

func TestDefer2() {
	person := &Person{28}

	// 1.
	defer fmt.Println(person.age)

	// 2.
	defer func(p *Person) {
		fmt.Println(p.age)
	}(person)

	// 3.
	defer func() {
		fmt.Println(person.age)
	}()

	person = &Person{29}
}








