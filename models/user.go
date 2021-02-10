package models

import (
	"log"
	"net/smtp"
	"os"
	"strings"
	"time"

	"github.com/Raise-the-Voices/raise-the-voices/database"
	u "github.com/Raise-the-Voices/raise-the-voices/utils"

	"github.com/dgrijalva/jwt-go"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

//User a struct to rep user account
type User struct {
	gorm.Model
	Name     string `json:"name"`
	Password string `json:"password"`
	Email    string `json:"email"`
	Phone    string `json:"phone,omitempty"`
	UserRole string `json:"user_role"`
}

//Password struct to can update user password for a new password
type Password struct {
	Email       string `json:"email"`
	OldPassword string `json:"old_password"`
	NewPassword string `json:"new_password"`
}

// Claims struct to JWT include email and userrole
type Claims struct {
	Email    string `json:"email"`
	UserRole string `json:"user_role"`
	jwt.StandardClaims
}

//ResetPassword a struct to reset password
type ResetPassword struct {
	gorm.Model
	UserID uint   `json:"user_ID"`
	Token  string `json:"token"`
	Email  string `json:"email"`
}

//Validate incoming user details...
func (user *User) Validate() (map[string]interface{}, bool) {
	user.Email = strings.ToLower(user.Email)
	user.UserRole = strings.ToLower(user.UserRole)

	if user.Name == "" {
		return u.Message("Name is required"), false
	}

	if len(user.Password) < 6 {
		return u.Message("Password must be at least 6 characters long"), false
	}

	if !strings.Contains(user.Email, "@") {
		return u.Message("Email address is required"), false
	}

	if !(user.UserRole == "admin" || user.UserRole == "editor") {
		return u.Message("User Role is required"), false
	}

	//Email must be unique
	temp := &User{}

	//check for errors and duplicate emails
	err := database.GetDB().Table("users").Where("email = ?", user.Email).First(temp).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return u.Message("Connection error. Please retry"), false
	}
	if temp.Email != "" {
		return u.Message("Email address already in use by another user."), false
	}

	return u.Message("Requirement passed"), true
}

//ValidateToUpdate user details...
func (user *User) ValidateToUpdate(idUserUpdate uint) (map[string]interface{}, bool) {
	user.Email = strings.ToLower(user.Email)
	user.UserRole = strings.ToLower(user.UserRole)
	userTemp := &User{}

	if user.Name == "" {
		return u.Message("Name is required"), false
	}

	if !strings.Contains(user.Email, "@") {
		return u.Message("Email address is required"), false
	}

	if !(user.UserRole == "admin" || user.UserRole == "editor") {
		return u.Message("User Role is required"), false
	}

	errToUpdate := database.GetDB().Table("users").Where("id = ?", idUserUpdate).First(userTemp).Error
	if errToUpdate != nil && errToUpdate != gorm.ErrRecordNotFound {
		return u.Message("Connection error. Please retry"), false
	}

	//Email must be unique
	temp := &User{}

	//check for errors and duplicate emails
	err := database.GetDB().Table("users").Where("email = ?", user.Email).First(temp).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return u.Message("Connection error. Please retry"), false
	}

	if temp.Email != "" && temp.Email != userTemp.Email {
		return u.Message("Email address already in use by another user."), false
	}

	return u.Message("Requirement passed"), true
}

//ValidatePassword password details...
func (password *Password) ValidatePassword() (map[string]interface{}, bool) {
	password.Email = strings.ToLower(password.Email)

	if !strings.Contains(password.Email, "@") {
		return u.Message("Email address is required"), false
	}

	if len(password.NewPassword) < 6 {
		return u.Message("Password must be at least 6 characters long"), false
	}

	if len(password.OldPassword) < 6 {
		return u.Message("Password must be at least 6 characters long"), false
	}

	//Email must be unique
	temp := &User{}

	//check for errors and duplicate emails
	err := database.GetDB().Table("users").Where("email = ?", password.Email).First(temp).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return u.Message("Email not found."), false
		}
		return u.Message("Connection error. Please retry"), false
	}

	err = bcrypt.CompareHashAndPassword([]byte(temp.Password), []byte(password.OldPassword))
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword { //Password does not match!
		return u.Message("Password does not match!"), false
	}

	return u.Message("Requirement passed"), true
}

