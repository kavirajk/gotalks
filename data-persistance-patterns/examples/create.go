package main

import (
	"fmt"
	"log"

	"github.com/kavirajk/bad-patterns/models"
)

// MAINSTART OMIT
func main() {
	models.InitModel()
	a := models.Album{Title: "Welcome Gophers1"}
	if err := a.Save(); err != nil {
		log.Fatal(err)
	}
	p := models.Picture{AlbumId: a.Id, Caption: "Learning a lot!! #golang"}
	if err := p.Save(); err != nil {
		log.Fatal(err)
	}
	fmt.Println(a)
}

// MAINEND OMIT
