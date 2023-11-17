package controller

import (
	// "fmt"
	// "math/rand"
	"net/mail"
	// "net/smtp"
	// "os"
	// "strconv"
	// "time"

	"github.com/Daviskelvin824/OldEgg/config"
	"github.com/Daviskelvin824/OldEgg/models"
	"github.com/gin-gonic/gin"
	// "github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

// func GetUser(c *gin.Context){
	
// }

func InsertUser(c *gin.Context){
	var newUser model.User;
	c.ShouldBindJSON(&newUser)

	if newUser.FirstName == "" {
		c.String(200, "Field Cannot be Null")
		return
	}

	if newUser.LastName == "" {
		c.String(200, "Field Cannot be Null")
		return
	}

	if newUser.Email == "" {
		c.String(200, "Field Cannot be Null")
		return
	}

	if newUser.MobilePhoneNumber == "" {
		c.String(200, "Field Cannot be Null")
		return
	}

	if newUser.Password == "" {
		c.String(200, "Field Cannot be Null")
		return
	}

	if newUser.RoleID == 0 {
		c.String(200, "Field Cannot be Null")
		return
	}

	if newUser.Status == "" {
		c.String(200, "Field Cannot be Null")
		return
	}


	var countEmail int64 = 0
	config.DB.Model(model.User{}).Where("email = ?", newUser.Email).Count(&countEmail)

	_, err := mail.ParseAddress(newUser.Email)
	if err != nil {
		c.String(200, "Invalid Email Format")
		return
	}

	if countEmail != 0 {
		c.String(200, "Email is Not Unique")
		return
	}

	var countPhone int64 = 0
	config.DB.Model(model.User{}).Where("mobile_phone_number = ?", newUser.MobilePhoneNumber).Count(&countPhone)

	if countPhone != 0 {
		c.String(200, "Mobile Phone Number is Not Unique")
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(newUser.Password), 10)

	if err != nil {
		panic(err)
	}

	newUser.Password = string(hashedPassword)

	config.DB.Create(&newUser)
	c.JSON(200, &newUser)

}
