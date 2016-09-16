package main

import (
	"fmt"
	"log"

	"github.com/kavirajk/bad-patterns/models"
)

// MAINSTART OMIT
func main() {
	models.InitModel()
	p := models.Picture{AlbumId: 99, Caption: "Where is my album?"} // HL
	if err := p.Save(); err != nil {
		log.Fatal(err)
	}
	fmt.Println(p)
}

// MAINEND OMIT
