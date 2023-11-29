package action

import (
	"fmt"
	"github.com/hacash/HVM/eval"
	"github.com/hacash/HVM/trait"
	"github.com/hacash/core/fields"
)

type ContractInvoke struct {
	ParamCount   fields.VarUint1
	Params       []trait.VMAction
	ContractAddr trait.VMAction
	FuncName     trait.VMAction
}

func (s *ContractInvoke) VMKind() uint8 {
	return 255
}

func (s *ContractInvoke) IsBurning90PersentTxFees() bool {
	return true
}

func (s *ContractInvoke) ChildActions() []trait.VMAction {
	var acts = []trait.VMAction{}
	acts = append(acts, s.Params...)
	acts = append(acts, s.ContractAddr)
	acts = append(acts, s.FuncName)
	return acts
}

/*
func (s *ContractInvoke) Childs() []trait.ASTNode {
	var acts = []trait.ASTNode{}
	acts = append(acts, s.Params...)
	acts = append(acts, s.ContractAddr)
	acts = append(acts, s.FuncName)
	return acts
}
*/

func (s *ContractInvoke) Parse(extca trait.ExtendCallExecutor, buf []byte, seek uint32) (uint32, error) {
	return 0, nil
}

func (s *ContractInvoke) Evaluate(ctx trait.Context) trait.EvalResult {
	return eval.ResultFatalErr(fmt.Errorf("not yet"))
}
