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
ether-abi --file abi.json --type encode --method setAllowance --args \[0xc02aaa39b223fe8d0a0e5c4f27ead9083c756cc2.0x0f5d2fb29fb7d3cfee444a200298f468908cc942\],0x2240dab907db71e64d3e0dba4800c83b5c502d4e 
```

notice:

- `\` before `[`,`]`
- splite `.` in `[]`

#### decode

TODO

## Development

```
go get -d github.com/outprog/ether-abi
cd $GOPATH/src/github.com/outprog/ether-abi
sh dep.sh
```

Run example:

```
go run ./cli.go --file abi.example.json --type encode --method setAllowance --args \[0xc02aaa39b223fe8d0a0e5c4f27ead9083c756cc2.0x0f5d2fb29fb7d3cfee444a200298f468908cc942\],0x2240dab907db71e64d3e0dba4800c83b5c502d4e
#0x46920bad00000000000000000000000000000000000000000000000000000000000000400000000000000000000000002240dab907db71e64d3e0dba4800c83b5c502d4e0000000000000000000000000000000000000000000000000000000000000002000000000000000000000000c02aaa39b223fe8d0a0e5c4f27ead9083c756cc20000000000000000000000000f5d2fb29fb7d3cfee444a200298f468908cc942
```

