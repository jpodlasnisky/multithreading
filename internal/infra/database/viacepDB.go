package database

import (
	"github.com/jpodlasnisky/multithreading/internal/model"
	"gorm.io/gorm"
)

type ViacepDB struct {
	DB gorm.DB
}

func NewDBViacep(db gorm.DB) ViacepDB {
	return ViacepDB{DB: db}
}

func (viacepdb ViacepDB) Create(viacep model.ViaCepRes) error {
	return viacepdb.DB.Create(viacep).Error
}
