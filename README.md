# ether-abi
A cli tool for decode &amp; encode Ethereum ABI

## Installation

```
go get github.com/outprog/ether-abi
```

## Examples

ABI must json file.

- abi.json

```
[
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
]
```

#### encode

```
ether-abi --file abi.json --type encode --method withdraw --args 10000
# 0x2e1a7d4d0000000000000000000000000000000000000000000000000000000000002710
```

#### decode

## Development

```
go get -d github.com/outprog/ether-abi
cd $GOPATH/src/github.com/outprog/ether-abi
sh dep.sh
```

Run example:

```
go run ./cli.go --file abi.example.json --method withdraw --args 10000
# 0x2e1a7d4d0000000000000000000000000000000000000000000000000000000000002710
```

