package trait

type EvalResult interface {
	FatalErr() error
	RetValue() []byte
	IsTrue() bool
	CheckInterrupt() bool
	EventType() uint8
	CleanEvent()
}

type ASTNode interface {
	Evaluate(Context) EvalResult
	//Childs() []ASTNode
}
