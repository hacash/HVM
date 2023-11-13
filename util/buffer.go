package util

import (
	"bytes"
)

func CopyVMBuffer(src *bytes.Buffer) *bytes.Buffer {
	var sbt = src.Bytes()
	var cpbts = make([]byte, len(sbt))
	copy(cpbts, sbt)
	return bytes.NewBuffer(cpbts)
}
