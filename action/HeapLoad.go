package action

import (
	"fmt"
	"github.com/hacash/HVM/eval"
	"github.com/hacash/HVM/heap"
	"github.com/hacash/HVM/trait"
)

type HeapLoad struct {
	Key trait.VMAction
}

func (s *HeapLoad) VMKind() uint8 {
	return 17
}

func (s *HeapLoad) IsBurning90PersentTxFees() bool {
	return true
}

func (s *HeapLoad) ChildActions() []trait.VMAction {
	return []trait.VMAction{s.Key}
}

func (s *HeapLoad) Childs() []trait.ASTNode {
	return []trait.ASTNode{s.Key}
}

func (s *HeapLoad) Evaluate(ctx trait.Context) trait.EvalResult {
	var key = s.Key.Evaluate(ctx)
	if key.CheckInterrupt() {
		return key
	}
	var kp = key.RetValue()
	var kpn = len(kp)
	if kpn < 1 || kpn > heap.HEAP_ITEM_MAX_SIZE {
		return eval.ResultFatalErr(fmt.Errorf("HeapLoad key length overflow"))
	}
	// load
	var hap = ctx.GetHeap()
	var v, e = hap.Read(string(kp))
	if e != nil {
		return eval.ResultFatalErr(e) // heapoverflow
	}
	resv := v.Bytes()
	if len(resv) > heap.HEAP_ITEM_MAX_SIZE {
		return eval.ResultFatalErr(fmt.Errorf("HeapLoad value over size %d over max %d",
			len(resv), heap.HEAP_ITEM_MAX_SIZE)) // stackoverflow
	}
	// ok
	return eval.ResultValue(resv)
}
