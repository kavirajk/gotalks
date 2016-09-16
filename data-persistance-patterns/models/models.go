package models

import (
	"log"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

var db *gorm.DB

func InitModel() {
	var err error
	db, err = gorm.Open("postgres", "user=kaviraj password=kaviraj database=badpatterns sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	if err := db.AutoMigrate(new(Picture), new(Album)).Error; err != nil {
		log.Fatal(err)
	}

}

// MODELSTART OMIT
type Model struct {
	Id        uint       `json:"id"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"` // HL
}

// MODELEND OMIT
