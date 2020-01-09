package main

import (
	"fmt"
	"log"
	"net/http"

	"./router"
	"./setupdb"
)

func main() {
	setupdb.SetUpTable()
	r := router.Router()
	fmt.Println("Starting server on the port 8080...")

	log.Fatal(http.ListenAndServe(":8080", r))
}
