package server

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/jpodlasnisky/multithreading/config"
	"github.com/jpodlasnisky/multithreading/internal/infra/webservers/handlers"
)

func InitAddressServer(appConfig config.Config, addrHandle handlers.AddressHandle) {

	router := chi.NewRouter()
	router.Use(middleware.Logger)

	router.Route("/myaddress", func(router chi.Router) {
		router.Get("/{cep}", addrHandle.GetAddressHandle)
	})

	http.ListenAndServe(appConfig.WebServerPort, router)
}
