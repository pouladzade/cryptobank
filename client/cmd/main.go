package main

import (
	"fmt"
	"github.com/cryptobank/client"
)

func main(){
	req := client.NewRequest()
	if req ==nil{
		panic(fmt.Errorf("Unable to create the request!"))
	}
	acc := make([]byte,client.AccountSize())
	err := req.CreateAccount(acc,"Ahmad",100000)
	if err != nil{
		fmt.Println(err)
	}
	req.Close()

	req = client.NewRequest()
	if req ==nil{
		panic(fmt.Errorf("Unable to create the request!"))
	}
	err = req.DeleteAccount(acc)
	if err != nil{
		fmt.Println(err)
	}
	req.Close()

	req = client.NewRequest()
	if req ==nil{
		panic(fmt.Errorf("Unable to create the request!"))
	}
	err = req.TransferFunds(acc,acc,1000)
	if err != nil{
		fmt.Println(err)
	}
	req.Close()
}