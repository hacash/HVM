package action

import (
	"github.com/hacash/HVM/eval"
	"github.com/hacash/HVM/trait"
)

type Continue struct {
}

func (s *Continue) VMKind() uint8 {
	return 29
}

func (s *Continue) IsBurning90PersentTxFees() bool {
	return true
}

func (s *Continue) ChildActions() []trait.VMAction {
	return []trait.VMAction{}
}

func (s *Continue) Childs() []trait.ASTNode {
	return []trait.ASTNode{}
}

func (s *Continue) Evaluate(ctx trait.Context) trait.EvalResult {
	// mark Continue
	return eval.ResultValueTy([]byte{}, eval.ResultEventTypeLoopContinue)
}
