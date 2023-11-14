package eval

import (
	"fmt"
	"github.com/hacash/HVM/stack"
	"github.com/hacash/HVM/trait"
	"math/big"
)

func CalcStackPtr(ctx trait.Context, ptr trait.ASTNode) (int, trait.EvalResult) {
	var rv = ptr.Evaluate(ctx)
	if rv.CheckInterrupt() {
		return 0, rv
	}
	if !rv.IsTrue() {
		return 0, ResultFatalErr(fmt.Errorf("StackLoad ptr num error"))
	}
	// get ptr
	var overr = ResultFatalErr(fmt.Errorf("StackLoad ptr num overflow"))
	var pbts = rv.RetValue()
	if len(pbts) > 4 {
		return 0, overr
	}
	var ptrb = big.NewInt(0).SetBytes(pbts).Uint64()
	if ptrb > uint64(stack.STACK_TOTAL_LENGTH) {
		return 0, overr
	}
	return int(ptrb), nil
}
