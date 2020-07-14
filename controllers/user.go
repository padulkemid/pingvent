package controller

import (
	"fmt"
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

func NyariUserPakeId(id string) (*model.User, error) {
	user := &model.User{ID: id}

	err := dbConnect.Select(user)

	if err != nil {
    return &model.User{}, fmt.Errorf("User ga ada bos!")
	}

	log.Printf("Nih lu minta user dari ID: %s", id)

	return user, nil
}

func EditUser(id string, newUser *model.EditUser) (*model.User, error) {
  oldUser, err := NyariUserPakeId(id)
  if err != nil {
    return &model.User{}, err
  }

  checked := utils.CheckPassword(newUser.PasswordLama, oldUser.Password)

  if !checked {
    return &model.User{}, fmt.Errorf("Password lama salah brok!")
  }

  hashed, _ := utils.HashPassword(newUser.PasswordBaru)

  editedUser := &model.User{
    ID:       id,
    Username: newUser.Username,
    Password: hashed,
    Role:     oldUser.Role,
  }

	findErr := dbConnect.Update(editedUser)

	if findErr != nil {
    return &model.User{}, fmt.Errorf("User ga ada boy!")
	}

	log.Printf("User udah diedit")

	return editedUser, nil
}

func DeleteUser(id string) (bool, error) {

	user := &model.User{ID: id}

	err := dbConnect.Delete(user)

	if err != nil {
    return false, fmt.Errorf("Ga ada brok, cari yg laen")
	}

	log.Printf("User id :%s , udah diapus", id)

	return true, nil
}

// login
type LoginData struct {
	Role       string `json:"id"`
	Username string `json:"username"`
}

func UsernameAdaGak(username string) (*LoginData, error)  {
  var user model.User
  err := dbConnect.Model(&user).Where("username=?", username).Select()

	if err != nil {
    panic(err)
	}

  data := &LoginData{
    Role: user.Role,
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
    Role: user.Role,
    Username: user.Username,
  }

  return data, check
}
