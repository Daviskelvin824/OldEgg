package controller

import (
	// "fmt"
	"math/rand"
	// "log"
	// "fmt"
	"net/mail"
	"net/smtp"
	"os"
	"strconv"
	"time"

	"github.com/Daviskelvin824/OldEgg/config"
	"github.com/Daviskelvin824/OldEgg/models"
	"github.com/gin-gonic/gin"

	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

// func GetUser(c *gin.Context){

// }

func InsertUser(c *gin.Context){
	var newUser models.User;
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
	config.DB.Model(models.User{}).Where("email = ?", newUser.Email).Count(&countEmail)

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
	config.DB.Model(models.User{}).Where("mobile_phone_number = ?", newUser.MobilePhoneNumber).Count(&countPhone)

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


func SignIn(c *gin.Context) {
	
	var attempt, user models.User
	c.ShouldBindJSON(&attempt)
	
	if attempt.Email == "" || attempt.Password == "" {
		c.String(200, "Fields Cannot be Empty")
		return
	}

	config.DB.First(&user, "email = ?", attempt.Email)

	if user.ID == 0 {
		c.String(200, "Invalid Email Address")
		return
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(attempt.Password))
	if err != nil {
		c.String(200, "Invalid Password")
		return
	}

	if user.Status != "Active" {
		c.String(200, "You Are Banned")
		return
	}


	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"subject": user.Email,
		"expire":  time.Now().Add(time.Hour * 24 * 30).Unix(),
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("SECRETKEY")))
	if err != nil {
		c.String(200, "Failed to Create Token")
		return
	}

	c.String(200, tokenString)

}

func Authenticate(c *gin.Context) {

	user, _ := c.Get("user")
	c.JSON(200, user)

}

func SubscribeToNewsletter(c *gin.Context) {

	var user models.User
	c.ShouldBindJSON(&user)

	var temp models.User
	config.DB.Model(&models.User{}).Where("email = ?", user.Email).First(&temp)

	if temp.ID == 0 {
		c.String(200, "Email Not Found")
		return
	}

	if temp.SubscribedToEmailOffersAndDiscounts {
		c.String(200, "You Are Already Subscribed")
		return
	}

	config.DB.Model(&models.User{}).Where("email = ?", user.Email).Updates(map[string]interface{}{
		"subscribed_to_email_offers_and_discounts": true,
	})

	c.String(200, "You Are Now Subscribed")

}

func UpdateUser(c *gin.Context) {

	var user models.User
	c.ShouldBindJSON(&user)

	config.DB.Model(&models.User{}).Where("email = ?", user.Email).Updates(map[string]interface{}{
		"status":                    user.Status,
		"mobile_phone_number":       user.MobilePhoneNumber,
		"two_factor_authentication": user.TwoFactorAuthentication,
		"currency":                  user.Currency,
	})

	c.JSON(200, &user)

}

func RequestTFACode(c *gin.Context) {

	type RequestBody struct {
		Email string `json:"email"`
	}

	var requestBody RequestBody
	c.ShouldBindJSON(&requestBody)

	// Validate Email Must Not be Empty
	if requestBody.Email == "" {
		c.String(200, "Field Can't be Empty")
		return
	}

	// Validate User Must Exist
	var user models.User
	config.DB.Model(models.User{}).Where("email = ?", requestBody.Email).First(&user)
	if user.ID == 0 {
		c.String(200, "Email isn't Registered")
		return
	}

	// Generate Code
	var code models.OneTimeCode
	code.Email = user.Email
	code.Code = strconv.Itoa(100000 + rand.Intn(999999-100000))

	// Save Code to Database
	// Only Create if Doesn't Exist
	var countEmail int64
	config.DB.Model(models.OneTimeCode{}).Where("email = ?", code.Email).Count(&countEmail)

	if countEmail == 0 {

		config.DB.Create(&code)

	} else {

		var user models.OneTimeCode
		config.DB.Model(models.OneTimeCode{}).Where("email = ?", code.Email).First(&user)
		user.Code = code.Code
		config.DB.Save(&user)

	}

	// Send Code to Email
	auth := smtp.PlainAuth("", "oldegg2023@gmail.com", "hhxxpsrtsgtidtek", "smtp.gmail.com")

	msg := "Subject: " + "Two Factor Authentication Code" + "\n" + "\nHere is your Code: " + code.Code
	var to []string
	to = append(to, code.Email)

	err := smtp.SendMail(
		"smtp.gmail.com:587",
		auth,
		"oldegg2023@gmail.com",
		to,
		[]byte(msg),
	)

	if err != nil {
		c.String(200, "Send Error")
		return
	}

	c.String(200, "Email Sent Successfully")

}