// UserCreation register to DB of a new user
func (user *User) UserCreation() (map[string]interface{}, uint) {

	if resp, ok := user.Validate(); !ok {
		return resp, 400
	}

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	user.Password = string(hashedPassword)

	database.GetDB().Create(user)

	if user.ID <= 0 {
		return u.Message("Failed to create account, connection error."), 400
	}

	user.Password = "" //delete password

	response := u.Message("User has been created")
	response["user"] = user
	return response, 201
}

// UpdatePassword from a user
func (password *Password) UpdatePassword() (map[string]interface{}, uint) {

	if resp, ok := password.ValidatePassword(); !ok {
		return resp, 400
	}

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password.NewPassword), bcrypt.DefaultCost)
	password.NewPassword = string(hashedPassword)

	userToUpdate := &User{}

	errorUpdateUser := database.GetDB().Table("users").Where("email = ?", password.Email).First(&userToUpdate).Updates(
		User{
			Password: password.NewPassword}).Error

	if errorUpdateUser != nil {
		response := u.Message("User not found")
		return response, 404
	}

	response := u.Message("user Password have been updated")
	return response, 200
}

// UpdateUser from a user
func (user *User) UpdateUser(idUserUpdate uint) (map[string]interface{}, uint) {

	if resp, ok := user.ValidateToUpdate(idUserUpdate); !ok {
		return resp, 400
	}

	userToUpdate := &User{}

	errorUpdateUser := database.GetDB().Table("users").Where("id = ?", idUserUpdate).First(&userToUpdate).Error

	if errorUpdateUser != nil {
		response := u.Message("User not found")
		return response, 404
	}

	adminList := &[]User{}
	database.GetDB().Table("users").Where("user_role = ?", "admin").Find(&adminList)

	if len((*adminList)) < 2 && userToUpdate.UserRole != user.UserRole && userToUpdate.UserRole == "admin" {
		response := u.Message("Error! this is the last Administrator user please create an additional administrator user.")
		return response, 404
	}

	userToUpdate.Name = user.Name
	userToUpdate.Email = user.Email
	userToUpdate.Phone = user.Phone
	userToUpdate.UserRole = user.UserRole

	errorUpdateUser = database.GetDB().Table("users").Save(userToUpdate).Error
	if errorUpdateUser != nil {
		response := u.Message("Error! updating user")
		return response, 404
	}

	response := u.Message("user have been updated")
	return response, 200
}

// GetSpecificUser get specific user
func GetSpecificUser(iduser uint) (map[string]interface{}, uint) {

	specificUser := &User{}

	err := database.GetDB().Table("users").Where("id = ?", iduser).First(&specificUser).Error
	if err != nil {
		response := u.Message("User not found")
		return response, 404
	}

	specificUser.Password = ""

	response :=
		map[string]interface{}{
			"user": specificUser,
		}
	return response, 200
}

// SearchUserbyName search user by his name
func SearchUserbyName(nameOfUser string) (map[string]interface{}, uint) {

	specificUsers := &[]User{}

	err := database.GetDB().Order("id").Select([]string{"id", "created_at", "updated_at", "name", "email", "phone", "user_role"}).Table("users").Where("name ILIKE ?", "%"+nameOfUser+"%").Find(&specificUsers).Error
	if err != nil {
		response := u.Message("User not found")
		return response, 404
	}

	response :=
		map[string]interface{}{
			"users": specificUsers,
		}
	return response, 200
}

// GetAllUsers get all users
func GetAllUsers() (map[string]interface{}, uint) {

	usersList := &[]User{}

	err := database.GetDB().Table("users").Order("id").Find(&usersList).Error
	if err != nil {
		response := u.Message("users not found")
		return response, 404
	}

	for i := 0; i < len(*usersList); i++ {
		(*usersList)[i].Password = ""
	}

	response :=
		map[string]interface{}{
			"users": usersList,
		}
	return response, 200
}

