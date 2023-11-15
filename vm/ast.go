package vm

import (
	"github.com/hacash/HVM/action"
)

func (vm *HacashVM) BuildAST(buf []byte, seek uint32) (*action.List, uint32, error) {
	var list = &action.List{}
	seek2, e := list.Parse(vm.extca, buf, seek)
	if e != nil {
		return nil, 0, e
	}
	// ok
	return list, seek2, nil
}
