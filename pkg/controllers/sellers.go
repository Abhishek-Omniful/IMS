package controllers

import (
	"strconv"

	"github.com/Abhishek-Omniful/IMS/pkg/models"
	"github.com/gin-gonic/gin"
	"github.com/omniful/go_commons/i18n"
)

// sellers
func GetSellers(c *gin.Context) {
	sellers, err := models.GetSellers()
	if err != nil {
		c.JSON(400, gin.H{"error": i18n.Translate(c, "Failed to fetch sellers")})
		return
	}
	c.JSON(200, gin.H{"sellers": sellers})
}

func CreateSeller(c *gin.Context) {
	var seller = &models.Seller{}
	if err := c.ShouldBindJSON(seller); err != nil {
		c.JSON(400, gin.H{"error": i18n.Translate(c, "Failed to bind seller JSON")})
		return
	}

	seller, err := models.CreateSeller(seller)
	if err != nil {
		c.JSON(400, gin.H{"error": i18n.Translate(c, "Failed to create seller")})
		return
	}

	c.JSON(201, gin.H{"message": i18n.Translate(c, "Seller created successfully"), "seller": seller})
}

func UpdateSeller(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{"error": i18n.Translate(c, "Invalid seller ID")})
		return
	}

	var seller = &models.Seller{}
	if err := c.ShouldBindJSON(seller); err != nil {
		c.JSON(400, gin.H{"error": i18n.Translate(c, "Failed to bind seller JSON")})
		return
	}

	seller.ID = id
	seller, err = models.UpdateSeller(seller)
	if err != nil {
		c.JSON(400, gin.H{"error": i18n.Translate(c, "Failed to update seller")})
		return
	}

	c.JSON(200, gin.H{"message": i18n.Translate(c, "Seller updated successfully"), "seller": seller})
}

func DeleteSeller(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{"error": i18n.Translate(c, "Invalid seller ID")})
		return
	}

	seller, err := models.DeleteSeller(id)
	if err != nil {
		c.JSON(400, gin.H{"error": i18n.Translate(c, "Failed to delete seller")})
		return
	}

	c.JSON(200, gin.H{"message": i18n.Translate(c, "Seller deleted successfully"), "seller": seller})
}
