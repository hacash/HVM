package action

import (
	"github.com/hacash/HVM/eval"
	"github.com/hacash/HVM/trait"
)

type Break struct {
	Value trait.VMAction
}

func (s *Break) Type() uint16 {
	return 65530
}

func (s *Break) IsBurning90PersentTxFees() bool {
	return false
}

func (s *Break) ChildActions() []trait.VMAction {
	return []trait.VMAction{s.Value}
}

func (s *Break) Childs() []trait.ASTNode {
	return []trait.ASTNode{s.Value}
}

func (s *Break) Evaluate(ctx trait.Context) trait.EvalResult {
	var v = s.Value.Evaluate(ctx)
	if v.CheckInterrupt() {
		return v
	}
	// mark break
	return eval.ResultValueTy(v.RetValue(), eval.ResultEventTypeLoopBreak)
}
