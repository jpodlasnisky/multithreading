package handlers

import (
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/go-chi/chi"
	"github.com/jpodlasnisky/multithreading/config"
	"github.com/jpodlasnisky/multithreading/internal/infra/database"
	"github.com/jpodlasnisky/multithreading/internal/infra/webservers/api"
	"github.com/jpodlasnisky/multithreading/internal/model"
	utils "github.com/jpodlasnisky/multithreading/internal/utils"
)

type AddressHandle struct {
	config      config.Config
	viacepDB    database.ViaCepInterface
	brasilcepDB database.BrasilCepInterface
}

func NewAddressHandle(appConfig config.Config, viacepDB database.ViaCepInterface, brasilcepDB database.BrasilCepInterface) AddressHandle {
	return AddressHandle{
		config:      appConfig,
		viacepDB:    viacepDB,
		brasilcepDB: brasilcepDB,
	}
}

func (handle *AddressHandle) GetAddressHandle(res http.ResponseWriter, req *http.Request) {
	cep := chi.URLParam(req, "cep")

	var wg sync.WaitGroup
	wg.Add(1)

	channelViacep := make(chan interface{})
	channelBrasilCep := make(chan interface{})

	go api.GetAddressFromCEP(channelViacep, res, handle.config.ViaCepHost+cep+"/json/", &wg, false)
	go api.GetAddressFromCEP(channelBrasilCep, res, handle.config.BrasilCepHost+cep, &wg, true)

	select {
	case viaCepRes := <-channelViacep:

		log.Println("viaCEPResponse:", viaCepRes)
		handle.viacepDB.Create(viaCepRes.(model.ViaCepRes))
		utils.AddJsonBodyIntoRes(&viaCepRes, res)

	case brasilCepRes := <-channelBrasilCep:

		log.Println("brasilCEPResponse:", brasilCepRes)
		handle.brasilcepDB.Create(brasilCepRes.(model.BrasilCepRes))
		utils.AddJsonBodyIntoRes(&brasilCepRes, res)

	case <-time.After(time.Second):

		log.Println("Timeout limit exceeded")
		http.Error(res, "Timeout limit exceeded", http.StatusRequestTimeout)
		return
	}

	wg.Wait()
}
