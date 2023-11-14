package action

import (
	"github.com/hacash/HVM/eval"
	"github.com/hacash/HVM/trait"
)

type Not struct {
	Numv trait.VMAction
}

func (s *Not) Type() uint16 {
	return 65530
}

func (s *Not) IsBurning90PersentTxFees() bool {
	return false
}

func (s *Not) ChildActions() []trait.VMAction {
	return []trait.VMAction{s.Numv}
}

func (s *Not) Childs() []trait.ASTNode {
	return []trait.ASTNode{s.Numv}
}

func (s *Not) Evaluate(ctx trait.Context) trait.EvalResult {
	var rv = s.Numv.Evaluate(ctx)
	if rv.CheckInterrupt() {
		return rv
	}
	if rv.IsTrue() {
		return eval.ResultFalse() // false
	} else {
		return eval.ResultTrue() // true
	}
}
