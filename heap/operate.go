package heap

import (
	"bytes"
	"fmt"
	"github.com/hacash/HVM/util"
)

func (h VMHeap) Read(key string) (*bytes.Buffer, error) {
	if res, ok := h.data_space[key]; ok && res != nil {
		return res, nil
	}
	// not find
	return nil, fmt.Errorf("Heap read uninitialized")
}

func (h *VMHeap) Write(key string, buf *bytes.Buffer) error {
	var kl = len(key)
	if kl < 1 {
		return fmt.Errorf("Heap write key cannot empty")
	}
	if kl > HEAP_ITEM_MAX_SIZE {
		return fmt.Errorf("Heap write key size %d overflow %d",
			kl, HEAP_ITEM_MAX_SIZE)
	}
	if buf == nil {
		return fmt.Errorf("Heap write value cannot is nil")
	}
	var vl = buf.Len()
	if vl < 1 {
		return fmt.Errorf("Heap write value cannot empty")
	}
	if vl > HEAP_ITEM_MAX_SIZE {
		return fmt.Errorf("Heap write value size %d overflow %d",
			vl, HEAP_ITEM_MAX_SIZE)
	}
	// count used size
	var used_size_change = 0
	if val, has := h.data_space[key]; has {
		used_size_change = vl - val.Len()
	} else {
		used_size_change = kl + vl
	}
	var new_used_size = h.used_size + used_size_change
	if new_used_size > HEAP_TOTAL_MAX_SIZE {
		return fmt.Errorf("Heap write memory overflow")
	}
	// change used size
	h.used_size = new_used_size
	// write
	h.data_space[key] = util.CopyVMBuffer(buf)
	// ok
	return nil
}
