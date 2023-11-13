package heap

import (
	"bytes"
	"fmt"
	"testing"
)

func Test1(t *testing.T) {

	var vmh = NewVMHeap()

	vmh.Write("abc", bytes.NewBuffer([]byte("1234567890")))
	vmh.Write("abc", bytes.NewBuffer([]byte("12345")))
	vmh.Write("foo", bytes.NewBuffer([]byte("qwert")))
	vmh.Write("aa", bytes.NewBuffer([]byte("bb")))
	vmh.Write("foo", bytes.NewBuffer([]byte("1")))

	fmt.Printf("vmh.used_size = %d\n", vmh.used_size)

	if vmh.used_size != 16 {
		panic("vmh.used_size != 16")
	}

}
