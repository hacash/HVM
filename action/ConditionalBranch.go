package action

import "github.com/hacash/HVM/trait"

type ConditionalBranch struct {
	Condition   trait.VMAction
	TrueBranch  trait.VMAction
	FalseBranch trait.VMAction
}

func (s *ConditionalBranch) Kind() uint16 {
	return 65532
}

func (s *ConditionalBranch) IsBurning90PersentTxFees() bool {
	return false
}

func (s *ConditionalBranch) ChildActions() []trait.VMAction {
	return []trait.VMAction{
		s.Condition, s.TrueBranch, s.FalseBranch,
	}
}

func (s *ConditionalBranch) Childs() []trait.ASTNode {
	return []trait.ASTNode{
		s.Condition, s.TrueBranch, s.FalseBranch,
	}
}

func (s *ConditionalBranch) Evaluate(ctx trait.Context) trait.EvalResult {
	var condv = s.Condition.Evaluate(ctx)
	if condv.CheckInterrupt() {
		return condv
	}
	if condv.IsTrue() {
		return s.TrueBranch.Evaluate(ctx) // true
	} else {
		return s.FalseBranch.Evaluate(ctx) // false
	}
}
