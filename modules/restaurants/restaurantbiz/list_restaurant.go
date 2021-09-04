package restaurantbiz

import (
	"context"
	"food-delivery/modules/restaurants/model"
)

type ListRestaurantStore interface {
	ListData(ctx context.Context) ([]model.Restaurant, error)
}

type listRestaurantBiz struct {
	store ListRestaurantStore
}

func NewListRestaurantBiz(store ListRestaurantStore) *listRestaurantBiz {
	return &listRestaurantBiz{store: store}
}

func (biz *listRestaurantBiz) ListRestaurant(ctx context.Context) ([]model.Restaurant, error) {
	result, err := biz.store.ListData(ctx)
	return result, err
}
