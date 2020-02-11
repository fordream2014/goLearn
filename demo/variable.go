package demo

import (
	"unsafe"
	"fmt"
)

func TestVariable() {
	var a int = 10
	var p *int =&a
	var c *int64
	c= (*int64)(unsafe.Pointer(p))
	fmt.Println(c)
}
