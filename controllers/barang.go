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

// Barang
func BuatBarangKeDb(barang *model.Barang) error {
	err := dbConnect.Insert(barang)

	if err != nil {
		panic(err)
	}

	log.Printf("Barang dah masuk boi!")

	return nil
}

func NyariBarangDiDb() []*model.Barang {
	var barang []*model.Barang
	err := dbConnect.Model(&barang).Select()

	if err != nil {
		panic(err)
	}

	log.Printf("Noh gua kasi semua barang!")

	return barang
}

func NyariBarangPakeId(id string) *model.Barang {
	barang := &model.Barang{ID: id}

	err := dbConnect.Select(barang)

	if err != nil {
		panic(err)
	}

	log.Printf("Nih lu minta barang dari ID: %s", id)

	return barang
}

func EditBarang(id string, barang *model.Barang) *model.Barang {
	barangLama := NyariBarangPakeId(id)

	editedBarang := &model.Barang{
		ID:        id,
		Nama:      barang.Nama,
		Harga:     barang.Harga,
		Stock:     barang.Stock,
		Vendor:    barang.Vendor,
		CreatedAt: barangLama.CreatedAt,
		UpdatedAt: barang.UpdatedAt,
	}

	err := dbConnect.Update(editedBarang)

	if err != nil {
		panic(err)
	}

	log.Printf("Barang udah diapdet")

	return editedBarang
}

func DeleteBarang(id string) bool {
	barang := &model.Barang{ID: id}

	err := dbConnect.Delete(barang)

	if err != nil {
		panic(err)
	}

	log.Printf("Barang udah diapus")

	return true
}
