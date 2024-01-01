package controllers

import (
	"net/http"
	"path/filepath"
	"task5-pbi-btpns-holidmuhamadsalman/database"
	"task5-pbi-btpns-holidmuhamadsalman/models"

	"github.com/gin-gonic/gin"
)

func GetPhotos(c *gin.Context) {
	var photos []models.Photo
	database.DB.Find(&photos)

	c.JSON(http.StatusOK, gin.H{"data": photos})
}

func GetPhotoById(c *gin.Context) {
	id := c.Param("id")

	var photo models.Photo
	if err := database.DB.First(&photo, id).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"Status": "Failed", "Message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": photo})
}

func CreatePhoto(c *gin.Context) {
	reqID := c.GetUint("reqID")

	file, err := c.FormFile("photo_url")
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"Status": "Failed", "Message": err.Error()})
		return
	}

	filepath := filepath.Join("uploads", file.Filename)
	if err := c.SaveUploadedFile(file, filepath); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"Status": "Failed", "Message": err.Error()})
		return
	}

	newPhoto := models.Photo{
		Title:    c.PostForm("title"),
		Caption:  c.PostForm("caption"),
		PhotoUrl: filepath + file.Filename,
		UserID:   reqID,
	}

	if err := database.DB.Create(&newPhoto).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"Status": "Failed", "Message": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"Status": "Success", "Message": "New photo created", "Data": newPhoto})
}

func UpdatePhoto(c *gin.Context) {
	id := c.Param("id")
	reqID := c.GetUint("reqID")

	file, err := c.FormFile("photo_url")
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"Status": "Failed", "Message": err.Error()})
		return
	}

	filepath := filepath.Join("uploads", file.Filename)
	if err := c.SaveUploadedFile(file, filepath); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"Status": "Failed", "Message": err.Error()})
		return
	}

	updatePhoto := models.Photo{
		Title:    c.PostForm("title"),
		Caption:  c.PostForm("caption"),
		PhotoUrl: filepath + file.Filename,
		UserID:   reqID,
	}

	if err := database.DB.Model(&updatePhoto).Where("id = ?", id).Updates(&updatePhoto).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"Status": "Failed", "Message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"Status": "Success", "Message": "Photo updated", "Data": updatePhoto})
}

func DeletePhoto(c *gin.Context) {
	id := c.Param("id")

	var photo models.Photo
	if err := database.DB.First(&photo, id).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"Status": "Failed", "Message": err.Error()})
		return
	}

	if err := database.DB.Delete(&photo).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"Status": "Failed", "Message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"Status": "Success", "Message": "Photo deleted"})
}