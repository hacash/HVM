package eval

import (
	"fmt"
	"github.com/hacash/HVM/stack"
	"github.com/hacash/HVM/trait"
)

const ResultEventTypeFatal uint8 = 1
const ResultEventTypeFunctionReturn uint8 = 2
const ResultEventTypeLoopBreak uint8 = 3

type ResultObj struct {
	eventTy  uint8 // 0: normal   1: fatal panic   2: func return   3: loop break
	fatalErr error
	retValue []byte
}

func (r *ResultObj) FatalErr() error {
	return r.fatalErr
}

func (r *ResultObj) RetValue() []byte {
	return r.retValue
}

func (r *ResultObj) IsTrue() bool {
	if r.fatalErr != nil {
		return false
	}
	if r.retValue == nil {
		return false
	}
	var size = len(r.retValue)
	if 0 == size {
		return false
	}
	if 1 == size && r.retValue[0] == 0 {
		return false
	}
	return true
}

func (r *ResultObj) IsNone() bool {
	if r.fatalErr != nil {
		return true
	}
	if r.retValue == nil {
		return true
	}
	if len(r.retValue) == 0 {
		return true
	}
	// have value
	return false
}

func (r *ResultObj) CheckInterrupt() bool {
	if r.fatalErr != nil {
		return true
	}
	if r.eventTy > 0 {
		return true
	}
	if len(r.retValue) > stack.STACK_ITEM_MAX_SIZE {
		r.fatalErr = fmt.Errorf("Result value item size %d overflow %d",
			len(r.retValue), stack.STACK_ITEM_MAX_SIZE)
		return true
	}
	return false
}

func (r *ResultObj) IsFatal() bool {
	if r.fatalErr != nil {
		return true
	}
	return false
}

/*************/

func ResultFatalErr(err error) trait.EvalResult {
	return &ResultObj{
		ResultEventTypeFatal,
		err,
		nil,
	}
}

func ResultValue(data []byte) trait.EvalResult {
	return &ResultObj{
		0,
		nil,
		data,
	}
}

func ResultValueTy(data []byte, ty uint8) trait.EvalResult {
	return &ResultObj{
		ty,
		nil,
		data,
	}
}

func ResultNone() trait.EvalResult {
	return &ResultObj{
		0,
		nil,
		[]byte{},
	}
}

func ResultFalse() trait.EvalResult {
	return &ResultObj{
		0,
		nil,
		[]byte{0},
	}
}

func ResultOK() trait.EvalResult {
	return &ResultObj{
		0,
		nil,
		[]byte{1},
	}
}
