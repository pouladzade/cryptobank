package server

import (
	"fmt"
	"github.com/cryptobank/cryptobank"
	"net"
	"zombiezen.com/go/capnproto2/rpc"
)

func server(c net.Conn) error {

	main := cryptobank.CoreBanking_ServerToClient(Service{})

	conn := rpc.NewConn(rpc.StreamTransport(c), rpc.MainInterface(main.Client))
	// Wait for connection to abort.
	fmt.Println("Server is running....")
	err := conn.Wait()
	return err
}
