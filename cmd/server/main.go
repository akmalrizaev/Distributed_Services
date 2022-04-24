package main

import (
	"log"

	"github.com/akmalrizaev/Distributed_Services/internal/server"
)

func main() {
	srv := server.NewHTTPServer(":8080")
	log.Fatal(srv.ListenAndServe())
}
