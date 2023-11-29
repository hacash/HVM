package machine

import "bytes"

func RunCode(code []byte, stack [][]byte) []*bytes.Buffer {

	var codelen = len(code)
	var opseek int = 0
	var stkptr int = 0

	for {
		switch code[opseek] {

		case PSH:
			l := opseek + 1
			n := code[l]
			dts := code[l : l+int(n)]
			stack[stkptr] = dts
			stkptr++
			opseek += 1 + int(n)

		case POP:
			stkptr--

		}
		if opseek >= codelen {
			// return nothing
			return []*bytes.Buffer{}
		}
	}

}
