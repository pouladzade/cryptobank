package main

import (
	"fmt"
	"github.com/cryptobank/acm"
	"github.com/cryptobank/client"
	"github.com/cryptobank/config"
)

func main() {
	conf := config.LoadCreateConfig()
	req := client.NewRequest(conf)
	if req == nil {
		panic(fmt.Errorf("Unable to create the request!"))
	}
	var acc1 acm.Account
	acc1.Name = "Ahmad"
	acc1.Balance = 99999999
	acc1.SetAccountIdString("94E6D699FC57B3575E8E5A56CA18CF9632430A31D566705B4C3CAA06134F58B0")
	err := req.CreateAccount(acc1)
	if err != nil {
		fmt.Println(err)
	}
	req.Close()

	req = client.NewRequest(conf)
	var acc2 acm.Account
	acc2.Name = "Max"
	acc2.Balance = 8888888
	acc2.SetAccountIdString("73757FAA063959ECDACAB1D845786F196A792811EA0D1E638AD0A1BD8B1DF03B")
	err = req.CreateAccount(acc2)
	if err != nil {
		fmt.Println(err)
	}
	req.Close()
	/*
		req = client.NewRequest()
		if req == nil {
			panic(fmt.Errorf("Unable to create the request!"))
		}
		err = req.DeleteAccount(acc)
		if err != nil {
			fmt.Println(err)
		}
		req.Close()

		req = client.NewRequest()
		if req == nil {
			panic(fmt.Errorf("Unable to create the request!"))
		}
		err = req.TransferFunds(acc, acc, 1000)
		if err != nil {
			fmt.Println(err)
		}
		req.Close()
	*/
}
