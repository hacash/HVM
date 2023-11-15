package action

import (
	"fmt"
	"github.com/hacash/HVM/eval"
	"github.com/hacash/HVM/trait"
	"github.com/hacash/core/fields"
)

type Sequential struct {
	CondAtLeast fields.VarUint2
	CondAtMost  fields.VarUint2
	Count       fields.VarUint2
	Lists       []trait.VMAction
}

func (s *Sequential) VMKind() uint8 {
	return 33
}

func (s *Sequential) IsBurning90PersentTxFees() bool {
	return false
}

func (s *Sequential) ChildActions() []trait.VMAction {
	return s.Lists
}

func (s *Sequential) Childs() []trait.ASTNode {
	var size = len(s.Lists)
	var ary = make([]trait.ASTNode, size)
	for i := 0; i < size; i++ {
		ary[i] = s.Lists[i]
	}
	return ary
}

func (s *Sequential) Evaluate(ctx trait.Context) trait.EvalResult {
	if s.CondAtLeast < 1 {
		return eval.ResultFatalErr(fmt.Errorf("Sequential CondAtLeast cannot less than 1"))
	}
	if s.CondAtMost < s.CondAtLeast {
		return eval.ResultFatalErr(fmt.Errorf("Sequential CondAtMost %d cannot less than CondAtLeast %d",
			s.CondAtMost, s.CondAtLeast))
	}
	if s.Count < s.CondAtMost {
		return eval.ResultFatalErr(fmt.Errorf("Sequential Count %d cannot less than CondAtMost %d",
			s.Count, s.CondAtMost))
	}
	if int(s.Count) != len(s.Lists) {
		return eval.ResultFatalErr(fmt.Errorf("Sequential Count %d not match real list length %d",
			s.Count, len(s.Lists)))
	}
	// execute each item in list
	var retv trait.EvalResult = nil
	var success int = 0
	for i := 0; i < int(s.Count); i++ {
		var res = s.Lists[i].Evaluate(ctx)
		if res.CheckInterrupt() {
			return res // fatal error
		}
		if res.IsTrue() {
			success++
		}
		retv = res
		if success >= int(s.CondAtMost) {
			break // condition have already finished
		}
	}
	// check condition
	if success < int(s.CondAtLeast) {
		return eval.ResultFatalErr(fmt.Errorf("Sequential Success %d cannot less than CondAtLeast %d",
			success, s.CondAtLeast))
	}
	// finish return
	return retv
}
