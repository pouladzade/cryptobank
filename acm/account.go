package acm

import (
	"encoding/hex"
	"fmt"
	crp "github.com/cryptobank/cryptobank"
	"github.com/monax/bosmarmot/monax/log"
)

type Address []byte
type Account struct {
	accId   Address `json:"accountId"`
	Name    string  `json:"name"`
	Balance uint64  `json:"balance"`
}

func (ac *Account) SetAccountIdString(str string) {
	b, err := hex.DecodeString(str)
	if len(b) == int(crp.AccountSize) && err == nil {
		ac.accId = make(Address, crp.AccountSize)
		copy(ac.accId, b)
	} else {
		log.Warn(fmt.Sprintf("[%s] Can not convert to byte\n Error=[%s]"), str, err)
	}
}

func (ac *Account) SetAccountId(accid Address) {
	if len(accid) == int(crp.AccountSize) {
		ac.accId = make(Address, crp.AccountSize)
		copy(ac.accId, accid)
	} else {
		log.Warn(fmt.Sprintf("[%v] Account Size should be 32 bytes "), accid)
	}
}

func (ac *Account) AccountId() []byte {
	return ac.accId
}

func (ac *Account) AccountIdString() string {
	return hex.EncodeToString(ac.accId)
}
