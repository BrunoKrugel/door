package db

import (
	"strconv"

	"github.com/hashicorp/go-memdb"
)

type Card struct {
	ID int
}

var mydb *memdb.MemDB

func Init() {
	if mydb != nil {
		return
	}

	// Create the DB schema
	schema := &memdb.DBSchema{
		Tables: map[string]*memdb.TableSchema{
			"card": {
				Name: "card",
				Indexes: map[string]*memdb.IndexSchema{
					"id": {
						Name:    "id",
						Unique:  true,
						Indexer: &memdb.StringFieldIndex{Field: "ID"},
					},
				},
			},
		},
	}
	// Create a new data base
	db, err := memdb.NewMemDB(schema)
	if err != nil {
		panic(err)
	}
	mydb = db
}

func GetDb() *memdb.MemDB {
	return mydb
}

func Write(index int) error {
	db := mydb
	txn := db.Txn(true)

	card := []*Card{
		{index},
	}
	for _, p := range card {
		if err := txn.Insert("card", p); err != nil {
			return err
		}
	}
	txn.Commit()
	return nil
}

func Read(index int) error {
	db := mydb
	txn := db.Txn(false)
	defer txn.Abort()

	raw, err := txn.First("card", "id", strconv.Itoa(index))
	if err != nil {
		return err
	}
	if raw == nil {
		return nil
	}
	return nil
}
