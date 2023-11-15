package vm

import "github.com/hacash/HVM/trait"

func (vm *HacashVM) GetStack() trait.Stack {
	return vm.stack
}

func (vm *HacashVM) GetHeap() trait.Heap {
	return vm.heap
}

func (vm *HacashVM) GetExtendCallExecutor() trait.ExtendCallExecutor {
	return vm.extca
}
