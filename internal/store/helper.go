package store

import (
	"errors"

	"gorm.io/gorm"
)

var (
	ErrRecordNotFound = errors.New("No record found for given filter.")
)

// non generic structs can not have generic methods, thus this helper
// returns error if record does not exist
func findOne[T any](db *gorm.DB, filter any) (*T, error) {
	var entity T
	result := db.Where(filter).First(&entity)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, ErrRecordNotFound
		}
		return nil, result.Error
	}

	return &entity, nil
}

// returns empty slice if no record found
func findMany[T any](db *gorm.DB, filter any) ([]T, error) {
	var entities []T

	result := db.Where(filter).Find(&entities)

	if result.Error != nil {
		return nil, result.Error
	}

	return entities, nil
}
