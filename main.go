package main

import (
	"log"
	"net/http"

	"github.com/jhoefker/borgdir-media/app/route"
)

func main() {
	route.MapToController()
	log.Println("Listening on Port 3000 ...")
	http.ListenAndServe(":3000", nil)
}
