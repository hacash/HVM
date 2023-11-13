package action

import (
	"github.com/hacash/HVM/eval"
	"github.com/hacash/HVM/trait"
)

type Null struct {
}

func (s *Null) Type() uint16 {
	return 65535
}

func (s *Null) IsBurning90PersentTxFees() bool {
	return false
}

func (s *Null) ChildActions() []trait.VMAction {
	return []trait.VMAction{}
}

func (s *Null) Childs() []trait.ASTNode {
	return []trait.ASTNode{}
}

func (s *Null) Evaluate(ctx trait.Context) trait.EvalResult {
	return eval.ResultNone()
}
