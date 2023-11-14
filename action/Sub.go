package action

import (
	"github.com/hacash/HVM/eval"
	"github.com/hacash/HVM/trait"
	"math/big"
)

type Sub struct {
	Left  trait.VMAction
	Right trait.VMAction
}

func (s *Sub) Type() uint16 {
	return 65522
}

func (s *Sub) IsBurning90PersentTxFees() bool {
	return false
}

func (s *Sub) ChildActions() []trait.VMAction {
	return []trait.VMAction{s.Left, s.Right}
}

func (s *Sub) Childs() []trait.ASTNode {
	return []trait.ASTNode{s.Left, s.Right}
}

func (s *Sub) Evaluate(ctx trait.Context) trait.EvalResult {
	return eval.DoCompute(ctx, s.Left, s.Right,
		func(l, r *big.Int) (*big.Int, trait.EvalResult) {
			return l.Sub(l, r), nil
		})
}
