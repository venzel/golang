package main

import (
	"log"
	"net/http"
	"webApp/src/routes"
)

func main() {
	router := routes.GetRoutes()

	log.Fatal(http.ListenAndServe(":3000", router))
}
