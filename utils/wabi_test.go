package utils

import (
	"math/big"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestEncode(t *testing.T) {
	json := `[
		{
			"constant": false,
			"inputs": [
				{
					"name": "wad",
					"type": "uint256"
				}
			],
			"name": "withdraw",
			"outputs": [],
			"payable": false,
			"stateMutability": "nonpayable",
			"type": "function"
		}
	]`

	wabi := New(json)

	hex, err := wabi.Encode("withdraw", big.NewInt(10000000000000000))
	require.NoError(t, err)
	assert.Equal(t, "2e1a7d4d000000000000000000000000000000000000000000000000002386f26fc10000", hex)
}
