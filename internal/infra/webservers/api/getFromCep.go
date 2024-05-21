package api

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"sync"

	modelID "github.com/jpodlasnisky/multithreading/pkg/model"

	"github.com/jpodlasnisky/multithreading/internal/model"
)

func GetAddressFromCEP(channel chan interface{}, resHandle http.ResponseWriter, url string, wg *sync.WaitGroup, isAPIBrasil bool) {

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Println("Falha ao montar request para viacep:", err)
		http.Error(resHandle, err.Error(), http.StatusInternalServerError)
		return
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Println("Falha ao buscar viacep:", err)
		http.Error(resHandle, err.Error(), http.StatusInternalServerError)
		return
	}
	defer res.Body.Close()

	resBody, err := io.ReadAll(res.Body)
	if err != nil {
		log.Println("Falha ao ler resposta do viacep:", err)
		http.Error(resHandle, err.Error(), http.StatusInternalServerError)
		return
	}

	var cepVia model.ViaCepRes
	var cepBrasil model.BrasilCepRes

	if isAPIBrasil {
		if err = json.Unmarshal(resBody, &cepBrasil); err != nil {
			log.Println("Falha ao ler resposta do BrasilCEP:", err)
			http.Error(resHandle, err.Error(), http.StatusInternalServerError)
			return
		}
		cepBrasil.Id = modelID.NewID()
		channel <- cepBrasil

	} else {

		if err = json.Unmarshal(resBody, &cepVia); err != nil {
			log.Println("Falha ao ler resposta do ViaCEP:", err)
			http.Error(resHandle, err.Error(), http.StatusInternalServerError)
			return
		}
		cepVia.Id = modelID.NewID()
		channel <- cepVia
	}

	wg.Done()

}
