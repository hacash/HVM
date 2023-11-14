package action

import (
	"github.com/hacash/HVM/eval"
	"github.com/hacash/HVM/trait"
)

type And struct {
	Left  trait.VMAction
	Right trait.VMAction
}

func (s *And) Type() uint16 {
	return 65528
}

func (s *And) IsBurning90PersentTxFees() bool {
	return false
}

func (s *And) ChildActions() []trait.VMAction {
	return []trait.VMAction{s.Left, s.Right}
}

func (s *And) Childs() []trait.ASTNode {
	return []trait.ASTNode{s.Left, s.Right}
}

func (s *And) Evaluate(ctx trait.Context) trait.EvalResult {
	var fe, lv, rv = eval.EvalLeftRight(ctx, s.Left, s.Right)
	if fe != nil {
		return fe
	}
	// ret
	if lv.IsTrue() && rv.IsTrue() {
		return eval.ResultTrue() // true
	} else {
		return eval.ResultFalse() // false
	}
}
