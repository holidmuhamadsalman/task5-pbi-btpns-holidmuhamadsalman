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

func GetUsers(c *gin.Context) {
	var users []models.User
	database.DB.Preload("Photos").Find(&users)

	c.JSON(http.StatusOK, gin.H{"data": users})
}

func GetUserById(c *gin.Context) {
	id := c.Param("id")
	
	var user models.User
	if err := database.DB.Preload("Photos").First(&user, id).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"Status": "Failed", "Message": "Data Not Found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": user})
}

func UpdateUser(c *gin.Context) {
	id := c.Param("id")
	userInput := app.UserUpdate{
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

	if err := database.DB.Model(&user).Where("id = ?", id).Updates(&user).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"Status": "Failed", "Message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"Status": "Success", "Message": "User updated", "Data": user})
}

func DeleteUser(c *gin.Context) {
	id := c.Param("id")

	var user models.User
	if err := database.DB.First(&user, id).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"Status": "Failed", "Message": err.Error()})
		return
	}

	if err := database.DB.Unscoped().Delete(&user).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"Status": "Failed", "Message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"Status": "Success", "Message": "User deleted"})
}