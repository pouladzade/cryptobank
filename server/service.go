package server

import (
	"github.com/cryptobank/acm"
	crb "github.com/cryptobank/cryptobank"
	"github.com/cryptobank/server/cdb"
	"log"
	"zombiezen.com/go/capnproto2"
)

type Service struct {
	c  crb.CoreBanking
	db *cdb.CryptoDb
}

func (c *Service) SetDb(db *cdb.CryptoDb) {
	c.db = db
}

func (c *Service) CreateAccount(call crb.CoreBanking_createAccount) error {

	res, _ := crb.NewResponse(call.Results.Segment())
	var acc acm.Account
	acc.Name, _ = call.Params.Name()
	acc.Balance = call.Params.Balance()
	acid, err := call.Params.AccountId()
	if err != nil {
		res.SetMessage(err.Error())
		res.SetCode(-1)
		return call.Results.SetRes(res)
	}
	acc.SetAccountId(acid)
	c.db.InsertAccount(acc)
	res.SetMessage("Succesfull")
	res.SetCode(0)
	c.db.Commit()
	return call.Results.SetRes(res)
}

func (c *Service) DeleteAccount(call crb.CoreBanking_deleteAccount) error {
	res, _ := crb.NewResponse(call.Results.Segment())
	var acc acm.Account
	acid, err := call.Params.AccountId()
	if err != nil {
		res.SetMessage(err.Error())
		res.SetCode(-2)
		return call.Results.SetRes(res)
	}
	acc.SetAccountId(acid)
	c.db.DeleteAccount(acc.AccountIdString())
	res.SetMessage("Succesfull")
	res.SetCode(0)
	c.db.Commit()
	return call.Results.SetRes(res)
}

func (c *Service) TransferFunds(call crb.CoreBanking_transferFunds) error {
	res, _ := crb.NewResponse(call.Results.Segment())
	var src, des acm.Account
	srcAdd, err := call.Params.Source()
	if err != nil {
		res.SetMessage(err.Error())
		res.SetCode(-3)
		return call.Results.SetRes(res)
	}
	src.SetAccountId(srcAdd)
	desAdd, err := call.Params.Destination()
	if err != nil {
		res.SetMessage(err.Error())
		res.SetCode(-4)
		return call.Results.SetRes(res)
	}
	des.SetAccountId(desAdd)
	src = c.db.LoadAccount(src.AccountIdString())
	des = c.db.LoadAccount(des.AccountIdString())
	mess, code := c.settle(src, des, call.Params.Amount())
	res.SetMessage(mess)
	res.SetCode(code)
	return call.Results.SetRes(res)
}

func (c *Service) createResponse(seg *capnp.Segment, code int32, message string) (crb.Response, error) {
	res, err := crb.NewResponse(seg)
	if err != nil {
		log.Fatal("Can not make the response!")
	}
	res.SetMessage(message)
	res.SetCode(code)
	return res, err
}

func (c *Service) settle(src, des acm.Account, amount uint64) (string, int32) {
	if des.Balance < amount {
		return "Insuficient balance", -5
	}
	src.Balance += amount
	des.Balance -= amount
	c.db.UpdateAccount(src)
	c.db.UpdateAccount(des)
	c.db.Commit()

	return "Successfull", 0
}
