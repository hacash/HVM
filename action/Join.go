package action

import (
	"bytes"
	"fmt"
	"github.com/hacash/HVM/eval"
	"github.com/hacash/HVM/stack"
	"github.com/hacash/HVM/trait"
)

type Join struct {
	Left  trait.VMAction
	Right trait.VMAction
}

func (s *Join) VMKind() uint8 {
	return 15
}

func (s *Join) IsBurning90PersentTxFees() bool {
	return false
}

func (s *Join) ChildActions() []trait.VMAction {
	return []trait.VMAction{s.Left, s.Right}
}

func (s *Join) Childs() []trait.ASTNode {
	return []trait.ASTNode{s.Left, s.Right}
}

func (s *Join) Evaluate(ctx trait.Context) trait.EvalResult {
	var fte, lv, rv = eval.EvalLeftRight(ctx, s.Left, s.Right)
	if fte != nil {
		return fte
	}
	var lt, rt = lv.RetValue(), rv.RetValue()
	if len(lt)+len(rt) > stack.STACK_ITEM_MAX_SIZE {
		return eval.ResultFatalErr(fmt.Errorf("Join stackoverflow"))
	}
	// ok ret
	var buf = bytes.NewBuffer(lt)
	buf.Write(rt) // join
	return eval.ResultValue(buf.Bytes())
}