// DeleteUserByID func to delete a user by its id
func DeleteUserByID(idUser uint) (map[string]interface{}, uint) {

	userToDelete := &User{}
	errorDeleteUser := database.GetDB().Table("users").Where("id = ?", idUser).First(&userToDelete).Error
	if errorDeleteUser != nil {
		response := u.Message("user Not Found")
		return response, 404
	}

	adminList := &[]User{}
	database.GetDB().Table("users").Where("user_role = ?", "admin").Find(&adminList)

	if userToDelete.UserRole == "admin" && len((*adminList)) < 2 {
		response := u.Message("Error! this is the last Administrator user please create an additional administrator user.")
		return response, 404
	}

	errorDeleteUser = database.GetDB().Delete(&userToDelete).Error
	if errorDeleteUser != nil {
		response := u.Message("Error! to delete user.")
		return response, 404
	}

	response := u.Message("the user was deleted successfully")
	return response, 200
}

// Login Function user
func Login(email, password string) (map[string]interface{}, uint) {

	user := &User{}
	err := database.GetDB().Table("users").Where("email = ?", email).First(user).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			response := u.Message("Email address not found")
			return response, 404
		}
		return u.Message("Connection error. Please retry"), 404
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword { //Password does not match!
		return u.Message("Invalid login credentials. Please try again"), 404
	}
	//Worked! Logged In
	user.Password = ""

	// Declare the expiration time of the token
	// here, we have kept it as 24 hours
	expirationTime := time.Now().Add(24 * time.Hour)
	// Create the JWT claims, which includes the username and expiry time
	claims := &Claims{
		Email:    user.Email,
		UserRole: user.UserRole,
		StandardClaims: jwt.StandardClaims{
			// In JWT, the expiry time is expressed as unix milliseconds
			ExpiresAt: expirationTime.Unix(),
		},
	}

	// Declare the token with the algorithm used for signing, and the claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// Create the JWT string
	tokenString, err := token.SignedString([]byte(os.Getenv("token_password")))
	if err != nil {
		// If there is an error in creating the JWT return an internal server error
		return u.Message("Internal Server Error"), 500
	}

	response :=
		map[string]interface{}{
			"token":   tokenString,
			"message": "Logged In",
			"expires": expirationTime,
			"role":    user.UserRole,
			"name":    user.Name,
		}

	return response, 200
}

// RefreshUserToken function to refresh time token
func RefreshUserToken(claims *Claims) (map[string]interface{}, uint) {

	if time.Unix(claims.ExpiresAt, 0).Sub(time.Now()) > 30*time.Second {
		return u.Message("Still have more than 30 seconds for the token time to expire"), 400
	}

	// Now, create a new token for the current use, with a renewed expiration time
	// Declare the expiration time of the token
	// here, we have kept it as 24 hours
	expirationTime := time.Now().Add(24 * time.Hour)
	claims.ExpiresAt = expirationTime.Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(os.Getenv("token_password")))
	if err != nil {
		return u.Message("Internal Server Error"), 500
	}

	response :=
		map[string]interface{}{
			"token":   tokenString,
			"message": "Refreshed token",
			"Expires": expirationTime,
		}

	return response, 200
}

