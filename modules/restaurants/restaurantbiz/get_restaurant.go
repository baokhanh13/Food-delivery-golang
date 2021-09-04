package restaurantbiz

import (
	"context"
	"food-delivery/modules/restaurants/model"
)

type GetRestaurantStore interface {
	FindData(ctx context.Context, conditions map[string]interface{}) (*model.Restaurant, error)
}

type getRestaurantBiz struct {
	store GetRestaurantStore
}

func NewGetRestaurantBiz(store GetRestaurantStore) *getRestaurantBiz {
	return &getRestaurantBiz{store: store}
}

func (biz *getRestaurantBiz) GetRestaurant(ctx context.Context, id int) (*model.Restaurant, error) {
	result, err := biz.store.FindData(ctx, map[string]interface{}{"id": id})

	return result, err
}
