package main

import (
	"fmt"
	"github.com/cryptobank/client"
)

func main() {
	var cli client.Cli
	cli.LoadFlags()
	err := cli.Commit()
	fmt.Println(err.Error())
}