func ChangePassword(c *gin.Context) {

	type RequestBody struct {
		Email       string `json:"email"`
		OldPassword string `json:"old_password"`
		NewPassword string `json:"new_password"`
	}

	var requestBody RequestBody
	c.ShouldBindJSON(&requestBody)

	// Get User with Email
	var user models.User
	config.DB.Model(models.User{}).Where("email = ?", requestBody.Email).First(&user)

	// Validate Old Password Must be Correct
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(requestBody.OldPassword))

	if err != nil {
		c.String(200, "Old Password is Wrong")
		return
	}

	// If Correct, Save new Password
	newHashedPassword, err := bcrypt.GenerateFromPassword([]byte(requestBody.NewPassword), 10)

	if err != nil {
		panic(err)
	}

	user.Password = string(newHashedPassword)
	config.DB.Save(&user)

	c.String(200, "Password Saved!")

}

func SignInWithOneTimeCode(c *gin.Context) {

	type RequestBody struct {
		Email string `json:"email"`
		Code  string `json:"code"`
	}

	var requestBody RequestBody
	c.ShouldBindJSON(&requestBody)

	var code models.OneTimeCode
	config.DB.Model(models.OneTimeCode{}).Where("email = ?", requestBody.Email).Where("code = ?", requestBody.Code).First(&code)

	if code.ID == 0 {
		c.String(200, "Invalid Code")
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"subject": code.Email,
		"expire":  time.Now().Add(time.Hour * 24 * 30).Unix(),
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("SECRETKEY")))
	if err != nil {
		c.String(200, "Failed to Create Token")
		return
	}

	// Check if Code is Still Valid
	_, _, _, _, min, _ := diff(code.UpdatedAt, time.Now())
	if min >= 15 {
		c.String(200, "Code is Not Longer Valid")
		return
	}

	c.String(200, tokenString)

}
func diff(a, b time.Time) (year, month, day, hour, min, sec int) {

	if a.Location() != b.Location() {
		b = b.In(a.Location())
	}
	if a.After(b) {
		a, b = b, a
	}
	y1, M1, d1 := a.Date()
	y2, M2, d2 := b.Date()

	h1, m1, s1 := a.Clock()
	h2, m2, s2 := b.Clock()

	year = int(y2 - y1)
	month = int(M2 - M1)
	day = int(d2 - d1)
	hour = int(h2 - h1)
	min = int(m2 - m1)
	sec = int(s2 - s1)

	if sec < 0 {
		sec += 60
		min--
	}

	if min < 0 {
		min += 60
		hour--
	}

	if hour < 0 {
		hour += 24
		day--
	}

	if day < 0 {
		t := time.Date(y1, M1, 32, 0, 0, 0, 0, time.UTC)
		day += 32 - t.Day()
		month--
	}

	if month < 0 {
		month += 12
		year--
	}

	return

}

