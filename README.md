# go-ethrpc

[![Open Source Love](https://badges.frapsoft.com/os/v1/open-source.svg?v=103)](https://github.com/ellerbrock/open-source-badges/)
[![License: Apache 2.0](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](https://opensource.org/licenses/Apache-2.0)


go-ethrpc is a Go library use for interacting to the ethereum node from your server with JSON-RPC which is a standard protocol for blockchain.

This library provides the easiest way to send JSON-RPC.

## Installation
Use go get to install and update.

```
go get github.com/KeisukeYamashita/go-ethrpc.git
```

## Setup
You need to setup environmental variables.

Firstly, copy the `.env.sample` as `.env`

```
cp .env.sample .env
```

Setup in your `.env`.

```
GETH_ENDPOINT: NODE_ENDPOINT
```

## Usage and Example
This shows you that easiest request to the node which is getting the infos.

```go
package main

import (
  "fmt"
  btcrpc "github.com/KeisukeYamashita/go-ethrpc"
  )

func main() {
	c := NewRPCClient(os.Getenv("ETHD_ENDPOINT"))
	address := "my88QLpf2RYYDdNMmDwYvfx6TFc6NXaELa"
	balance := c.GetBalance(address)
	fmt.Print(balance) // 0.13514 BTC
}
```

### Equal curl command

```
curl -X "POST" "<YOUR_GETH_NODE>" \
     -H 'Content-Type: application/json; charset=utf-8' \
     -d $'{
  "jsonrpc":"2.0",
  "method": "getbalance",
  "id": "1",
  "params": [
    "my88QLpf2RYYDdNMmDwYvfx6TFc6NXaELa"
  ]
}'
```

It'll return a JSON.

### Available Methods

| method| discription |
|:----:|:----:|
| getNewAddress | creates a new address to your account. |
| getBalance | get the balance of your address. |
| getBlockHash | get the hash of the block. |
| getBlock | get the block by hash of the block. |
| getBlockCount | get the newest block count by hash of the block. |
| decodeTransaction | decode the raw transaction to humanreadable transaction by hash of the block. |
| getrawTransaction | get the raw transaction hash block count by hash of the block. |


## Use tests
Set up your environmental valiables in `.env` to conduct this test.

```
cp .env.sample .env
```

Then write in your endpoint in this file.


Finally run your test. It will pass if your bitcoin node is setted up correctly.

```
GO_ENV=test go test ethrpc
```

## Contribution
To contribute, just send me a pull request!
If it is valid, you will be added on the contribution doc in `/doc/contributor.md` .

## Other Blockchain Libraries
There is a bitcoin version for this, it might help your application.

- [go-btcrpc](https://github.com/KeisukeYamashita/go-btcrpc)

## License
Copyright 2018 Keisuke Yamashita Hayashi Masaya.
Licensed under the Apache 2.0 license.
