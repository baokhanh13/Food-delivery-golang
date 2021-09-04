package storage

import (
	"context"
	"food-delivery/modules/restaurants/model"
)

func (s *sqlStore) FindData(ctx context.Context, conditions map[string]interface{}) (*model.Restaurant, error) {
	db := s.db

	var data model.Restaurant

	if err := db.Where(conditions).First(&data).Error; err != nil {
		return nil, err
	}
	return &data, nil
}
