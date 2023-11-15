package trait

type Context interface {
	GetStack() Stack
	GetHeap() Heap
	GetExtendCallExecutor() ExtendCallExecutor
}
