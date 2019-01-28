package client

import (
	"context"
	"fmt"
	"github.com/cryptobank/acm"
	crb "github.com/cryptobank/cryptobank"
	"log"
	"zombiezen.com/go/capnproto2/rpc"
)

type Request struct {
	cb   crb.CoreBanking
	ctx  context.Context
	conn *rpc.Conn
}

func (r Request) CreateAccount(acc acm.Account) error {
	log.Printf("Calling CreateAccount:\n AccountId = %s , name = %s , balance = %d\n\n",
		acc.AccountIdString(), acc.Name, acc.Balance)
	result, err := r.cb.CreateAccount(r.ctx, func(p crb.CoreBanking_createAccount_Params) error {
		p.SetBalance(acc.Balance)
		if p.SetAccountId(acc.AccountId()) == nil && p.SetName(acc.Name) == nil {
			return nil
		}
		return fmt.Errorf("Error : can not set parameters!")
	}).Struct()
	if err != nil {
		log.Printf("Error in CreateAccount : %v", err)
		return err
	}
	res, err := result.Res()
	log.Println(res)

	return err
}

func (r Request) DeleteAccount(acc acm.Account) error {
	log.Printf("Calling DeleteAccount: AccountId = %v \n\n", acc)
	result, err := r.cb.DeleteAccount(r.ctx, func(p crb.CoreBanking_deleteAccount_Params) error {
		err := p.SetAccountId(acc.AccountId())
		return err
	}).Struct()
	if err != nil {
		log.Printf("Error in DeleteAccount : %v", err)
		return err
	}
	res, err := result.Res()
	log.Println(res)
	return err
}

func (r Request) TransferFunds(src, des acm.Account, amount uint64) error {
	log.Printf("Calling TransferFunds: source = %v , destination = %v , amount = %d\n\n", src, des, amount)
	result, err := r.cb.TransferFunds(r.ctx, func(p crb.CoreBanking_transferFunds_Params) error {
		p.SetAmount(amount)
		if p.SetSource(src.AccountId()) == nil && p.SetDestination(des.AccountId()) == nil {
			return nil
		}
		return fmt.Errorf("Error : can not set parameters!")
	}).Struct()
	if err != nil {
		log.Printf("Error in TransferFunds : %v", err)
		return err
	}
	res, err := result.Res()
	log.Println(res)
	return err
}

func (r Request) Close() {
	r.conn.Close()
}
