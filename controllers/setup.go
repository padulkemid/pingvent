package controller

import (
	"log"

	"github.com/go-pg/pg"
	"github.com/go-pg/pg/orm"
	"github.com/padulkemid/pingpos/graph/model"
)

// table

func BuatTableBarang(db *pg.DB) error {

	options := &orm.CreateTableOptions{
		IfNotExists: true,
	}

	err := db.CreateTable(&model.Barang{}, options)

	if err != nil {
		panic(err)
	}

	log.Printf("Table barang udah dibikin coy!")

	return nil

}

func BuatTableUser(db *pg.DB) error {

	options := &orm.CreateTableOptions{
		IfNotExists: true,
	}

	err := db.CreateTable(&model.User{}, options)

	if err != nil {
		panic(err)
	}

	log.Printf("Table user juga sekuyyyy!")

	return nil

}

// initiate
var dbConnect *pg.DB

func InitiateDB(db *pg.DB) {
	dbConnect = db
}
