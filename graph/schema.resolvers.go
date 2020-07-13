package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	guuid "github.com/google/uuid"
	controller "github.com/padulkemid/pingpos/controllers"
	"github.com/padulkemid/pingpos/graph/generated"
	"github.com/padulkemid/pingpos/graph/model"
	"github.com/padulkemid/pingpos/utils"
)

func (r *mutationResolver) BuatBarang(ctx context.Context, input model.BarangBaru) (*model.Barang, error) {
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

func (r *mutationResolver) EditBarang(ctx context.Context, id string, input model.BarangBaru) (*model.Barang, error) {
	barang := &model.Barang{
		Nama:      input.Nama,
		Harga:     input.Harga,
		Stock:     input.Stock,
		Vendor:    input.Vendor,
		UpdatedAt: utils.JamWaktu(),
	}

	data := controller.EditBarang(id, barang)

	return data, nil
}

func (r *mutationResolver) HapusBarang(ctx context.Context, id string) (bool, error) {
	data := controller.DeleteBarang(id)

	return data, nil
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

func (r *mutationResolver) EditUser(ctx context.Context, id string, input model.UserBaru) (*model.User, error) {
	hashed, _ := utils.HashPassword(input.Password)

	user := &model.User{
		Username: input.Username,
		Password: hashed,
		Role:     input.Role,
	}

	data := controller.EditUser(id, user)

	return data, nil
}

func (r *mutationResolver) HapusUser(ctx context.Context, id string) (bool, error) {
	data := controller.DeleteUser(id)

	return data, nil
}

func (r *mutationResolver) LoginUser(ctx context.Context, input model.LoginUser) (string, error) {
  data, check := controller.AuthUser(input.Username, input.Password)
  if !check {
    panic(check)
  }

  token, err := utils.GenerateToken(data.ID, data.Username)
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

  token, err := utils.GenerateToken(data.ID, data.Username)
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
	data := controller.NyariUserDiDb()

	return data, nil
}

func (r *queryResolver) UserPakeID(ctx context.Context, id string) (*model.User, error) {
	data := controller.NyariUserPakeId(id)

	return data, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
