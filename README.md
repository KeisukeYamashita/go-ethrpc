# go-ethrpc

[![Open Source Love](https://badges.frapsoft.com/os/v1/open-source.svg?v=103)](https://github.com/ellerbrock/open-source-badges/)
[![License: Apache 2.0](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](https://opensource.org/licenses/Apache-2.0)
[![CircleCI](https://circleci.com/gh/KeisukeYamashita/go-ethrpc/tree/develop.svg?style=svg)](https://circleci.com/gh/KeisukeYamashita/go-ethrpc/tree/develop)
[![Job Status](https://inspecode.rocro.com/badges/github.com/KeisukeYamashita/go-ethrpc/status?token=p273sFfQS5RMvjwD2lHfaiXcFW-B3-Ahs32p8DnrK40)](https://inspecode.rocro.com/jobs/github.com/KeisukeYamashita/go-ethrpc/latest?completed=true)
[![Report](https://inspecode.rocro.com/badges/github.com/KeisukeYamashita/go-ethrpc/report?token=p273sFfQS5RMvjwD2lHfaiXcFW-B3-Ahs32p8DnrK40&branch=master)](https://inspecode.rocro.com/reports/github.com/KeisukeYamashita/go-ethrpc/branch/master/summary)


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
  ethrpc "github.com/KeisukeYamashita/go-ethrpc"
  )

func main() {
	c := NewRPCClient(os.Getenv("ETHD_ENDPOINT"))
	address := "0x809826cceAb68c387726af962713b64Cb5Cb3CCA"
	balance := c.GetBalance(address)
	fmt.Print(balance) // 0.13514 ETH
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
    "0x809826cceAb68c387726af962713b64Cb5Cb3CCA"
  ]
}'
```

It'll return a JSON.

### Available Methods

| method| discription |
|:----:|:----:|
| eth_blockNumber | gets the most recent block height. |


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
To contribute, just send us a pull request!
If it is valid, you will be added on the contribution doc in `/doc/contributor.md` .

## Other Blockchain Libraries
There is a bitcoin version for this, it might help your application.

- [go-btcrpc](https://github.com/KeisukeYamashita/go-btcrpc)

## License
Copyright 2018 Keisuke Yamashita Hayashi Masaya.

Licensed under the Apache 2.0 license.
