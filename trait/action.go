package trait

type VMAction interface {

	//interfaces.Field

	ASTNode

	// the action type number
	Kind() uint16
	// burning fees
	IsBurning90PersentTxFees() bool // Whether to destroy 90% of the transaction cost of this transaction
	// chinds
	ChildActions() []VMAction
}
