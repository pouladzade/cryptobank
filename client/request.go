package client

import (
	"context"
	"fmt"
	"github.com/cryptobank/acm"
	crb "github.com/cryptobank/cryptobank"
	"github.com/monax/bosmarmot/monax/log"
	"zombiezen.com/go/capnproto2/rpc"
)

type Request struct {
	cb   crb.CoreBanking
	ctx  context.Context
	conn *rpc.Conn
}

func (r Request) CreateAccount(acc acm.Account) error {
	fmt.Printf("Calling CreateAccount:\n AccountId = %s , name = %s , balance = %d\n\n",
		acc.AccountIdString(), acc.Name, acc.Balance)
	result, err := r.cb.CreateAccount(r.ctx, func(p crb.CoreBanking_createAccount_Params) error {
		p.SetBalance(acc.Balance)
		if p.SetAccountId(acc.AccountId()) == nil && p.SetName(acc.Name) == nil {
			return nil
		}
		return fmt.Errorf("Error : can not set parameters!")
	}).Struct()
	if err != nil {
		log.Warn("Error in CreateAccount : %s", err.Error())
		return err
	}
	res, err := result.Res()
	if err == nil {
		r.logResult(&res)
	}

	return err
}

func (r Request) DeleteAccount(acc acm.Account) error {
	fmt.Printf("Calling DeleteAccount: AccountId = %v \n\n", acc)
	result, err := r.cb.DeleteAccount(r.ctx, func(p crb.CoreBanking_deleteAccount_Params) error {
		err := p.SetAccountId(acc.AccountId())
		return err
	}).Struct()
	if err != nil {
		log.Warn("Error in DeleteAccount : %s", err.Error())
		return err
	}
	res, err := result.Res()
	if err == nil {
		r.logResult(&res)
	}
	return err
}

func (r Request) TransferFunds(src, des acm.Account, amount uint64) error {
	fmt.Printf("Calling TransferFunds: source = %v , destination = %v , amount = %d\n\n", src, des, amount)
	result, err := r.cb.TransferFunds(r.ctx, func(p crb.CoreBanking_transferFunds_Params) error {
		p.SetAmount(amount)
		if p.SetSource(src.AccountId()) == nil && p.SetDestination(des.AccountId()) == nil {
			return nil
		}
		return fmt.Errorf("Error : can not set parameters!")
	}).Struct()
	if err != nil {
		log.Warn("Error in TransferFunds : %s", err.Error())
		return err
	}
	res, err := result.Res()
	if err == nil {
		r.logResult(&res)
	}
	return err
}

func (r Request) logResult(res *crb.Response) {
	message, _ := res.Message()
	fmt.Printf("Response : Code = %d , Message = %s\n", res.Code(), message)
}

func (r Request) Close() {
	r.conn.Close()
}
