package stack

import (
	"bytes"
	"fmt"
	"github.com/hacash/HVM/util"
)

func (s VMStack) is_out_of_bounds(ptr int) error {
	if ptr < s.seg_start || ptr >= s.seg_end {
		return fmt.Errorf("Stack read out of bounds")
	}
	return nil
}

func (s VMStack) Read(ptr int) (*bytes.Buffer, error) {
	var e = s.is_out_of_bounds(ptr)
	if e != nil {
		return nil, e
	}
	var res = s.data_space[ptr]
	if res == nil {
		return nil, fmt.Errorf("Stack read uninitialized")
	}
	// ok
	return res, nil
}

func (s *VMStack) Write(ptr int, buf *bytes.Buffer) error {
	if s.read_only {
		return fmt.Errorf("Stack write readonly")
	}
	var e = s.is_out_of_bounds(ptr)
	if e != nil {
		return e
	}
	if buf == nil {
		return fmt.Errorf("Stack write buffer cannot is nil")
	}
	var sl = buf.Len()
	if sl < 1 {
		return fmt.Errorf("Stack write buffer cannot empty")
	}
	if sl > STACK_ITEM_MAX_SIZE {
		return fmt.Errorf("Stack write buffer size %d overflow %d",
			sl, STACK_ITEM_MAX_SIZE)
	}
	// OK copy and write
	s.data_space[ptr] = util.CopyVMBuffer(buf)
	return nil
}
