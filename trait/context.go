package trait

type Context interface {
	GetStack() Stack
	GetHeap() Heap
}
