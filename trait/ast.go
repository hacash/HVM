package trait

type EvalResult interface {
	FatalErr() error
	RetValue() []byte
	IsTrue() bool
	CheckInterrupt() bool
}

type ASTNode interface {
	Evaluate(Context) EvalResult
	Childs() []ASTNode
}
