package controller

import (
	"log"

	"github.com/padulkemid/pingpos/graph/model"
	utils "github.com/padulkemid/pingpos/utils"
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
	err := dbConnect.Model(&user).Order("id ASC").Select()

	if err != nil {
		panic(err)
	}

	log.Printf("Noh semua user")

	return user
}

func NyariUserPakeId(id string) *model.User {
	user := &model.User{ID: id}

	err := dbConnect.Select(user)

	if err != nil {
		panic(err)
	}

	log.Printf("Nih lu minta user dari ID: %s", id)

	return user
}

func EditUser(id string, user *model.User) *model.User {

	editedUser := &model.User{
		ID:       id,
		Username: user.Username,
		Password: user.Password,
		Role:     user.Role,
	}

	err := dbConnect.Update(editedUser)

	if err != nil {
		panic(err)
	}

	log.Printf("User udah diedit")

	return editedUser
}

func DeleteUser(id string) bool {

	user := &model.User{ID: id}

	err := dbConnect.Delete(user)

	if err != nil {
		panic(err)
	}

	log.Printf("User id :%s , udah diapus", id)

	return true
}

// login
type LoginData struct {
	ID       string `json:"id"`
	Username string `json:"username"`
}

func UsernameAdaGak(username string) (*LoginData, error)  {
  user := &model.User{
    Username: username,
  }

	err := dbConnect.Select(user)

	if err != nil {
    panic(err)
	}

  data := &LoginData{
    ID: user.ID,
    Username: user.Username,
  }

	return data, nil
}

func AuthUser(username, password string) (*LoginData, bool) {
  var user model.User
  err := dbConnect.Model(&user).Where("username=?", username).Select()

	if err != nil {
    panic(err)
	}

  check := utils.CheckPassword(password, user.Password)
  data := &LoginData{
    ID: user.ID,
    Username: user.Username,
  }

  return data, check
}
