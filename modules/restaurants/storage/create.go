package storage

import (
	"context"
	"food-delivery/modules/restaurants/model"
)

func (s *sqlStore) Create(context context.Context, data *model.RestaurantCreate) error {
	db := s.db

	if err := db.Create(data).Error; err != nil {
		return err
	}
	return nil
}
