package action

import (
	"github.com/hacash/HVM/eval"
	"github.com/hacash/HVM/trait"
	"math/big"
)

type Add struct {
	Left  trait.VMAction
	Right trait.VMAction
}

func (s *Add) Type() uint16 {
	return 65521
}

func (s *Add) IsBurning90PersentTxFees() bool {
	return false
}

func (s *Add) ChildActions() []trait.VMAction {
	return []trait.VMAction{s.Left, s.Right}
}

func (s *Add) Childs() []trait.ASTNode {
	return []trait.ASTNode{s.Left, s.Right}
}

func (s *Add) Evaluate(ctx trait.Context) trait.EvalResult {
	return eval.DoCompute(ctx, s.Left, s.Right,
		func(l, r *big.Int) (*big.Int, trait.EvalResult) {
			return l.Add(l, r), nil
		})
}
