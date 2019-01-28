package cdb

import (
	"encoding/json"
	"fmt"
	"github.com/cryptobank/acm"
	"github.com/monax/bosmarmot/monax/log"
	"io/ioutil"
	"os"
)

const (
	CryptoDB_Name = "cryptodb.json"
	CryptoDB_Dir  = "./db"
)

type CryptoDb struct{ db map[string]acm.Account }

func (cdb *CryptoDb) InsertAccount(acc acm.Account) error {
	temp := cdb.LoadAccount(acc.AccountIdString())
	if temp.AccountIdString() != "" {
		return fmt.Errorf("Account already exists in database!")
	}
	cdb.db[acc.AccountIdString()] = acc
	return nil
}

func (cdb *CryptoDb) DeleteAccount(accid string) error {
	acc := cdb.LoadAccount(accid)
	if acc.AccountIdString() == "" {
		return fmt.Errorf("Can not find the Account in database!")
	}
	delete(cdb.db, accid)
	return nil
}

func (cdb *CryptoDb) UpdateAccount(acc acm.Account) {
	cdb.db[acc.AccountIdString()] = acc
}

func (cdb *CryptoDb) LoadAccount(accid string) acm.Account {
	return cdb.db[accid]
}

func (cdb *CryptoDb) Commit() {
	jsonString, _ := json.Marshal(cdb.db)
	fmt.Println(string(jsonString))
	var f *os.File
	if _, err := os.Stat(CryptoDB_Dir + "/" + CryptoDB_Name); os.IsNotExist(err) {
		os.Mkdir(CryptoDB_Dir, 0700)
		f, err = os.Create(CryptoDB_Dir + "/" + CryptoDB_Name)
		f.WriteString(string(jsonString))
		f.Close()
	} else {
		ioutil.WriteFile(CryptoDB_Dir+"/"+CryptoDB_Name, jsonString, 0700)
	}
}

func (cdb *CryptoDb) LoadDb() {
	cdb.db = make(map[string]acm.Account)
	dat, err := ioutil.ReadFile(CryptoDB_Dir + "/" + CryptoDB_Name)
	if err != nil {
		log.Warn(fmt.Sprintf("Can not find cdb file : [%s]", CryptoDB_Dir+"/"+CryptoDB_Name))
	} else {
		if json.Unmarshal(dat, cdb.db) != nil {
			log.Warn(err)
		}
	}
}
