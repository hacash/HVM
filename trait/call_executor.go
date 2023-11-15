package trait

/**
 *
 */
type ExtendCallExecutor interface {
	ExtendKindRange() uint8
	Parse(buf []byte, seek uint32) (VMAction, uint32, error)
	Evaluate(VMAction, Context) ([]byte, error)
}
