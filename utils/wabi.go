package utils

import (
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
)

// WABI is wrap of ethereum ABI
type WABI struct {
	wrapABI abi.ABI
}

// New .
func New(abiJSON string) *WABI {
	wrapABI, err := abi.JSON(strings.NewReader(abiJSON))
	if err != nil {
		panic(err)
	}
	return &WABI{wrapABI}
}

func (w *WABI) Encode(method string, args ...interface{}) (string, error) {
	by, err := w.wrapABI.Pack(method, args...)
	if err != nil {
		return "", err
	}
	return common.Bytes2Hex(by), nil
}
