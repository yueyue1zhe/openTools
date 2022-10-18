package dbutil

import (
	"database/sql/driver"
	"encoding/json"
	"gorm.io/gorm"
)

type SimpleCallback = func(db *gorm.DB) *gorm.DB

func SimpleCallbackEmpty(db *gorm.DB) *gorm.DB {
	return db
}

type JSONDbField struct {
}

func (p JSONDbField) Value() (driver.Value, error) {
	return json.Marshal(p)
}

func (p JSONDbField) Scan(data interface{}) error {
	return json.Unmarshal(data.([]byte), &p)
}
