package demo

import (
	"unsafe"
	"fmt"
)

var (
	size = 1024
	max_size = size * 2
)
func TestVariable() {
	var a int = 10
	var p *int =&a
	var c *int64
	c= (*int64)(unsafe.Pointer(p))
	fmt.Println(c)

	fmt.Println(size, max_size)
}
