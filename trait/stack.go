package trait

import "bytes"

type Stack interface {
	Alloc(int, bool) (Stack, error)
	Write(int, *bytes.Buffer) error
	Read(int) (*bytes.Buffer, error)
}
