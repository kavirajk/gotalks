package main

import (
	"fmt"
	"log"

	"github.com/kavirajk/bad-patterns/models"
)

// MAINSTART OMIT
func main() {
	models.InitModel()
	p := models.Picture{AlbumId: 99, Caption: "No one wants me :("}
	if err := p.Save(); err != nil {
		log.Fatal(err)
	}
	fmt.Println(p)
	a, err := models.GetAlbum(99) // HL
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(a)
}

// MAINEND OMIT
