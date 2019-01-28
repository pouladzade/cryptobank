## master  [![Build Status](https://api.travis-ci.org/pouladzade/cryptobank.svg?branch=master)](https://travis-ci.org/pouladzade/cryptobank)
## develop [![Build Status](https://api.travis-ci.org/pouladzade/cryptobank.svg?branch=develop)](https://travis-ci.org/pouladzade/cryptobank)

# Cryptobank
Online crypto-banking is a system that consist of a single bank that responsible for tracking accounts and performing 

cryptocurrency related operations, and the clients which initiate some operations.(a sample of usage capnp proto in golang)
## Compiling the code

You need to install [Go](https://golang.org/) (version 1.10.1 or higher)
<p>Then, <a href="https://capnproto.org/install.html" rel="nofollow">install the Cap'n Proto tools</a>.

After installing them, you can follow these steps to compile and build the cryptobank project:

```bash
mkdir -p $GOPATH/src/github.com/cryptobank
cd $GOPATH/src/github.com/cryptobank
git clone https://github.com/pouladzade/cryptobank.git
make
```
```
Two executable file will be created:

'bankclient' in this directory:
"$GOPATH/src/github.com/cryptobank/client/build"

 and 'bankserver' in this directory:
 $GOPATH/src/github.com/cryptobank/server/build
```

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
## bankserver(Cryptobank Server)
It's the server aplication which provides the services through capnp rpc.
and will manages the accounts in form of json file.
for starting server you need to provide a config.toml beside the server binaries,
however if server can not find the config.toml it will create it by default configuration.

#### config.toml
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

## bankclient(Cryptobank Client)
For initializing client you need to provide a config.toml beside the client binaries,
however if client can not find the config.toml it will create a default configuration file by itself.

#### config.toml
```js
[rpc]
  host = "127.0.0.1"
  port = "1362"
  type = "tcp"
```

Getting the the list of flags:

```
ahmad@ahmad-5570:~$ bankclient -h
Usage of bankclient:
  -accid string
    	a 32 bytes in hex-string format as AccountId for deleting or creating new account
  -amount string
    	an amount which will be use in transfer found
  -bal string
    	balance for creating new account
  -cmd string
    	function name(command) which you wanna send to server :
    		[CreateAccount|crt]
    		[DeleteAccount|del]
    		[TransferFunds|trf]
  -des string
    	a 32 bytes in hex-string format as AccountId of destination account in transfer found
  -name string
    	account holder name for creating new account
  -src string
    	a 32 bytes in hex-string format as AccountId of source account in transfer found

```
CreateAccount :
```
ahmad@ahmad-5570:~~$ bankclient -cmd="crt" -accid="a863fead151f388b781d62beea26712e59dd3af0e2f478da260f3c4aa5ee8904" -name="Ahmad" -bal="999999"

```
DeleteAccount :
```
ahmad@ahmad-5570:~~$ bankclient -cmd="del" -accid="a863fead151f388b781d62beea26712e59dd3af0e2f478da260f3c4aa5ee8904"

```

TransferFunds :
```
ahmad@ahmad-5570:~~$ bankclient -cmd="trf" -src="a863fead151f388b781d62beea26712e59dd3af0e2f478da260f3c4aa5ee8904" -des="9387024AD4E5C0645FCA292669889AACEDAF4B06D14E63D6C5C40F4B2A291588" -amount="1000"

```

bankclient and bank server sample log :

```
Client Side:

ahmad@ahmad-5570:~$ bankclient -cmd="crt" -accid="a863fead151f388b781d62beea26712e59dd3af0e2f478da260f3c4aa5ee8904" -name="Ahmad" -bal="999999"
2019/01/28 11:09:29 Calling CreateAccount:
 AccountId = a863fead151f388b781d62beea26712e59dd3af0e2f478da260f3c4aa5ee8904 , name = Ahmad , balance = 999999

2019/01/28 11:09:29 (code = 0, message = "Account succesfully created!")

ahmad@ahmad-5570:~$ bankclient -cmd="crt" -accid="9387024AD4E5C0645FCA292669889AACEDAF4B06D14E63D6C5C40F4B2A291588" -name="Max" -bal="999999"
2019/01/28 11:10:49 Calling CreateAccount:
 AccountId = 9387024ad4e5c0645fca292669889aacedaf4b06d14e63d6c5c40f4b2a291588 , name = Max , balance = 999999

2019/01/28 11:10:49 (code = 0, message = "Account succesfully created!")

ahmad@ahmad-5570:~$ bankclient -cmd="trf" -src="a863fead151f388b781d62beea26712e59dd3af0e2f478da260f3c4aa5ee8904" -des="9387024AD4E5C0645FCA292669889AACEDAF4B06D14E63D6C5C40F4B2A291588" -amount="1000"
2019/01/28 11:10:57 Calling TransferFunds: source = {[168 99 254 173 21 31 56 139 120 29 98 190 234 38 113 46 89 221 58 240 226 244 120 218 38 15 60 74 165 238 137 4]  0} , destination = {[147 135 2 74 212 229 192 100 95 202 41 38 105 136 154 172 237 175 75 6 209 78 99 214 197 196 15 75 42 41 21 136]  0} , amount = 1000

2019/01/28 11:10:57 (code = 0, message = "Amount Successfully transffered")

ahmad@ahmad-5570:~$ bankclient -cmd="del" -accid="a863fead151f388b781d62beea26712e59dd3af0e2f478da260f3c4aa5ee8904"
2019/01/28 11:11:54 Calling DeleteAccount: AccountId = {[168 99 254 173 21 31 56 139 120 29 98 190 234 38 113 46 89 221 58 240 226 244 120 218 38 15 60 74 165 238 137 4]  0} 

2019/01/28 11:11:54 (code = 0, message = "Account succesfully deleted!")
```
```
Server Side:

ahmad@ahmad-5570:~$ bankserver 
Listening on 127.0.0.1:1362

Server recieved a request :
2019/01/28 11:09:29 
{request: CreateAccount
params[accountId = [168 99 254 173 21 31 56 139 120 29 98 190 234 38 113 46 89 221 58 240 226 244 120 218 38 15 60 74 165 238 137 4], name=Ahmad, balance=999999]}
{"a863fead151f388b781d62beea26712e59dd3af0e2f478da260f3c4aa5ee8904":{"name":"Ahmad","balance":999999}}
2019/01/28 11:09:29 (code = 0, message = "Account succesfully created!")
2019/01/28 11:09:29 rpc: abort: rpc: aborted by remote: rpc: shutdown
Server recieved a request :
2019/01/28 11:10:49 
{request: CreateAccount
params[accountId = [147 135 2 74 212 229 192 100 95 202 41 38 105 136 154 172 237 175 75 6 209 78 99 214 197 196 15 75 42 41 21 136], name=Max, balance=999999]}
{"9387024ad4e5c0645fca292669889aacedaf4b06d14e63d6c5c40f4b2a291588":{"name":"Max","balance":999999},"a863fead151f388b781d62beea26712e59dd3af0e2f478da260f3c4aa5ee8904":{"name":"Ahmad","balance":999999}}
2019/01/28 11:10:49 (code = 0, message = "Account succesfully created!")
2019/01/28 11:10:49 rpc: abort: rpc: aborted by remote: rpc: shutdown
Server recieved a request :
2019/01/28 11:10:57 
{request: TransferFunds
2019/01/28 11:10:57 
{params[Source = [168 99 254 173 21 31 56 139 120 29 98 190 234 38 113 46 89 221 58 240 226 244 120 218 38 15 60 74 165 238 137 4], Destination=[147 135 2 74 212 229 192 100 95 202 41 38 105 136 154 172 237 175 75 6 209 78 99 214 197 196 15 75 42 41 21 136], Amount=1000]}
{"9387024ad4e5c0645fca292669889aacedaf4b06d14e63d6c5c40f4b2a291588":{"name":"Max","balance":998999},"a863fead151f388b781d62beea26712e59dd3af0e2f478da260f3c4aa5ee8904":{"name":"Ahmad","balance":1000999}}
2019/01/28 11:10:57 Amount Successfully transffered
2019/01/28 11:10:57 (code = 0, message = "Amount Successfully transffered")
2019/01/28 11:10:57 rpc: abort: rpc: aborted by remote: rpc: shutdown
Server recieved a request :
2019/01/28 11:11:54 
{request: DeleteAccount
params[accountId = [168 99 254 173 21 31 56 139 120 29 98 190 234 38 113 46 89 221 58 240 226 244 120 218 38 15 60 74 165 238 137 4]]}
{"9387024ad4e5c0645fca292669889aacedaf4b06d14e63d6c5c40f4b2a291588":{"name":"Max","balance":998999}}
2019/01/28 11:11:54 (code = 0, message = "Account succesfully deleted!")
2019/01/28 11:11:54 rpc: abort: rpc: aborted by remote: rpc: shutdown


```