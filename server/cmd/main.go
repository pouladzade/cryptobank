package main

import (
	"fmt"
	"github.com/cryptobank/cryptobank"
	"github.com/cryptobank/server"
	"github.com/cryptobank/config"
	"net"
	"os"
	"zombiezen.com/go/capnproto2/rpc"
)

func main() {
	conf, err := config.LoadFromFile(config.Config_File)
	if err != nil {
		fmt.Println(err)
		conf = config.DefaultConfig()
	}

	l, err := net.Listen(conf.RPC.Type, conf.RPC.Host+":"+conf.RPC.Port)
	if err != nil {
		fmt.Println("Error listening:", err.Error())
		os.Exit(1)
	}
	defer l.Close()

	fmt.Println("Listening on " + conf.RPC.Host + ":" + conf.RPC.Port)
	for {
		conn, err := l.Accept()
		if err != nil {
			fmt.Println("Error accepting: ", err.Error())
			os.Exit(1)
		}
		go serve(conn)
	}
}

func serve(c net.Conn) {
	main := cryptobank.CoreBanking_ServerToClient(server.Service{})
	conn := rpc.NewConn(rpc.StreamTransport(c), rpc.MainInterface(main.Client))
	fmt.Println("Waiting for request....")
	err := conn.Wait()
	if err !=nil{
		fmt.Println(err)
	}

}
