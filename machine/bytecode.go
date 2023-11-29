package machine

const (
	ERR byte = 10 // if true throw error of pop
	ADD byte = 11
	JMP byte = 12 // goto
	RET byte = 13 // function return
	IFJ byte = 14 // if false jump
	LOD byte = 15 //
	POP byte = 16 // pop stack and drop
	PSH byte = 17 // push value to stack

)
