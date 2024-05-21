package database

import "github.com/jpodlasnisky/multithreading/internal/model"

type ViaCepInterface interface {
	Create(viacepRes model.ViaCepRes) error
}

type BrasilCepInterface interface {
	Create(brasilcepRes model.BrasilCepRes) error
}
