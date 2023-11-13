package action

import (
	"fmt"
	"github.com/hacash/HVM/eval"
	"github.com/hacash/HVM/trait"
	"github.com/hacash/core/fields"
)

type Number struct {
	wide fields.VarUint1
	vbts []byte
}

func (s *Number) Type() uint16 {
	return 65525
}

func (s *Number) IsBurning90PersentTxFees() bool {
	return false
}

func (s *Number) ChildActions() []trait.VMAction {
	return []trait.VMAction{}
}

func (s *Number) Childs() []trait.ASTNode {
	return []trait.ASTNode{}
}

func (s *Number) Evaluate(ctx trait.Context) trait.EvalResult {
	if int(s.wide) != len(s.vbts) {
		return eval.ResultFatalErr(fmt.Errorf("Number bytes wide %d not macth count %d",
			len(s.vbts), s.wide))
	}
	if len(s.vbts) > eval.NUMBER_MAX_SIZE {
		return eval.ResultFatalErr(fmt.Errorf("Number bytes wide %d cannot over than max %d",
			len(s.vbts), eval.NUMBER_MAX_SIZE))
	}
	// ok
	return eval.ResultValue(s.vbts)
}
