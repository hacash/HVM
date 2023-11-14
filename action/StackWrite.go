package action

import (
	"bytes"
	"fmt"
	"github.com/hacash/HVM/eval"
	"github.com/hacash/HVM/stack"
	"github.com/hacash/HVM/trait"
)

type StackWrite struct {
	Ptrv trait.VMAction
	Data trait.VMAction
}

func (s *StackWrite) Type() uint16 {
	return 65519
}

func (s *StackWrite) IsBurning90PersentTxFees() bool {
	return false
}

func (s *StackWrite) ChildActions() []trait.VMAction {
	return []trait.VMAction{s.Ptrv, s.Data}
}

func (s *StackWrite) Childs() []trait.ASTNode {
	return []trait.ASTNode{s.Ptrv, s.Data}
}

func (s *StackWrite) Evaluate(ctx trait.Context) trait.EvalResult {
	var ptr, fte = eval.CalcStackPtr(ctx, s.Ptrv)
	if fte != nil {
		return fte
	}
	// data
	var dv = s.Data.Evaluate(ctx)
	if dv.CheckInterrupt() {
		return dv
	}
	var dts = dv.RetValue()
	if len(dts) > stack.STACK_ITEM_MAX_SIZE {
		return eval.ResultFatalErr(fmt.Errorf("StackWrite value over size %d over max %d",
			len(dts), stack.STACK_ITEM_MAX_SIZE)) // stackoverflow
	}
	// write
	var stk = ctx.GetStack()
	var e = stk.Write(int(ptr), bytes.NewBuffer(dts))
	if e != nil {
		return eval.ResultFatalErr(e) // stackoverflow
	}
	// ok return true
	return eval.ResultTrue()
}