func GetOneTimeSignInCode(c *gin.Context){
	type RequestBody struct {
		Email string `json:"email"`
	}

	var requestBody RequestBody
	c.ShouldBindJSON(&requestBody)

	var code models.OneTimeCode
	code.Email = requestBody.Email

	code.Code = strconv.Itoa(100000 + rand.Intn(999999-100000))

	// Only Create if Doesn't Exist
	var countEmail int64
	config.DB.Model(models.OneTimeCode{}).Where("email = ?", code.Email).Count(&countEmail)

	if countEmail == 0 {

		config.DB.Create(&code)

	} else {

		var user models.OneTimeCode
		config.DB.Model(models.OneTimeCode{}).Where("email = ?", code.Email).First(&user)
		user.Code = code.Code
		config.DB.Save(&user)

	}

	c.JSON(200, code)
}

func RequestForgotPasswordCode(c *gin.Context) {

	type RequestBody struct {
		Email string `json:"email"`
	}

	var requestBody RequestBody
	c.ShouldBindJSON(&requestBody)

	// Validate Email Must Not be Empty
	if requestBody.Email == "" {
		c.String(200, "Field Can't be Empty")
		return
	}

	// Validate User Must Exist
	var user models.User
	config.DB.Model(models.User{}).Where("email = ?", requestBody.Email).First(&user)
	if user.ID == 0 {
		c.String(200, "Email isn't Registered")
		return
	}

	// Generate Code
	var code models.OneTimeCode
	code.Email = user.Email
	code.Code = strconv.Itoa(100000 + rand.Intn(999999-100000))

	// Save Code to Database
	// Only Create if Doesn't Exist
	var countEmail int64
	config.DB.Model(models.OneTimeCode{}).Where("email = ?", code.Email).Count(&countEmail)

	if countEmail == 0 {

		config.DB.Create(&code)

	} else {

		var user models.OneTimeCode
		config.DB.Model(models.OneTimeCode{}).Where("email = ?", code.Email).First(&user)

		// 2 Minute Code Request Validation
		_, _, _, _, min, _ := diff(user.UpdatedAt, time.Now())


		if min < 2 {
			c.String(200, "You Can Only Request Code Every 2 Minutes")
			return
		}

		user.Code = code.Code
		config.DB.Save(&user)

	}

	// Send Code to Email
	auth := smtp.PlainAuth("", "oldegg2023@gmail.com", "hhxxpsrtsgtidtek", "smtp.gmail.com")

	msg := "Subject: " + "Forgot Password Code" + "\n" + "\nHere is your Code: " + code.Code
	var to []string
	to = append(to, code.Email)

	err := smtp.SendMail(
		"smtp.gmail.com:587",
		auth,
		"oldegg2023@gmail.com",
		to,
		[]byte(msg),
	)

	if err != nil {
		c.String(200, "Send Error")
		return
	}

	c.String(200, "Email Sent Successfully")

}

func ValidateForgotPasswordCode(c *gin.Context) {

	type RequestBody struct {
		Email string `json:"email"`
		Code  string `json:"code"`
	}

	var requestBody RequestBody
	c.ShouldBindJSON(&requestBody)

	var code models.OneTimeCode
	config.DB.Model(models.OneTimeCode{}).Where("email = ?", requestBody.Email).Where("code = ?", requestBody.Code).First(&code)

	if code.ID == 0 {
		c.String(200, "Invalid Code")
		return
	}

	// Check if Code is Still Valid
	_, _, _, _, min, _ := diff(code.UpdatedAt, time.Now())
	if min >= 5 {
		c.String(200, "Code is Not Longer Valid")
		return
	}

	c.String(200, "Valid Code")

}

func ResetPassword(c *gin.Context) {

	type RequestBody struct {
		Email       string `json:"email"`
		NewPassword string `json:"new_password"`
	}

	var requestBody RequestBody
	c.ShouldBindJSON(&requestBody)

	// Update the Password
	// Find
	var user models.User
	config.DB.Model(models.User{}).Where("email = ?", requestBody.Email).First(&user)

	newHashedPassword, err := bcrypt.GenerateFromPassword([]byte(requestBody.NewPassword), 10)

	if err != nil {
		panic(err)
	}

	// Set the new Password
	user.Password = string(newHashedPassword)

	// Save the user
	config.DB.Save(&user)

	c.String(200, "Password Saved!")

}
