package storage

import (
	"context"
	"food-delivery/common"
	"food-delivery/modules/restaurants/model"
)

func (s *sqlStore) ListData(ctx context.Context,
	paging *common.Paging,
	filter *model.Filter,
	moreKeys ...string,
) ([]model.Restaurant, error) {
	var result []model.Restaurant

	db := s.db

	for i := range moreKeys {
		db = db.Preload(moreKeys[i])
	}

	if v := filter; v != nil {
		if v.CityId > 0 {
			db = db.Where("city_id = ?", v.CityId)
		}
	}

	if err := db.Table(model.Restaurant{}.TableName()).Count(&paging.Total).Error; err != nil {
		return nil, err
	}

	if err := db.
		Offset((paging.Page - 1) * paging.Limit).
		Limit(paging.Limit).
		Order("id desc").
		Find(&result).Error; err != nil {
		return nil, err
	}

	return result, nil
}
