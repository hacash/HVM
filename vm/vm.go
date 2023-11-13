package vm

import (
	"github.com/hacash/HVM/config"
	"github.com/hacash/HVM/heap"
	"github.com/hacash/HVM/stack"
	"github.com/hacash/HVM/trait"
)

type HacashVM struct {
	cnf   *config.VMConfig
	stack *stack.VMStack
	heap  *heap.VMHeap
}

func NewHacashVM(cnf *config.VMConfig) *HacashVM {

	return &HacashVM{
		cnf:   cnf,
		stack: stack.NewVMStack(),
		heap:  heap.NewVMHeap(),
	}
}

func (vm *HacashVM) GetStack() trait.Stack {
	return vm.stack
}

func (vm *HacashVM) GetHeap() trait.Heap {
	return vm.heap
}