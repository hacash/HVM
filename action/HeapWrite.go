package action

import (
	"bytes"
	"fmt"
	"github.com/hacash/HVM/eval"
	"github.com/hacash/HVM/heap"
	"github.com/hacash/HVM/trait"
)

type HeapWrite struct {
	Key trait.VMAction
	Val trait.VMAction
}

func (s *HeapWrite) VMKind() uint8 {
	return 16
}

func (s *HeapWrite) IsBurning90PersentTxFees() bool {
	return true
}

func (s *HeapWrite) ChildActions() []trait.VMAction {
	return []trait.VMAction{s.Key, s.Val}
}

func (s *HeapWrite) Childs() []trait.ASTNode {
	return []trait.ASTNode{s.Key, s.Val}
}

func (s *HeapWrite) Evaluate(ctx trait.Context) trait.EvalResult {
	var key = s.Key.Evaluate(ctx)
	if key.CheckInterrupt() {
		return key
	}
	var kp = key.RetValue()
	var kpn = len(kp)
	if kpn < 1 || kpn > heap.HEAP_ITEM_MAX_SIZE {
		return eval.ResultFatalErr(fmt.Errorf("HeapWrite key length overflow"))
	}
	// data
	var dt = s.Key.Evaluate(ctx)
	if dt.CheckInterrupt() {
		return dt
	}
	var dts = dt.RetValue()
	if len(dts) > heap.HEAP_ITEM_MAX_SIZE {
		return eval.ResultFatalErr(fmt.Errorf("HeapWrite value over size %d over max %d",
			len(dts), heap.HEAP_ITEM_MAX_SIZE)) // stackoverflow
	}
	// write
	var hap = ctx.GetHeap()
	var e = hap.Write(string(kp), bytes.NewBuffer(dts))
	if e != nil {
		return eval.ResultFatalErr(e) // heapoverflow
	}
	// ok
	return eval.ResultValue(dts)
}
