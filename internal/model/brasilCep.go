package model

import "github.com/jpodlasnisky/multithreading/pkg/model"

type BrasilCepRes struct {
	Id           model.ID `json:"id"`
	Cep          string   `json:"cep"`
	State        string   `json:"state"`
	City         string   `json:"city"`
	Neighborhood string   `json:"neighborhood"`
	Street       string   `json:"street"`
	Service      string   `json:"service"`
}
