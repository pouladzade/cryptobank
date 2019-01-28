package client

import (
	"context"
	"flag"
	"fmt"
	"github.com/cryptobank/acm"
	"github.com/cryptobank/config"
	"github.com/cryptobank/cryptobank"
	crb "github.com/cryptobank/cryptobank"
	"net"
	"strconv"
	"zombiezen.com/go/capnproto2/rpc"
)

func NewRequest(conf *config.Config) *Request {
	ctx := context.Background()
	conn, err := net.Dial(conf.RPC.Type, conf.RPC.Host+":"+conf.RPC.Port)
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}
	rpcconn := rpc.NewConn(rpc.StreamTransport(conn))
	if rpcconn == nil {
		fmt.Println(fmt.Errorf("Can not create RPC connection!"))
		return nil
	}
	cb := crb.CoreBanking{Client: rpcconn.Bootstrap(ctx)}
	req := Request{cb, ctx, rpcconn}

	return &req
}

func AccountSize() uint8 {
	return cryptobank.AccountSize
}

type Cli struct {
	cmd     *string
	balance *string
	name    *string
	accId   *string
	src     *string
	des     *string
	amount  *string
}

func (cl *Cli) LoadFlags() {

	cl.cmd = flag.String("cmd", "", "function name(command) which you wanna send to server :\n\t[CreateAccount|crt]\n\t[DeleteAccount|del]\n\t[TransferFunds|trf]")
	cl.balance = flag.String("bal", "", "balance for creating new account")
	cl.name = flag.String("name", "", "account holder name for creating new account")
	cl.accId = flag.String("accid", "", "a 32 bytes in hex-string format as AccountId for deleting or creating new account")
	cl.src = flag.String("src", "", "a 32 bytes in hex-string format as AccountId of source account in transfer found")
	cl.des = flag.String("des", "", "a 32 bytes in hex-string format as AccountId of destination account in transfer found")
	cl.amount = flag.String("amount", "", "an amount which will be use in transfer found")

	flag.Parse()
}

func (cl *Cli) Commit() error {
	if cl.cmd == nil {
		return fmt.Errorf("Error : Please specifiy a function name using -cmd flag")
	}
	switch *cl.cmd {
	case "CreateAccount":
		return cl.createAccount()
	case "DeleteAccount":
		return cl.deleteAccount()
	case "TransferFunds":
		return cl.transferFunds()
	}
	switch *cl.cmd {
	case "crt":
		return cl.createAccount()
	case "del":
		return cl.deleteAccount()
	case "trf":
		return cl.transferFunds()
	}
	return fmt.Errorf("Error : Please specifiy a function name using -cmd flag")
}

func (cl *Cli) createAccount() error {
	conf := config.LoadCreateConfig()
	req := NewRequest(conf)

	if req == nil {
		return (fmt.Errorf("Unable to create the request!"))
	}
	defer req.Close()
	var acc acm.Account
	acc.Name = *cl.name
	bal, err := strconv.ParseUint(*cl.balance, 10, 64)
	if err != nil {
		return err
	}
	acc.Balance = bal
	acc.SetAccountIdString(*cl.accId)
	err = req.CreateAccount(acc)
	return err
}

func (cl *Cli) deleteAccount() error {
	conf := config.LoadCreateConfig()
	req := NewRequest(conf)

	if req == nil {
		return (fmt.Errorf("Unable to create the request!"))
	}
	defer req.Close()
	var acc acm.Account
	acc.SetAccountIdString(*cl.accId)
	err := req.DeleteAccount(acc)
	return err
}

func (cl *Cli) transferFunds() error {
	conf := config.LoadCreateConfig()
	req := NewRequest(conf)

	if req == nil {
		return (fmt.Errorf("Unable to create the request!"))
	}
	defer req.Close()
	var src, des acm.Account
	src.SetAccountIdString(*cl.src)
	des.SetAccountIdString(*cl.des)
	amount, err := strconv.ParseUint(*cl.amount, 10, 64)
	if err != nil {
		return err
	}
	err = req.TransferFunds(src, des, amount)
	return err
}
