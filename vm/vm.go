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
	extca trait.ExtendCallExecutor
	cttld trait.ContractLoader
}

func NewHacashVM(cnf *config.VMConfig,
	extCaller trait.ExtendCallExecutor,
	codeLoader trait.ContractLoader) *HacashVM {

	return &HacashVM{
		cnf:   cnf,
		stack: stack.NewVMStack(),
		heap:  heap.NewVMHeap(),
		extca: extCaller,
		cttld: codeLoader,
	}
}
