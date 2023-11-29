package stack

import (
	"bytes"
)

// total max size = 256*128 = 1024*32 = 32KB
const STACK_TOTAL_LENGTH int = 1024
const STACK_ITEM_MAX_SIZE int = 256
const STACK_FUNC_ITEM_LENGTH int = 128

type VMStack struct {
	data_space *[STACK_TOTAL_LENGTH]*bytes.Buffer
	seg_start  int
	seg_end    int
	read_only  bool
}

func NewVMStack() *VMStack {
	var new_space = new([STACK_TOTAL_LENGTH]*bytes.Buffer)
	var stc = &VMStack{
		new_space,
		0,
		0,
		false,
	}
	return stc
}
