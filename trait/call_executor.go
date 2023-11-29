package trait

/**
 *
 */
type ExtendCallExecutor interface {
	ExtendKindRange() uint8
	Parse(buf []byte, seek uint32) (VMAction, uint32, error)
	// return: ast_value, use_gas, fatal_err
	Evaluate(VMAction, Context) ([]byte, int32, error)
}
