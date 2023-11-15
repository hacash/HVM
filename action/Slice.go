package action

import (
	"fmt"
	"github.com/hacash/HVM/eval"
	"github.com/hacash/HVM/stack"
	"github.com/hacash/HVM/trait"
	"math/big"
)

type Slice struct {
	Src   trait.VMAction
	Start trait.VMAction
	End   trait.VMAction
}

func (s *Slice) VMKind() uint8 {
	return 14
}

func (s *Slice) IsBurning90PersentTxFees() bool {
	return false
}

func (s *Slice) ChildActions() []trait.VMAction {
	return []trait.VMAction{s.Src, s.Start, s.End}
}

func (s *Slice) Childs() []trait.ASTNode {
	return []trait.ASTNode{s.Src, s.Start, s.End}
}

func (s *Slice) Evaluate(ctx trait.Context) trait.EvalResult {
	// data
	var dv = s.Src.Evaluate(ctx)
	if dv.CheckInterrupt() {
		return dv
	}
	var bts = dv.RetValue()
	var datalen = len(bts)
	// start
	var sp = s.Start.Evaluate(ctx)
	if sp.CheckInterrupt() {
		return sp
	}
	var sdt = sp.RetValue()
	if len(sdt) > 2 {
		return eval.ResultFatalErr(fmt.Errorf("Slice start ptr size overflow"))
	}
	var si = int(big.NewInt(0).SetBytes(sdt).Uint64())
	// end
	var ep = s.End.Evaluate(ctx)
	if ep.CheckInterrupt() {
		return ep
	}
	var edt = ep.RetValue()
	if len(edt) > 2 {
		return eval.ResultFatalErr(fmt.Errorf("Slice end ptr size overflow"))
	}
	var ei = int(big.NewInt(0).SetBytes(edt).Uint64())
	// check
	if si >= ei {
		return eval.ResultFatalErr(fmt.Errorf("Slice start ptr cannot big than or equal to end ptr"))
	}
	if ei > datalen {
		return eval.ResultFatalErr(fmt.Errorf("Slice end ptr cannot big than data size"))
	}
	if ei-si > stack.STACK_ITEM_MAX_SIZE {
		return eval.ResultFatalErr(fmt.Errorf("Slice result size %d overflow %d",
			ei-si, stack.STACK_ITEM_MAX_SIZE))
	}
	// ok
	var resv = bts[si:ei]
	return eval.ResultValue(resv)
}
