package main

import (
	"github.com/jpodlasnisky/multithreading/config"
	"github.com/jpodlasnisky/multithreading/internal/infra/database"
	"github.com/jpodlasnisky/multithreading/internal/infra/webservers/handlers"
	"github.com/jpodlasnisky/multithreading/internal/infra/webservers/server"
)

func main() {
	appConfig, err := config.LoadConfig(".")
	if err != nil {
		panic(err)
	}
	dbBrasil, dbViaCep := database.InitDBConnections(appConfig)
	addrHandle := handlers.NewAddressHandle(appConfig, dbViaCep, dbBrasil)
	server.InitAddressServer(appConfig, addrHandle)
}
