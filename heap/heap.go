package heap

import "bytes"

const HEAP_TOTAL_MAX_SIZE = 1024 * 8 // 8KB = (128+128)*32
const HEAP_ITEM_MAX_SIZE = 256

type VMHeap struct {
	data_space map[string]*bytes.Buffer
	used_size  int // heap already used size
}

func NewVMHeap() *VMHeap {
	var heap = &VMHeap{
		data_space: make(map[string]*bytes.Buffer),
		used_size:  0,
	}
	return heap
}
