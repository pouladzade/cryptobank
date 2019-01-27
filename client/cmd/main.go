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
	var ahmad acm.Account
	ahmad.Name = "Ahmad"
	ahmad.Balance = 99999999
	ahmad.SetAccountIdString("94E6D699FC57B3575E8E5A56CA18CF9632430A31D566705B4C3CAA06134F58B0")
	err := req.CreateAccount(ahmad)
	if err != nil {
		fmt.Println(err)
	}
	req.Close()

	var max acm.Account
	req = client.NewRequest(conf)
	max.Name = "Max"
	max.Balance = 8888888
	max.SetAccountIdString("73757FAA063959ECDACAB1D845786F196A792811EA0D1E638AD0A1BD8B1DF03B")
	err = req.CreateAccount(max)
	if err != nil {
		fmt.Println(err)
	}
	req.Close()

	var alis acm.Account
	req = client.NewRequest(conf)
	alis.Name = "Alis"
	alis.Balance = 10000
	alis.SetAccountIdString("94D4F25C19FCEC53711FC77B839EBEF299E8467DCFB73A7AB504DE09912EBDFB")
	err = req.CreateAccount(alis)
	if err != nil {
		fmt.Println(err)
	}
	req.Close()

	var bob acm.Account
	req = client.NewRequest(conf)
	bob.Name = "Bob"
	bob.Balance = 30000000
	bob.SetAccountIdString("A863FEAD151F388B781D62BEEA26712E59DD3AF0E2F478DA260F3C4AA5EE8904")
	err = req.CreateAccount(bob)
	if err != nil {
		fmt.Println(err)
	}
	req.Close()

	req = client.NewRequest(conf)
	err = req.DeleteAccount(bob)
	if err != nil {
		fmt.Println(err)
	}
	req.Close()

	req = client.NewRequest(conf)
	err = req.TransferFunds(ahmad, max, 222222)
	if err != nil {
		fmt.Println(err)
	}
	req.Close()

}
