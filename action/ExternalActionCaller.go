package action

import (
	"github.com/hacash/HVM/trait"
	"github.com/hacash/core/fields"
)

type StaticExternalActionCaller struct{}

/********** Dynamic **********/

type DynamicExternalActionCaller struct {
	ExternalActionKind fields.VarUint2
	ActionBodyData     trait.VMAction
}

func (elm DynamicExternalActionCaller) Kind() uint16 {
	return 11111
}

func (elm DynamicExternalActionCaller) IsBurning90PersentTxFees() bool {
	return false
}

func (elm DynamicExternalActionCaller) ChildActions() []trait.VMAction {
	return []trait.VMAction{elm.ActionBodyData}
}

func (elm DynamicExternalActionCaller) Childs() []trait.ASTNode {
	return []trait.ASTNode{elm.ActionBodyData}
}

func (elm DynamicExternalActionCaller) Evaluate(ctx trait.Context) trait.EvalResult {
	return nil
}
