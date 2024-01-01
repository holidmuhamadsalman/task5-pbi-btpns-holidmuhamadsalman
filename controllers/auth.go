package controllers

import (
	"net/http"
	"task5-pbi-btpns-holidmuhamadsalman/app"
	"task5-pbi-btpns-holidmuhamadsalman/database"
	"task5-pbi-btpns-holidmuhamadsalman/helpers"
	"task5-pbi-btpns-holidmuhamadsalman/models"

	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context) {
	userInput := app.UserRegister{
		Username: c.PostForm("username"),
		Email:    c.PostForm("email"),
		Password: c.PostForm("password"),
	}

	if _, err := govalidator.ValidateStruct(userInput); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"Status": "Failed", "Message": err.Error()})
		return
	}

	hashedPassword, _ := helpers.HashPassword(userInput.Password)

	user := models.User{
		Username: userInput.Username,
		Email:    userInput.Email,
		Password: hashedPassword,
	}

	if err := database.DB.Create(&user).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"Status": "Failed", "Message": "Email has been created", "Error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"Status": "Success","Message": "User Created", "Data": user})
}

func Login(c *gin.Context)  {
	userInput := app.UserLogin{
		Email: c.PostForm("email"),
		Password: c.PostForm("password"),
	}

	if _, err := govalidator.ValidateStruct(userInput); err != nil{
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"Status": "Failed", "Message": "Accoun Not Found, Silahkan Register", "Error": err.Error()})
		return
	}

	var user models.User
	if err := database.DB.First(&user, "email = ?", userInput.Email).Error; err != nil{
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"Status": "Failed", "Message": "Invalid Email", "Error": err.Error()})
		return
	}

	if err := helpers.CheckPasswordHash(userInput.Password, user.Password); err != nil{
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"Status": "Failed", "Message": "Invalid Password", "Error": err.Error()})
		return
	}

	accessToken, _ := helpers.GenerateToken(user.ID)

	c.JSON(http.StatusCreated, gin.H{"Status": "Success", "Message": "Login Success", "userID": user.ID, "token": accessToken})

}