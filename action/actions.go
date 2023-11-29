package action

import (
	"fmt"
	"github.com/hacash/HVM/trait"
)

func ParseVMAction(extca trait.ExtendCallExecutor, buf []byte, seek uint32) (trait.VMAction, uint32, error) {
	var k = buf[seek]
	var act trait.VMAction = nil
	switch k {
	//case 255:
	//	act = &ContractInvoke{}
	//case 254:
	//	act = &DynamicExternalActionCall{}
	case 200:
		act = &List{}
	}
	if act == nil {
		return nil, 0, fmt.Errorf("VM ParseAction Error: cannot find kind <%d>", k)
	}
	// parse
	seek2, e := act.Parse(extca, buf, seek+1)
	if e != nil {
		return nil, 0, e // parse error
	}
	// ok
	return act, seek2, nil
}

func ParseAction(extca trait.ExtendCallExecutor, buf []byte, seek uint32) (trait.VMAction, uint32, error) {
	var kdmx = extca.ExtendKindRange()
	var vk = buf[seek]
	if vk <= kdmx {
		// is external action
		var act, sk, e = extca.Parse(buf, seek)
		if e != nil {
			return nil, 0, e
		}
		// ok
		return &StaticExternalActionCall{act}, sk, nil
	}
	// try vm action
	return ParseVMAction(extca, buf, seek)
}
