package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
  "github.com/myk4040okothogodo/Microservices4Jumia/ServiceB_Jumia/internal/orders"
	"github.com/myk4040okothogodo/Microservices4Jumia/ServiceB_Jumia/graph/generated"
	"github.com/myk4040okothogodo/Microservices4Jumia/ServiceB_Jumia/graph/model"
)

func (r *mutationResolver) CreateOrder(ctx context.Context, input model.NewOrder) (*model.Order, error) {
	var order orders.Order
  order.Parcelweight = input.Parcelweight
  order.Country =  input.Country
  order.Email  =   input.Email
  order.Phone  =   input.Phone
  orderId := order.Save()
  
  return &model.Order{ID: strconv.FormatInt(orderId, 10), Parcelweight: order.Parcelweight, Country: order.Country, Email: order.Email, Phone: order.Phone}, nil
}

func (r *mutationResolver) CreateUser(ctx context.Context, input model.NewUser) (string, error) {
	var user users.User
  user.Username = input.Username
  user.Password = input.Password
  user.Create()
  return &model.User{Username: user.Username}, nil
}

func (r *mutationResolver) Login(ctx context.Context, input model.Login) (string, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) RefreshToken(ctx context.Context, input model.RefreshTokenInput) (string, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Orders(ctx context.Context) ([]*model.Order, error) {
	var resultOrders  []*model.Order
  var dbOrders  []orders.Order
  dbOrders = orders.GetAll()
  for _, order := range dbOrders {
      resultOrders = append(resultOrders, &model.Order{ID: order.ID, Parcelweight: order.Parcelweight,Country: order.Country, Email: order.Email, Phone: order.Phone })
  }
  return resultOrders, nil
}

func (r *queryResolver) Users(ctx context.Context) ([]*model.User, error) {
	panic(fmt.Errorf("not implemented"))
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
