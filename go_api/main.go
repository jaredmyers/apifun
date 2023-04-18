package main

import (
	"flag"
	"log"

	"github.com/jaredmyers/apifun/go_api/api"
	"github.com/jaredmyers/apifun/go_api/services"
	"github.com/jaredmyers/apifun/go_api/storage"
)

func main() {
	listenAddr := flag.String("listenAddr", ":8000", "server port")
	flag.Parse()

	store, err := storage.NewMySqlStore("testdb")
	if err != nil {
		log.Fatal(err)
	}

	cache, err := storage.NewRedisCache()
	if err != nil {
		log.Fatal(err)
	}

	userService := services.NewUserService(store, cache)
	server := api.NewServer(*listenAddr, userService)
	server.RegisterRoutes()
	server.Run()
}
