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

func (r *queryResolver) SemuaBarang(ctx context.Context) ([]*model.Barang, error) {
	data := controller.NyariBarangDiDb()

	return data, nil
}

func (r *queryResolver) SemuaUser(ctx context.Context) ([]*model.User, error) {
	data := controller.NyariUserDiDb()

	return data, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
