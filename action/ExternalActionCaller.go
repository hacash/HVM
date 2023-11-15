package action

import (
	"fmt"
	"github.com/hacash/HVM/eval"
	"github.com/hacash/HVM/trait"
)

/********** Static **********/

type StaticExternalActionCaller struct {
	ExternalAction trait.VMAction
}

func (elm StaticExternalActionCaller) VMKind() uint8 {
	return 0
}

func (elm StaticExternalActionCaller) IsBurning90PersentTxFees() bool {
	return false
}

func (s *StaticExternalActionCaller) Parse(extca trait.ExtendCallExecutor, buf []byte, seek uint32) (uint32, error) {
	panic(any("Never call StaticExternalActionCaller.Parse"))
}

func (elm StaticExternalActionCaller) ChildActions() []trait.VMAction {
	return []trait.VMAction{elm.ExternalAction}
}

func (elm StaticExternalActionCaller) Childs() []trait.ASTNode {
	return []trait.ASTNode{elm.ExternalAction}
}

func (elm StaticExternalActionCaller) Evaluate(ctx trait.Context) trait.EvalResult {
	// check
	var exec = ctx.GetExtendCallExecutor()
	var evmk = elm.ExternalAction.VMKind()
	if evmk > exec.ExtendKindRange() {
		return eval.ResultFatalErr(fmt.Errorf("ExternalAction VMKind %d error", evmk))
	}
	// evaluate
	var resv, err = exec.Evaluate(elm.ExternalAction, ctx)
	if err != nil {
		return eval.ResultFatalErr(err) // error
	}
	// ok
	return eval.ResultValue(resv)
}

/********** Dynamic **********/

type DynamicExternalActionCaller struct {
	ActionOverallData trait.VMAction
}

func (elm DynamicExternalActionCaller) VMKind() uint8 {
	return 11
}

func (elm DynamicExternalActionCaller) IsBurning90PersentTxFees() bool {
	return false
}

func (elm DynamicExternalActionCaller) ChildActions() []trait.VMAction {
	return []trait.VMAction{elm.ActionOverallData}
}

func (elm DynamicExternalActionCaller) Childs() []trait.ASTNode {
	return []trait.ASTNode{elm.ActionOverallData}
}

func (elm DynamicExternalActionCaller) Evaluate(ctx trait.Context) trait.EvalResult {
	var exec = ctx.GetExtendCallExecutor()
	var actdts = elm.ActionOverallData.Evaluate(ctx)
	if actdts.CheckInterrupt() {
		return actdts
	}
	var data = actdts.RetValue()
	if len(data) < 2 {
		return eval.ResultFatalErr(fmt.Errorf("Action data length %d less than kind length %d",
			len(data), 2))
	}
	// build action
	var actobj, _, e = exec.Parse(data, 0)
	if e != nil {
		return eval.ResultFatalErr(fmt.Errorf("Action build error: %s ", e))
	}
	// evaluate
	var resv, err = ctx.GetExtendCallExecutor().Evaluate(actobj, ctx)
	if err != nil {
		return eval.ResultFatalErr(err) // error
	}
	// ok
	return eval.ResultValue(resv)

}
