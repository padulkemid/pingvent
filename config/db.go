package config

import (
	"log"

	"github.com/go-pg/pg"
	controller "github.com/padulkemid/pingpos/controllers"
)

func Connection() *pg.DB {
	option := &pg.Options{
		User:     "padulkemid",
		Password: "",
		Addr:     "localhost:5432",
		Database: "pingpos",
	}

	var db *pg.DB = pg.Connect(option)

	if db == nil {
		log.Printf("DB not exist")
	}

	log.Printf("db is connected")
	controller.BuatTableBarang(db)
	controller.BuatTableUser(db)
	controller.InitiateDB(db)

	return db
}
