package client

import (
	"context"
	"fmt"
	crb "github.com/cryptobank/cryptobank"
	"github.com/monax/bosmarmot/monax/log"
	"zombiezen.com/go/capnproto2/rpc"
)

type Request struct {
	cb  crb.CoreBanking
	ctx context.Context
	conn *rpc.Conn
}

func (r Request) CreateAccount(acc []byte, name string, bal uint64) error {
	fmt.Printf("Calling CreateAccount:\n AccountId = %v , Name = %s , Balance = %d\n\n", acc, name, bal)
	result, err := r.cb.CreateAccount(r.ctx, func(p crb.CoreBanking_createAccount_Params) error {
		p.SetBalance(bal)
		if p.SetAccountId(acc) == nil && p.SetName(name) == nil {
			return nil
		}
		return fmt.Errorf("Error : can not set parameters!")
	}).Struct()
	if err != nil {
		log.Warn("Error in CreateAccount : %v", err)
		return err
	}
	res, err := result.Res()
	r.logResult(res)
	return err
}

func (r Request) DeleteAccount(acc []byte) error {
	fmt.Printf("Calling DeleteAccount: AccountId = %v \n\n", acc)
	result, err := r.cb.DeleteAccount(r.ctx, func(p crb.CoreBanking_deleteAccount_Params) error {
		err := p.SetAccountId(acc)
		return err
	}).Struct()
	if err != nil {
		log.Warn("Error in DeleteAccount : %v", err)
		return err
	}
	res, err := result.Res()
	r.logResult(res)
	return err
}

func (r Request) TransferFunds(src []byte, des []byte, amount uint64) error {
	fmt.Printf("Calling TransferFunds: source = %v , destination = %v , amount = %d\n\n", src, des, amount)
	result, err := r.cb.TransferFunds(r.ctx, func(p crb.CoreBanking_transferFunds_Params) error {
		p.SetAmount(amount)
		if p.SetSource(src) == nil && p.SetDestination(des) == nil {
			return nil
		}
		return fmt.Errorf("Error : can not set parameters!")
	}).Struct()
	if err != nil {
		log.Warn("Error in TransferFunds : %v", err)
		return err
	}
	res, err := result.Res()
	r.logResult(res)
	return err
}

func (r Request) logResult(res crb.Response) {
	message, _ := res.Message()
	fmt.Printf("Response : Code = %d , Message = %s\n", res.Code(), message)
}

func (r Request) Close(){
	r.conn.Close()
}