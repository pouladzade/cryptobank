using Go = import "/go.capnp";
@0xaaa3acc39d94e0f8;
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
  getAccountInfo@3 (accountId :Data) -> (name :Text, balance : UInt64,res :Response);
}

