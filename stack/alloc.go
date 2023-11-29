package stack

import (
	"fmt"
	"github.com/hacash/HVM/trait"
)

func allocVMStack(base *VMStack, seg_size int, is_read_only bool) (*VMStack, error) {
	if is_read_only == false && base.read_only == true {
		return nil, fmt.Errorf("Stack alloc must readonly")
	}
	if seg_size < 1 {
		return nil, fmt.Errorf("Stack seg_size cannot less than 1")
	}
	if seg_size > STACK_FUNC_ITEM_LENGTH {
		return nil, fmt.Errorf("Sub Stack seg_size %d cannot more than %d",
			seg_size, STACK_ITEM_MAX_SIZE)
	}
	var new_start = base.seg_end
	var new_end = new_start + seg_size
	if new_end >= STACK_TOTAL_LENGTH {
		return nil, fmt.Errorf("Stack alloc overflow, %d + %d = %d >= %d",
			new_start, seg_size, new_end, STACK_TOTAL_LENGTH)
	}
	var stc = &VMStack{
		base.data_space,
		new_start,
		new_end,
		is_read_only,
	}
	return stc, nil

}

func (s *VMStack) AllocSub(seg_size int, is_read_only bool) (*VMStack, error) {
	return allocVMStack(s, seg_size, is_read_only)
}

func (s *VMStack) Alloc(seg_size int, is_read_only bool) (trait.Stack, error) {
	return s.AllocSub(seg_size, is_read_only)
}
