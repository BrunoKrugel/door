package db

import (
	"log"

	"github.com/arriqaaq/flashdb"
)

var mydb *flashdb.FlashDB

func Init() {
	config := &flashdb.Config{Path: "", EvictionInterval: 10}
	db, err := flashdb.New(config)
	if err != nil {
		log.Fatal(err)
	}
	mydb = db
}

func Get() *flashdb.FlashDB {
	return mydb
}

func Close() {
	mydb.Close()
}

func Update(key string, value string) error {
	err := mydb.Update(func(tx *flashdb.Tx) error {
		err := tx.Set(key, value)
		return err
	})
	return err
}

func Read(key string) (string, error) {
	var value string
	err := mydb.View(func(tx *flashdb.Tx) error {
		val, err := tx.Get(key)
		if err != nil {
			return err
		}
		value = val
		return nil
	})
	return value, err
}
