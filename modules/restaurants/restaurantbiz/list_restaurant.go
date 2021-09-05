package restaurantbiz

import (
	"context"
	"food-delivery/common"
	"food-delivery/modules/restaurants/model"
)

type ListRestaurantStore interface {
	ListData(ctx context.Context, paging *common.Paging, filter *model.Filter, moreKeys ...string) ([]model.Restaurant, error)
}

type listRestaurantBiz struct {
	store ListRestaurantStore
}

func NewListRestaurantBiz(store ListRestaurantStore) *listRestaurantBiz {
	return &listRestaurantBiz{store: store}
}

func (biz *listRestaurantBiz) ListRestaurant(ctx context.Context, paging *common.Paging, filter *model.Filter) ([]model.Restaurant, error) {
	result, err := biz.store.ListData(ctx, paging, filter)
	return result, err
}
