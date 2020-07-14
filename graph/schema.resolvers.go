package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	guuid "github.com/google/uuid"
	"github.com/padulkemid/pingpos/auth"
	controller "github.com/padulkemid/pingpos/controllers"
	"github.com/padulkemid/pingpos/graph/generated"
	"github.com/padulkemid/pingpos/graph/model"
	"github.com/padulkemid/pingpos/utils"
)

func (r *mutationResolver) BuatBarang(ctx context.Context, input model.BarangBaru) (*model.Barang, error) {
	user, ok := auth.ForContext(ctx)

	if !ok {
		return &model.Barang{}, fmt.Errorf("Kamu bukan siapa-siapa!")
	} else {
		if user.Role == "user" {
			return &model.Barang{}, fmt.Errorf("Kamu bukan admin / penjual!")
		} else {
			barang := &model.Barang{
				ID:        guuid.New().String(),
				Nama:      input.Nama,
				Harga:     input.Harga,
				Stock:     input.Stock,
				Vendor:    input.Vendor,
				CreatedAt: utils.JamWaktu(),
				UpdatedAt: utils.JamWaktu(),
			}

			err := controller.BuatBarangKeDb(barang)

			if err != nil {
				panic(err)
			}

			return barang, nil
		}
	}
}

func (r *mutationResolver) EditBarang(ctx context.Context, id string, input model.BarangBaru) (*model.Barang, error) {
	user, ok := auth.ForContext(ctx)

	if !ok {
		return &model.Barang{}, fmt.Errorf("Kamu bukan siapa-siapa!")
	} else {
		if user.Role == "user" {
			return &model.Barang{}, fmt.Errorf("Kamu bukan admin / penjual!")
		} else {
			barang := &model.Barang{
				Nama:      input.Nama,
				Harga:     input.Harga,
				Stock:     input.Stock,
				Vendor:    input.Vendor,
				UpdatedAt: utils.JamWaktu(),
			}

			data, err := controller.EditBarang(id, barang)

			if err != nil {
				return &model.Barang{}, err
			}

			return data, nil

		}
	}
}

func (r *mutationResolver) HapusBarang(ctx context.Context, id string) (bool, error) {
	user, ok := auth.ForContext(ctx)

	if !ok {
		return false, fmt.Errorf("Kamu bukan siapa-siapa!")
	} else {
		if user.Role == "user" {
			return false, fmt.Errorf("Kamu bukan admin / penjual!")
		} else {
			data := controller.DeleteBarang(id)

			return data, nil

		}
	}
}

func (r *mutationResolver) BuatUser(ctx context.Context, input model.UserBaru) (*model.User, error) {
	hashed, _ := utils.HashPassword(input.Password)

	user := &model.User{
		ID:       guuid.New().String(),
		Username: input.Username,
		Password: hashed,
		Role:     input.Role,
	}

	err := controller.BuatUserKeDb(user)

	if err != nil {
		panic(err)
	}

	return user, nil
}

func (r *mutationResolver) EditUser(ctx context.Context, id string, input model.EditUser) (*model.User, error) {
	user, ok := auth.ForContext(ctx)

	if !ok {
		return &model.User{}, fmt.Errorf("Kamu bukan siapa-siapa!")
	} else {
		if user.Role == "seller" {
			return &model.User{}, fmt.Errorf("Kamu bukan admin / user!")
		} else {
      editedUser := model.EditUser{
        Username: input.Username,
        PasswordLama: input.PasswordLama,
        PasswordBaru: input.PasswordBaru,
      }

      data, err := controller.EditUser(id, &editedUser)

      if err != nil {
        return &model.User{}, err
      }

      return data, nil
		}
	}
}

func (r *mutationResolver) HapusUser(ctx context.Context, id string) (bool, error) {
	user, ok := auth.ForContext(ctx)

	if !ok {
		return false, fmt.Errorf("Kamu bukan siapa-siapa!")
	} else {
		if user.Role != "admin" {
			return false, fmt.Errorf("Kamu bukan admin!")
		} else {
			data, err := controller.DeleteUser(id)

			if err != nil {
				return false, err
			}

			return data, nil

		}
	}
}

func (r *mutationResolver) LoginUser(ctx context.Context, input model.LoginUser) (string, error) {
	data, check := controller.AuthUser(input.Username, input.Password)
	if !check {
		panic(check)
	}

	token, err := utils.GenerateToken(data.Role, data.Username)
	if err != nil {
		panic(err)
	}

	return token, nil
}

func (r *mutationResolver) RefreshToken(ctx context.Context, input model.RefreshTokenData) (string, error) {
	data, err := utils.ParseToken(input.Token)
	if err != nil {
		panic(err)
	}

	token, err := utils.GenerateToken(data.Role, data.Username)
	if err != nil {
		panic(err)
	}

	return token, nil
}

func (r *queryResolver) SemuaBarang(ctx context.Context) ([]*model.Barang, error) {
	data := controller.NyariBarangDiDb()

	return data, nil
}

func (r *queryResolver) BarangPakeID(ctx context.Context, id string) (*model.Barang, error) {
	data := controller.NyariBarangPakeId(id)

	return data, nil
}

func (r *queryResolver) SemuaUser(ctx context.Context) ([]*model.User, error) {
	user, ok := auth.ForContext(ctx)

	if !ok {
		return []*model.User{}, fmt.Errorf("Kamu bukan siapa-siapa!")
	} else {
		if user.Role != "admin" {
			return []*model.User{}, fmt.Errorf("Kamu bukan admin!")
		} else {
			data := controller.NyariUserDiDb()

			return data, nil
		}
	}
}

func (r *queryResolver) UserPakeID(ctx context.Context, id string) (*model.User, error) {
	user, ok := auth.ForContext(ctx)

	if !ok {
		return &model.User{}, fmt.Errorf("Kamu bukan siapa-siapa!")
	} else {
		if user.Role != "admin" {
			return &model.User{}, fmt.Errorf("Kamu bukan admin!")
		} else {
			data, err := controller.NyariUserPakeId(id)
			if err != nil {
				return &model.User{}, err
			}

			return data, nil
		}
	}
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
