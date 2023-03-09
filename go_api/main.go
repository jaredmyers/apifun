package main

import (
	"flag"
	"log"

	"github.com/jaredmyers/apifun/go_api/api"
	"github.com/jaredmyers/apifun/go_api/storage"
)

func main() {
	listenAddr := flag.String("listenAddr", ":8000", "server port")
	flag.Parse()

	store, err := storage.NewMySqlStore()
	if err != nil {
		log.Fatal(err)
	}

	userService := storage.NewUserService(store)

	server := api.NewServer(*listenAddr, userService)
	server.Run()
}
