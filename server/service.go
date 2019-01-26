package server

import (
	crb "github.com/cryptobank/cryptobank"
	"log"
	"zombiezen.com/go/capnproto2"
)

// coreBanking is a local implementation of CoreBanking.
type Service struct{ c crb.CoreBanking }

func (c Service) CreateAccount(call crb.CoreBanking_createAccount) error {

	res, _ := crb.NewResponse(call.Results.Segment())
	res.SetMessage("salam")
	res.SetCode(123)
	return call.Results.SetRes(res)
}

func (c Service) DeleteAccount(call crb.CoreBanking_deleteAccount) error {
	res, _ := crb.NewResponse(call.Results.Segment())
	res.SetMessage("salam")
	res.SetCode(123)
	return call.Results.SetRes(res)
}

func (c Service) TransferFunds(call crb.CoreBanking_transferFunds) error {
	res, _ := crb.NewResponse(call.Results.Segment())
	res.SetMessage("salam")
	res.SetCode(123)
	return call.Results.SetRes(res)
}

func (c Service) createResponse(seg *capnp.Segment, code int32, message string) (crb.Response, error) {
	res, err := crb.NewResponse(seg)
	if err != nil {
		log.Fatal("Can not make the response!")
	}
	res.SetMessage(message)
	res.SetCode(code)
	return res, err
}
