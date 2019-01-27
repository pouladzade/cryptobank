package client

import (
	"context"
	"fmt"
	"github.com/cryptobank/config"
	"github.com/cryptobank/cryptobank"
	crb "github.com/cryptobank/cryptobank"
	"net"
	"zombiezen.com/go/capnproto2/rpc"
)

func NewRequest() *Request {
	ctx := context.Background()
	conf, err := config.LoadFromFile(config.Config_File)
	if err != nil {
		fmt.Println(err)
		conf = config.DefaultConfig()
	}
	conn, err := net.Dial(conf.RPC.Type, conf.RPC.Host+":"+conf.RPC.Port)
	if err != nil {
		fmt.Println(err)
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
