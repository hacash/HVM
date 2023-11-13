package eval

import (
	"fmt"
	"github.com/hacash/HVM/trait"
	"math/big"
)

func ComputeLeftRight(ctx trait.Context, l, r trait.ASTNode) (trait.EvalResult, *big.Int, *big.Int) {

	var lv = l.Evaluate(ctx)
	if lv.CheckInterrupt() {
		return lv, nil, nil
	}
	var rv = r.Evaluate(ctx)
	if rv.CheckInterrupt() {
		return lv, nil, nil
	}
	var ln = big.NewInt(0)
	if lv.IsTrue() {
		ln = ln.SetBytes(lv.RetValue())
	}
	var rn = big.NewInt(0)
	if rv.IsTrue() {
		rn = rn.SetBytes(rv.RetValue())
	}
	return nil, ln, rn
}

func DoCompute(ctx trait.Context, l trait.ASTNode, r trait.ASTNode, opfn func(l, r *big.Int) (*big.Int, trait.EvalResult)) trait.EvalResult {

	fte, ln, rn := ComputeLeftRight(ctx, l, r)
	if fte != nil {
		return fte // fatal error
	}
	// do add
	resn, fte := opfn(ln, rn)
	if fte != nil {
		return fte // fatal error
	}
	if resn == nil {
		return ResultFatalErr(fmt.Errorf("Compute result cannot be nil"))
	}
	var resbts = resn.Bytes()
	// check size
	var re = CheckNumberSize(len(resbts), "tip")
	if re != nil {
		return re
	}
	// ok
	return ResultValue(resbts)
}
