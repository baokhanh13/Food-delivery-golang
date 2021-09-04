package restaurantbiz

import (
	"context"
	"food-delivery/modules/restaurants/model"
)

type CreateRestaurantStore interface {
	Create(ctx context.Context, data *model.RestaurantCreate) error
}

type createRestaurantBiz struct {
	store CreateRestaurantStore
}

func NewCreateRestaurantBiz(store CreateRestaurantStore) *createRestaurantBiz {
	return &createRestaurantBiz{store: store}
}

func (biz *createRestaurantBiz) CreateRestaurant(ctx context.Context, data *model.RestaurantCreate) error {
	err := biz.store.Create(ctx, data)
	return err
}
