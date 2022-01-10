package repository

import (
	"context"
	"gorm.io/gorm"
)

type DataRepo struct {
	db *gorm.DB
}

func NewDataRepo(db *gorm.DB) *DataRepo {
	return &DataRepo{db: db}
}

func (r *DataRepo) Exp(ctx context.Context) string {
	return "index"
}