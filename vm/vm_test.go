package vm

import (
	"encoding/hex"
	"fmt"
	"math/big"
	"testing"
)

func Test1(t *testing.T) {

	//var vvv = fields.NewAmountSmallValue(1, 248).GetValue()
	//var bts = vvv.Bytes()
	//
	//fmt.Println(len(bts), hex.EncodeToString(bts), vvv.String())

	var vvv = big.NewInt(-10)
	vvv = vvv.Abs(vvv)
	vvv = vvv.Mod(vvv, big.NewInt(7))
	fmt.Println(vvv.String())

	var v2 = big.NewInt(512)
	fmt.Println(hex.EncodeToString(v2.Bytes()), v2.String())

}
