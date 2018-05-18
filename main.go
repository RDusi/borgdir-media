package main

import (
	"log"
	"net/http"

	"github.com/jhoefker/borgdir-media/app/controller/test"
	"github.com/jhoefker/borgdir-media/app/route"
)

func main() {
	route.MapToController()
	test.Funktion1()
	log.Println("Listening on Port 3000 ...")
	http.ListenAndServe(":3000", nil)
}
