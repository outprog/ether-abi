package utils

import (
	"fmt"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/math"
)

// WABI is wrap of ethereum ABI
type WABI struct {
	wrapABI abi.ABI
}

// New .
func NewWABI(abiJSON string) *WABI {
	wrapABI, err := abi.JSON(strings.NewReader(abiJSON))
	if err != nil {
		panic(err)
	}
	return &WABI{wrapABI}
}

func (w *WABI) Encode(method string, args ...string) (string, error) {
	inputs := w.wrapABI.Methods[method].Inputs
	if len(inputs) != len(args) {
		return "", fmt.Errorf("argument count mismatch: %d of %d", len(args), len(inputs))
	}

	var iArgs []interface{}
	for i, v := range args {
		switch inputs[i].Type.String() {
		case "address":
			iArgs = append(iArgs, common.HexToAddress(v))
		case "uint256":
			iArgs = append(iArgs, math.MustParseBig256(v))
		case "address[]":
			addrs := []common.Address{}
			value := strings.Replace(strings.Replace(v, "[", "", 1), "]", "", 1)
			values := strings.Split(value, ".")
			for _, vv := range values {
				addrs = append(addrs, common.HexToAddress(vv))
			}
			iArgs = append(iArgs, addrs)
		default:
			return "", fmt.Errorf("unsupported arg type")
		}
	}

	by, err := w.wrapABI.Pack(method, iArgs...)
	if err != nil {
		return "", err
	}
	return common.Bytes2Hex(by), nil
}
