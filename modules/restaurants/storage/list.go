package storage

import (
	"context"
	"food-delivery/modules/restaurants/model"
)

func (s *sqlStore) ListData(ctx context.Context) ([]model.Restaurant, error) {
	var result []model.Restaurant

	db := s.db

	if err := db.Find(&result).Error; err != nil {
		return nil, err
	}
	return result, nil
}
