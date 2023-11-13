package trait

import "bytes"

type Heap interface {
	Write(string, *bytes.Buffer) error
	Read(string) (*bytes.Buffer, error)
}
