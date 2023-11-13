package eval

import (
	"fmt"
	"github.com/hacash/HVM/trait"
)

const NUMBER_MAX_SIZE = 32

func CheckNumberSize(size int, tip string) trait.EvalResult {
	if size > NUMBER_MAX_SIZE {
		return ResultFatalErr(fmt.Errorf("Number operate <%s> result size %d overflow %d",
			tip, size, NUMBER_MAX_SIZE))
	}
	return nil
}
