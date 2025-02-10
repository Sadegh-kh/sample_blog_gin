package main

import (
	"blog/api"
	"blog/server"
	"blog/service"
	"blog/storage/postgresql"
	"log"
)

func Init() server.Server {
	storage, err := postgresql.New()
	if err != nil {
		log.Fatalln("error when connect to storage :", err)
	}

	services := service.New(storage)

	ah := api.New(services.AuthService)

	return server.Server{ApiHanddler: ah}

}

func main() {
	server := Init()
	server.Run()
}
