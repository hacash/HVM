package vm

import "github.com/hacash/HVM/trait"

func (vm *HacashVM) Evaluate(node trait.ASTNode) trait.EvalResult {
	return node.Evaluate(vm)
}
