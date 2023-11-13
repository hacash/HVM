package stack

import (
	"fmt"
	"testing"
)

func Test1(t *testing.T) {

	var stk = NewVMStack()
	child, e := stk.Alloc(10, false)
	fmt.Println(e)
	sunzi, e := child.Alloc(1023, false)
	fmt.Println(e)

	fmt.Printf("%d, %d\n", sunzi.seg_start, sunzi.seg_end)

}
