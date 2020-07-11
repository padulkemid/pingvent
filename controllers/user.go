package controller

import (
	"log"

	"github.com/padulkemid/pingpos/graph/model"
)

// User
func BuatUserKeDb(user *model.User) error {
	err := dbConnect.Insert(user)

	if err != nil {
		panic(err)
	}

	log.Printf("User dah masuk boi!")

	return nil
}

func NyariUserDiDb() []*model.User {
	var user []*model.User
	err := dbConnect.Model(&user).Select()

	if err != nil {
		panic(err)
	}

	return user
}
