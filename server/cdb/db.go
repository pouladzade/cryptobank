package cdb

import (
	"fmt"
	"github.com/cryptobank/acm"
	"github.com/syndtr/goleveldb/leveldb"
)

const (
	CryptoDB_Name = "cryptodb"
	CryptoDB_Dir  = "./db"
)

type CryptoDb struct{ db *leveldb.DB }

func (cdb *CryptoDb) InsertAccount(acc acm.Account) error {
	temp, _ := cdb.LoadAccount(acc.AccountId())
	if temp != nil {
		return fmt.Errorf("Account already exists in database!")
	}
	return cdb.UpdateAccount(acc)
}

func (cdb *CryptoDb) DeleteAccount(acc acm.Account) error {
	return cdb.db.Delete([]byte(acc.AccountId()), nil)
}

func (cdb *CryptoDb) UpdateAccount(acc acm.Account) error {
	cdb.DeleteAccount(acc)
	bs, _ := acc.Encode()
	return cdb.db.Put(acc.AccountId(), bs, nil)
}

func (cdb *CryptoDb) LoadAccount(accid []byte) (*acm.Account, error) {
	bs, err := cdb.db.Get(accid, nil)

	if err != nil {
		return nil, err
	}
	ac := new(acm.Account)
	err = ac.Decode(bs, accid)
	return ac, err
}

func (cdb *CryptoDb) Commit() error {
	return nil
}

func (cdb *CryptoDb) LoadDb() error {
	var err error
	cdb.db, err = leveldb.OpenFile(CryptoDB_Dir, nil)
	return err
}
