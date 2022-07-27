package db

import (
	"github.com/prometheus/common/model"
	"gorm.io/gorm/clause"
)

func createOrUpdate(*record) error {
	return db.Get().Model(&model.Name{}).Clauses(clause.OnConflict{
		UpdateAll: true,
	}).Create(record).Error
}
