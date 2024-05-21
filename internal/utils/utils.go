package utils

import (
	"encoding/json"
	"log"
	"net/http"
)

func AddJsonBodyIntoRes(viacepRes *interface{}, res http.ResponseWriter) {
	res.Header().Set("Content-Type", "application/json")

	address, err := json.Marshal(viacepRes)
	if err != nil {
		log.Println("Falha ao montar request para viacep:", err)
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}

	res.Write(address)
}
