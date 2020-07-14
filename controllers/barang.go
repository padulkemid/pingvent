package controller

import (
	"fmt"
	"log"

	"github.com/padulkemid/pingpos/graph/model"
)

// Barang
func BuatBarangKeDb(barang *model.Barang) error {
	err := dbConnect.Insert(barang)

	if err != nil {
    return fmt.Errorf("Barang sudah ada boi, ganti la")
	}

	log.Printf("Barang dah masuk boi!")

	return nil
}

func NyariBarangDiDb() []*model.Barang {
	var barang []*model.Barang
	err := dbConnect.Model(&barang).Order("id ASC").Select()

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

func EditBarang(id string, barang *model.Barang) (*model.Barang, error) {
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
    return &model.Barang{}, fmt.Errorf("Ga ada barangnya...")
	}

	log.Printf("Barang udah diapdet")

	return editedBarang, nil
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
