## master  [![Build Status](https://api.travis-ci.org/pouladzade/cryptobank.svg?branch=master)](https://travis-ci.org/pouladzade/cryptobank)
## develop [![Build Status](https://api.travis-ci.org/pouladzade/cryptobank.svg?branch=develop)](https://travis-ci.org/pouladzade/cryptobank)

# Cryptobank
Online crypto-banking is a system that consist of a single bank that responsible for tracking accounts and performing 

cryptocurrency related operations, and the clients which initiate some operations.(a sample of usage capnp proto in golang)

This is the capn'p proto schema :
```js
using Go = import "/go.capnp";
@0x85d3acc39d94e0f8;
$Go.package("cryptobank");
$Go.import("cryptobank");


const accountSize :UInt8 = 32;

struct Response{
    code    @0 : Int32;
    message @1 : Text;
}

interface CoreBanking {
  createAccount @0 (accountId :Data, name :Text, balance : UInt64) -> (res:Response);
  deleteAccount @1 (accountId : Data) -> (res :Response);
  transferFunds @2 (source : Data, destination : Data, amount : UInt64) -> (res :Response);
}

```
## bankserver
It's the server aplication which provides the services through capnp rpc.
and will manages the accounts in form of json file.
for starting server you need to provide a config.toml beside the server binaries,
however if server can not find the config.toml it will create it by default configuration.

## config.toml
```js
[rpc]
  host = "127.0.0.1"
  port = "1362"
  type = "tcp"
```

Server will save and load accounts in a json file in the default director.
./db/cryptodb.json
```js
{
    "73757faa063959ecdacab1d845786f196a792811ea0d1e638ad0a1bd8b1df03b": {
        "name": "Max",
        "balance": 8888888
    },
    "94d4f25c19fcec53711fc77b839ebef299e8467dcfb73a7ab504de09912ebdfb": {
        "name": "Alis",
        "balance": 10000
    },
    "94e6d699fc57b3575e8e5a56ca18cf9632430a31d566705b4c3caa06134f58b0": {
        "name": "Ahmad",
        "balance": 99999999
    },
    "a863fead151f388b781d62beea26712e59dd3af0e2f478da260f3c4aa5ee8904": {
        "name": "Bob",
        "balance": 30000000
    }
}
```  
# Will be updated very soon....
