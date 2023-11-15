package action

import (
	"github.com/hacash/HVM/eval"
	"github.com/hacash/HVM/trait"
	"math/big"
)

type Abs struct {
	Numv trait.VMAction
}

func (s *Abs) VMKind() uint8 {
	return 26
}

func (s *Abs) IsBurning90PersentTxFees() bool {
	return false
}

func (s *Abs) ChildActions() []trait.VMAction {
	return []trait.VMAction{s.Numv}
}

func (s *Abs) Childs() []trait.ASTNode {
	return []trait.ASTNode{s.Numv}
}

func (s *Abs) Evaluate(ctx trait.Context) trait.EvalResult {
	var rv = s.Numv.Evaluate(ctx)
	if rv.CheckInterrupt() {
		return rv
	}
	if !rv.IsTrue() {
		return eval.ResultFalse()
	}
	// do abs
	var bv = big.NewInt(0).SetBytes(rv.RetValue())
	bv = bv.Abs(bv)
	return eval.ResultValue(bv.Bytes())
}
