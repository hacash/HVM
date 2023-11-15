package action

import (
	"github.com/hacash/HVM/eval"
	"github.com/hacash/HVM/trait"
)

type Loop struct {
	Cond trait.VMAction
	Eval trait.VMAction
}

func (s *Loop) VMKind() uint8 {
	return 31
}

func (s *Loop) IsBurning90PersentTxFees() bool {
	return true
}

func (s *Loop) ChildActions() []trait.VMAction {
	return []trait.VMAction{s.Cond, s.Eval}
}

func (s *Loop) Childs() []trait.ASTNode {
	return []trait.ASTNode{s.Cond, s.Eval}
}

func (s *Loop) Evaluate(ctx trait.Context) trait.EvalResult {
	for {
		// check cond
		var cv = s.Cond.Evaluate(ctx)
		if cv.CheckInterrupt() {
			return cv
		}
		if !cv.IsTrue() {
			return eval.ResultNone() // cond break
		}
		// do loop
		var bv = s.Eval.Evaluate(ctx)
		if bv.CheckInterrupt() {
			var ty = bv.EventType()
			if ty == eval.ResultEventTypeLoopBreak {
				//break
				bv.CleanEvent()
				return bv
			} else if ty == eval.ResultEventTypeLoopContinue {
				continue
			}
			// func return or fatal
			return bv
		}
	}
}
