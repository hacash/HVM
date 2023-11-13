package action

import (
	"fmt"
	"github.com/hacash/HVM/eval"
	"github.com/hacash/HVM/trait"
	"github.com/hacash/core/fields"
)

type List struct {
	Count fields.VarUint2
	Lists []trait.VMAction
}

func (s *List) Kind() uint16 {
	return 65534
}

func (s *List) IsBurning90PersentTxFees() bool {
	return false
}

func (s *List) ChildActions() []trait.VMAction {
	return s.Lists
}

func (s *List) Childs() []trait.ASTNode {
	var size = len(s.Lists)
	var ary = make([]trait.ASTNode, size)
	for i := 0; i < size; i++ {
		ary[i] = s.Lists[i]
	}
	return ary
}

func (s *List) Evaluate(ctx trait.Context) trait.EvalResult {
	if int(s.Count) != len(s.Lists) {
		return eval.ResultFatalErr(fmt.Errorf("Sequential Count %d not match real list length %d",
			s.Count, len(s.Lists)))
	}
	// execute each item in list
	var retv trait.EvalResult = nil
	for i := 0; i < int(s.Count); i++ {
		var res = s.Lists[i].Evaluate(ctx)
		if res.FatalErr() != nil {
			return res // fatal error
		}
		retv = res
	}
	// finish return
	return retv
}
