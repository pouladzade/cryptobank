package main

import (
	"fmt"
	"github.com/cryptobank/config"
	"github.com/cryptobank/cryptobank"
	"github.com/cryptobank/server"
	"github.com/cryptobank/server/cdb"
	"net"
	"os"
	"zombiezen.com/go/capnproto2/rpc"
)

func main() {

	conf := config.LoadCreateConfig()
	l, err := net.Listen(conf.RPC.Type, conf.RPC.Host+":"+conf.RPC.Port)
	if err != nil {
		fmt.Printf("Error listening: [%s] \n\n", err.Error())
		os.Exit(1)
	}
	defer l.Close()
	var db cdb.CryptoDb
	db.LoadDb()
	var service server.Service
	service.SetDb(&db)

	fmt.Println("Listening on " + conf.RPC.Host + ":" + conf.RPC.Port)
	for {
		conn, err := l.Accept()
		if err != nil {
			fmt.Printf("Error accepting:[%s] \n\n", err.Error())
			os.Exit(1)
		}
		go serve(conn, &service)
	}
	db.Commit()
}

func serve(c net.Conn, service *server.Service) {
	main := cryptobank.CoreBanking_ServerToClient(service)
	conn := rpc.NewConn(rpc.StreamTransport(c), rpc.MainInterface(main.Client))
	fmt.Println("Server recieved a request :")
	err := conn.Wait()
	if err != nil {
		fmt.Println(err.Error())
	}
}
