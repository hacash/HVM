package action

import (
	"github.com/hacash/HVM/eval"
	"github.com/hacash/HVM/trait"
)

type Or struct {
	Left  trait.VMAction
	Right trait.VMAction
}

func (s *Or) Type() uint16 {
	return 65529
}

func (s *Or) IsBurning90PersentTxFees() bool {
	return false
}

func (s *Or) ChildActions() []trait.VMAction {
	return []trait.VMAction{s.Left, s.Right}
}

func (s *Or) Childs() []trait.ASTNode {
	return []trait.ASTNode{s.Left, s.Right}
}

func (s *Or) Evaluate(ctx trait.Context) trait.EvalResult {
	var fe, lv, rv = eval.EvalLeftRight(ctx, s.Left, s.Right)
	if fe != nil {
		return fe
	}
	// ret
	if lv.IsTrue() || rv.IsTrue() {
		return eval.ResultTrue() // true
	} else {
		return eval.ResultFalse() // false
	}
}
