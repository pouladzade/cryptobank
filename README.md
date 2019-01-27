# master [![Build Status](https://api.travis-ci.org/pouladzade/cryptobank.svg?branch=master)](https://travis-ci.org/pouladzade/cryptobank)
# develop [![Build Status](https://api.travis-ci.org/pouladzade/cryptobank.svg?branch=develop)](https://travis-ci.org/pouladzade/cryptobank)


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

# Will be updated very soon....
