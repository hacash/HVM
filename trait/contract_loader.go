package trait

import "github.com/hacash/core/fields"

type CodePackage interface {
	CodeTy() uint8       // code type
	StackSize() uint8    // max 128
	ASTNodes() []ASTNode // vm script or nil
	ByteCodes() []byte   // vm byte codes or nil
}

type ContractLoader interface {
	Load(addr fields.Address, fnsig []byte)
}
