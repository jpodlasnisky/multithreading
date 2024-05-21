package database

import (
	"github.com/jpodlasnisky/multithreading/internal/model"
	"gorm.io/gorm"
)

type BrasilcepDB struct {
	DB gorm.DB
}

func NewDBBrasilCep(db gorm.DB) BrasilcepDB {
	return BrasilcepDB{DB: db}
}

func (brasilcepdb BrasilcepDB) Create(brasilcep model.BrasilCepRes) error {
	return brasilcepdb.DB.Create(brasilcep).Error
}