// GenerateEmailToResetPassword to user email
func (user *User) GenerateEmailToResetPassword() (map[string]interface{}, uint) {

	if !strings.Contains(user.Email, "@") {
		response := u.Message("Email address is required")
		return response, 404
	}

	tempUser := &User{}

	//check for validate if exist Email
	err := database.GetDB().Table("users").Where("email = ?", user.Email).First(tempUser).Error
	if err != nil {

		if err == gorm.ErrRecordNotFound {
			if ok := send("Sorry someone is trying to enter your email on our website", user.Email); !ok {
				response := u.Message("Error to send email")
				return response, 404
			}

			response := u.Message("the email was sent successfully.")
			return response, 200
		}

		response := u.Message("Connection error. Please retry.")
		return response, 404
	}

	// Delete any existing reset tokens for the user
	tokenToDelete := &[]ResetPassword{}
	errorDeleteToken := database.GetDB().Table("reset_passwords").Where("email = ?", user.Email).Find(&tokenToDelete).Delete(&tokenToDelete).Error
	if errorDeleteToken != nil {
		// fmt.Println("Email Not Found")
	}

	token, err := u.GenerateRandomString()

	resetPassword := &ResetPassword{
		UserID: tempUser.ID,
		Token:  token,
		Email:  tempUser.Email}

	database.GetDB().Create(resetPassword)

	if ok := send("Click here to reset your password  https://raisethevoices.org/reset/"+token, user.Email); !ok {
		response := u.Message("Error to send email")
		return response, 404
	}

	response := u.Message("the email was sent successfully")
	return response, 200
}

// ValidateResetToken validate if the token is validate and not expired
func ValidateResetToken(token string) (map[string]interface{}, uint) {

	// Validate if exist the reset token for the user
	tokenToValidate := &ResetPassword{}
	errorValidateToken := database.GetDB().Table("reset_passwords").Where("token = ?", token).Find(&tokenToValidate).Error
	if errorValidateToken != nil {
		response := u.Message("Error! The Token was not found or not valid")
		return response, 404
	}

	TokenOneHour := tokenToValidate.CreatedAt.Add(1 * time.Hour)
	if time.Now().After(TokenOneHour) {
		response := u.Message("the Token is expired")
		return response, 401
	}

	response := u.Message("the Token is valid")
	response["user-email"] = tokenToValidate.Email
	return response, 200
}

// ResetPasswordFromUser with his email.
func (user *User) ResetPasswordFromUser() (map[string]interface{}, uint) {
	if resp, ok := user.ValidateResetPassword(); !ok {
		return resp, 400
	}

	// Delete any existing reset tokens for the user
	tokenToDelete := &[]ResetPassword{}
	errorDeleteToken := database.GetDB().Table("reset_passwords").Where("email = ?", user.Email).Find(&tokenToDelete).Delete(&tokenToDelete).Error
	if errorDeleteToken != nil {
		// fmt.Println("Email Not Found")
	}

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	user.Password = string(hashedPassword)

	userToUpdate := &User{}

	errorUpdateUser := database.GetDB().Table("users").Where("email = ?", user.Email).First(&userToUpdate).Updates(
		User{
			Password: user.Password}).Error

	if errorUpdateUser != nil {
		response := u.Message("Error! changing password.")
		return response, 404
	}

	user.Password = "" //delete password

	response := u.Message("User Password have been modified")
	return response, 200
}

//ValidateResetPassword password details...
func (user *User) ValidateResetPassword() (map[string]interface{}, bool) {
	user.Email = strings.ToLower(user.Email)

	if !strings.Contains(user.Email, "@") {
		return u.Message("Email address is required"), false
	}

	if len(user.Password) < 6 {
		return u.Message("Password must be at least 6 characters long"), false
	}

	temp := &User{}

	//check for if exist email
	err := database.GetDB().Table("users").Where("email = ?", user.Email).First(temp).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return u.Message("Email not found."), false
		}
		return u.Message("Connection error. Please retry"), false
	}

	return u.Message("Requirement passed"), true
}

func send(body string, toEmail string) bool {

	// add the email and password of the email to use to send reset password emails. ****Important

	from := "Email-From@gmail.com"
	pass := "Password-From-Email"
	to := toEmail

	msg := "From: " + from + "\n" +
		"To: " + to + "\n" +
		"Subject: Reset Password\n\n" +
		body

	err := smtp.SendMail("smtp.gmail.com:587",
		smtp.PlainAuth("", from, pass, "smtp.gmail.com"),
		from, []string{to}, []byte(msg))

	if err != nil {
		log.Printf("smtp error: %s", err)
		return false
	}
	return true
}
