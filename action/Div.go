package action

import (
	"fmt"
	"github.com/hacash/HVM/eval"
	"github.com/hacash/HVM/trait"
	"math/big"
)

type Div struct {
	Left  trait.VMAction
	Right trait.VMAction
}

func (s *Div) Type() uint16 {
	return 65526
}

func (s *Div) IsBurning90PersentTxFees() bool {
	return false
}

func (s *Div) ChildActions() []trait.VMAction {
	return []trait.VMAction{s.Left, s.Right}
}

func (s *Div) Childs() []trait.ASTNode {
	return []trait.ASTNode{s.Left, s.Right}
}

func (s *Div) Evaluate(ctx trait.Context) trait.EvalResult {
	return eval.DoCompute(ctx, s.Left, s.Right,
		func(l, r *big.Int) (*big.Int, trait.EvalResult) {
			if r.Uint64() == 0 {
				return nil, eval.ResultFatalErr(fmt.Errorf("Div cannot be divided by zero"))
			}
			return l.Div(l, r), nil
		})
}
