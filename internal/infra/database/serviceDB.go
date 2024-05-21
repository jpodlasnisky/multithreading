package database

import (
	"github.com/jpodlasnisky/multithreading/config"
	"github.com/jpodlasnisky/multithreading/internal/model"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func InitDBConnections(appConfig config.Config) (BrasilcepDB, ViacepDB) {
	dbGorm, err := gorm.Open(sqlite.Open(appConfig.DBFilePath))
	if err != nil {
		panic(err)
	}
	dbGorm.AutoMigrate(model.ViaCepRes{}, model.BrasilCepRes{})

	dbBrasil := NewDBBrasilCep(*dbGorm)
	dbVia := NewDBViacep(*dbGorm)

	return dbBrasil, dbVia
}
