package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ServeHome(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Welcome to the IMS Service",
	})
}

// validators
func ValidateHubAndSKU(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Hub and SKU validated successfully",
	})
}

func ValidateAndUpdateInventory(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Inventory validated and updated successfully",
	})
}

// hubs
func GetHubs(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"hubs": []string{"Hub1", "Hub2", "Hub3"},
	})
}

func CreateHub(c *gin.Context) {
	c.JSON(http.StatusCreated, gin.H{
		"message": "Hub created successfully",
	})
}

func UpdateHub(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Hub updated successfully",
	})
}

func DeleteHub(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Hub deleted successfully",
	})
}

// skus
func GetSKUs(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"skus": []string{"SKU1", "SKU2", "SKU3"},
	})
}

func CreateSKU(c *gin.Context) {
	c.JSON(http.StatusCreated, gin.H{
		"message": "SKU created successfully",
	})
}

func UpdateSKU(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "SKU updated successfully",
	})
}

func DeleteSKU(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "SKU deleted successfully",
	})
}
