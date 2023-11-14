package action

import (
	"github.com/hacash/HVM/eval"
	"github.com/hacash/HVM/trait"
	"math/big"
)

type Mod struct {
	Numv trait.VMAction
	Base trait.VMAction
}

func (s *Mod) Type() uint16 {
	return 65525
}

func (s *Mod) IsBurning90PersentTxFees() bool {
	return false
}

func (s *Mod) ChildActions() []trait.VMAction {
	return []trait.VMAction{s.Numv, s.Base}
}

func (s *Mod) Childs() []trait.ASTNode {
	return []trait.ASTNode{s.Numv, s.Base}
}

func (s *Mod) Evaluate(ctx trait.Context) trait.EvalResult {
	return eval.DoCompute(ctx, s.Numv, s.Base,
		func(l, r *big.Int) (*big.Int, trait.EvalResult) {
			return l.Mod(l, r), nil
		})
}
