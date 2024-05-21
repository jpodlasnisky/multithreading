package model

import "github.com/jpodlasnisky/multithreading/pkg/model"

type ViaCepRes struct {
	Id          model.ID `json:"id"`
	Cep         string   `json:"cep"`
	Logradouro  string   `json:"logradouro"`
	Complemento string   `json:"complemento"`
	Bairro      string   `json:"bairro"`
	Localidade  string   `json:"localidade"`
	Uf          string   `json:"uf"`
	Ibge        string   `json:"ibge"`
	Gia         string   `json:"gia"`
	Ddd         string   `json:"ddd"`
	Siafi       string   `json:"siafi"`
}
