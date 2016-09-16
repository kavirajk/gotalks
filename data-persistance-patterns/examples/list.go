package main

import (
	"fmt"
	"log"

	"github.com/kavirajk/bad-patterns/models"
)

// MAINSTART OMIT
func main() {
	models.InitModel()
	albums, err := models.GetAllAlbums()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Albums")
	for i := range albums {
		fmt.Println(albums[i])
	}
	fmt.Println("Pictures")
	pics, err := models.GetAllPictures()
	if err != nil {
		log.Fatal(err)
	}
	for i := range pics {
		fmt.Println(pics[i])
	}
}

// MAINEND OMIT
