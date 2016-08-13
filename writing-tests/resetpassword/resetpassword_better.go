// WARNING: This command is supposed to be used only by admin

package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/kavirajk/my-server/config"
	"github.com/kavirajk/my-server/models"
	"github.com/kavirajk/my-server/utils"
)

// START1 OMIT
type Saver interface { // HL
	Save() *models.AppError
}

var defaultUserGetter = func(email string) *models.User {
	user, err := models.GetUserByEmail(email)
	if err != nil {
		log.Fatalf("failed password-reset: %v\n", err)
	}
	return user
}

var defaultPasswordSaver = func(s Saver) {
	if err := s.Save(); err != nil {
		log.Fatalf("error saving new password: %v\n", err)
	}
}

// END1 OMIT
// START2 OMIT

var (
	email        = flag.String("email", "", "user's email to reset the password")
	password     = flag.String("password", "", "new password")
	userGetter   = defaultUserGetter    // HL
	savePassword = defaultPasswordSaver // HL
)

// END2 OMIT

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
	ResetPassword(*email, *password) // HL
	fmt.Println("Password reset success.")
}

// ResetPassword takes email and new-password and resets the password with new-password
func ResetPassword(email, newPassword string) {
	user := userGetter(email)
	user.Password = utils.HashString(newPassword)
	savePassword(user)
}
