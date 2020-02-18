package demo

import (
	"fmt"
	"encoding/json"
)
//Go 语言中，大括号不能放在单独的一行。

//以下代码，编译错误 syntax error: unexpected semicolon or newline before {
//语法错误
//func TestMain()
//{
//
//}

//未使用变量：如果有未使用的变量代码将编译失败。但也有例外，函数中声明的变量必须要使用，但可以有未使用的全局变量。函数的参数未使用也是可以的。
//如果你给未使用的变量分配了一个新值，代码也还是会编译失败。你需要在某个地方使用这个变量，才能让编译器愉快的编译。
//以下代码编译失败
//var gvar int
//func Test021801() {
//	var one int
//
//	func(unusedvar string) {
//		fmt.Println("Hello world")
//	}("nihao")
//}


//参考答案及解析：运行时错误。如果类型实现 String() 方法，当格式化输出时会自动使用 String() 方法。
// 上面这段代码是在该类型的 String() 方法内使用格式化输出，导致递归调用，最后抛错。
type ConfigOne struct {
	Daemon string
}

func (c *ConfigOne) String() string {
	return fmt.Sprintf("print: %v", c)
}

func Test021802() {
	c := &ConfigOne{}
	c.String()
}

//rune 是 int32 的别名一样，byte 是 uint8 的别名，别名类型无序转换，可直接转换。

//不像变量，常量未使用是能编译通过的。

//:= 操作符不能用于结构体字段赋值。

type Peopler struct {
	Name string `json:"name"`  //如果为name，则输出 {}。知识点：结构体访问控制，因为 name 首字母是小写，导致其他包不能访问，所以输出为空结构体。
}

func Test021803() {
	js := `{
		"name": "jack"	
	}`
	var p Peopler
	err := json.Unmarshal([]byte(js), &p)
	if err != nil {
		fmt.Println("err:", err)
		return
	}

	fmt.Println(p)
}



