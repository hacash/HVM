package action

import (
	"fmt"
	"github.com/hacash/HVM/eval"
	"github.com/hacash/HVM/stack"
	"github.com/hacash/HVM/trait"
)

type StackLoad struct {
	Ptrv trait.VMAction
}

func (s *StackLoad) VMKind() uint8 {
	return 19
}

func (s *StackLoad) IsBurning90PersentTxFees() bool {
	return false
}

func (s *StackLoad) ChildActions() []trait.VMAction {
	return []trait.VMAction{s.Ptrv}
}

func (s *StackLoad) Childs() []trait.ASTNode {
	return []trait.ASTNode{s.Ptrv}
}

func (s *StackLoad) Evaluate(ctx trait.Context) trait.EvalResult {
	var ptr, fte = eval.CalcStackPtr(ctx, s.Ptrv)
	if fte != nil {
		return fte
	}
	// load
	var stk = ctx.GetStack()
	var v, e = stk.Read(ptr)
	if e != nil {
		return eval.ResultFatalErr(e) // stackoverflow
	}
	resv := v.Bytes()
	if len(resv) > stack.STACK_ITEM_MAX_SIZE {
		return eval.ResultFatalErr(fmt.Errorf("StackLoad value over size %d over max %d",
			len(resv), stack.STACK_ITEM_MAX_SIZE)) // stackoverflow
	}
	// ok
	return eval.ResultValue(resv)
}
