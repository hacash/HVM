package action

import (
	"github.com/hacash/HVM/eval"
	"github.com/hacash/HVM/trait"
	"math/big"
)

type Mul struct {
	Left  trait.VMAction
	Right trait.VMAction
}

func (s *Mul) Type() uint16 {
	return 65523
}

func (s *Mul) IsBurning90PersentTxFees() bool {
	return false
}

func (s *Mul) ChildActions() []trait.VMAction {
	return []trait.VMAction{s.Left, s.Right}
}

func (s *Mul) Childs() []trait.ASTNode {
	return []trait.ASTNode{s.Left, s.Right}
}

func (s *Mul) Evaluate(ctx trait.Context) trait.EvalResult {
	return eval.DoCompute(ctx, s.Left, s.Right,
		func(l, r *big.Int) (*big.Int, trait.EvalResult) {
			return l.Mul(l, r), nil
		})
}
