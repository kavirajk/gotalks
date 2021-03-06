package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/kavirajk/my-server/config"
	"github.com/kavirajk/my-server/models"
	"github.com/kavirajk/my-server/utils"
)

var (
	email    = flag.String("email", "", "user's email to reset the password")
	password = flag.String("password", "", "new password")
)

func main() {
	flag.Parse()
	if *email == "" {
		fmt.Print("email: ")
		fmt.Scanln(email)
	}
	if *password == "" {
		fmt.Print("new-password: ")
		fmt.Scanln(password)
	}
	models.InitModel(config.DBDriver, fmt.Sprintf(config.DBDataSource))
	user, err := models.GetUserByEmail(*email) // HL
	if err != nil {
		log.Fatalf("error getting user: %v\n", err)
	}
	user.Password = utils.HashString(*password)
	if err := user.Save(); err != nil { // HL
		log.Fatalf("error saving new password: %v\n", err)
	}
	fmt.Println("Password reset success.")
}
