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
	defer log.Println(res)
	var acc acm.Account
	acc.Name, _ = call.Params.Name()
	acc.Balance = call.Params.Balance()
	acid, err := call.Params.AccountId()
	log.Printf("\n{request: CreateAccount\nparams[accountId = %v, name=%s, balance=%d]}\n",
		acid, acc.Name, acc.Balance)
	if err != nil {
		res.SetMessage(err.Error())
		res.SetCode(-1)
		return call.Results.SetRes(res)
	}
	acc.SetAccountId(acid)
	err = c.db.InsertAccount(acc)
	if err != nil {
		res.SetMessage(err.Error())
		res.SetCode(-10)
		return call.Results.SetRes(res)
	}
	res.SetMessage("Account succesfully created!")
	res.SetCode(0)
	c.db.Commit()
	return call.Results.SetRes(res)
}

func (c *Service) DeleteAccount(call crb.CoreBanking_deleteAccount) error {
	res, _ := crb.NewResponse(call.Results.Segment())
	defer log.Println(res)
	var acc acm.Account
	acid, err := call.Params.AccountId()
	log.Printf("\n{request: DeleteAccount\nparams[accountId = %v]}\n",
		acid)
	if err != nil {
		res.SetMessage(err.Error())
		res.SetCode(-2)
		return call.Results.SetRes(res)
	}
	acc.SetAccountId(acid)
	err = c.db.DeleteAccount(acc)
	if err != nil {
		res.SetMessage(err.Error())
		res.SetCode(-3)
		return call.Results.SetRes(res)
	}
	res.SetMessage("Account succesfully deleted!")
	res.SetCode(0)
	c.db.Commit()
	return call.Results.SetRes(res)
}

func (c *Service) TransferFunds(call crb.CoreBanking_transferFunds) error {
	res, _ := crb.NewResponse(call.Results.Segment())
	defer log.Println(res)
	srcAcc, err := call.Params.Source()
	log.Printf("\n{request: TransferFunds")
	if err != nil {
		res.SetMessage(err.Error())
		res.SetCode(-4)
		return call.Results.SetRes(res)
	}
	desAcc, err := call.Params.Destination()
	if err != nil {
		res.SetMessage(err.Error())
		res.SetCode(-5)
		return call.Results.SetRes(res)
	}
	log.Printf("\n{params[Source = %v, Destination=%v, Amount=%d]}\n",
		srcAcc, desAcc, call.Params.Amount())

	src, err := c.db.LoadAccount(srcAcc)
	if err != nil {
		res.SetMessage("Can not find the source account!")
		res.SetCode(-6)
		return call.Results.SetRes(res)
	}
	des, err := c.db.LoadAccount(desAcc)
	if err != nil {

		res.SetMessage("Can not find the destination account!")
		res.SetCode(-7)
		return call.Results.SetRes(res)
	}
	mess, code := c.settle(*src, *des, call.Params.Amount())
	res.SetMessage(mess)
	res.SetCode(code)
	return call.Results.SetRes(res)
}

func (c *Service) GetAccountInfo(call crb.CoreBanking_getAccountInfo) error {
	res, _ := crb.NewResponse(call.Results.Segment())
	defer log.Println(res)
	acid, err := call.Params.AccountId()
	log.Printf("\n{request: GetAccountInfo\nparams[accountId = %v]}\n",
		acid)
	if err != nil {
		res.SetMessage(err.Error())
		res.SetCode(-1)
		return call.Results.SetRes(res)
	}

	ac, err := c.db.LoadAccount(acid)
	if err != nil || ac == nil {
		res.SetMessage(err.Error())
		res.SetCode(-10)
		return call.Results.SetRes(res)
	}
	res.SetMessage("GetAccountInfo Succesfully done!")
	res.SetCode(0)
	call.Results.SetBalance(ac.Balance)
	call.Results.SetName(ac.Name)
	c.db.Commit()
	return call.Results.SetRes(res)
}

func (c *Service) createResponse(seg *capnp.Segment, code int32, message string) (crb.Response, error) {
	res, err := crb.NewResponse(seg)
	if err != nil {
		log.Println("Can not make the response!")
	}
	res.SetMessage(message)
	res.SetCode(code)
	return res, err
}

func (c *Service) settle(src, des acm.Account, amount uint64) (string, int32) {
	if src.Balance < amount {
		log.Println("Insuficient balance")
		return "Insuficient balance", -8
	}
	src.Balance = src.Balance - amount
	des.Balance = des.Balance + amount
	if c.db.UpdateAccount(src) == nil {
		if c.db.UpdateAccount(des) != nil {
			//Todo Need to implement some reverse plan for source account
			return "can not update account!", 15
		}
	}

	c.db.Commit()
	log.Println("Amount Successfully transffered")
	return "Amount Successfully transffered", 0
}
