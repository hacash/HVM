package action

import (
	"fmt"
	"github.com/hacash/HVM/eval"
	"github.com/hacash/HVM/stack"
	"github.com/hacash/HVM/trait"
	"github.com/hacash/core/fields"
)

type Bytes struct {
	wide fields.VarUint1
	vbts []byte
}

func (s *Bytes) VMKind() uint8 {
	return 20
}

func (s *Bytes) IsBurning90PersentTxFees() bool {
	return false
}

func (s *Bytes) ChildActions() []trait.VMAction {
	return []trait.VMAction{}
}

func (s *Bytes) Childs() []trait.ASTNode {
	return []trait.ASTNode{}
}

func (s *Bytes) Evaluate(ctx trait.Context) trait.EvalResult {
	if int(s.wide) != len(s.vbts) {
		return eval.ResultFatalErr(fmt.Errorf("Bytes wide %d not macth count %d",
			len(s.vbts), s.wide))
	}
	if len(s.vbts) > stack.STACK_ITEM_MAX_SIZE {
		return eval.ResultFatalErr(fmt.Errorf("Bytes bytes wide %d cannot over than max %d",
			len(s.vbts), stack.STACK_ITEM_MAX_SIZE))
	}
	// ok
	return eval.ResultValue(s.vbts)
}
