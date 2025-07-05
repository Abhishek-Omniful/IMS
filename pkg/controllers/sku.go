package controllers

import (
	"strconv"

	"github.com/Abhishek-Omniful/IMS/pkg/models"
	"github.com/gin-gonic/gin"
	"github.com/omniful/go_commons/i18n"
)

// skus
func GetSKUs(c *gin.Context) {
	skus, err := models.GetSKUs()
	if err != nil {
		c.JSON(400, gin.H{"error": i18n.Translate(c, "Failed to fetch SKUs")})
		return
	}
	c.JSON(200, gin.H{"skus": skus})
}

func CreateSKU(c *gin.Context) {
	var sku = &models.SKU{}
	if err := c.ShouldBindJSON(sku); err != nil {
		c.JSON(400, gin.H{"error": i18n.Translate(c, "Failed to create SKU")})
		return
	}

	sku, err := models.CreateSKU(sku)
	if err != nil {
		c.JSON(400, gin.H{"error": i18n.Translate(c, "Failed to create SKU")})
		return
	}

	c.JSON(201, gin.H{"message": i18n.Translate(c, "SKU created successfully"), "sku": sku})
}

func UpdateSKU(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{"error": i18n.Translate(c, "Unable to parse SKU ID")})
		return
	}

	var sku = &models.SKU{}
	if err := c.ShouldBindJSON(sku); err != nil {
		c.JSON(400, gin.H{"error": i18n.Translate(c, "Failed to update SKU")})
		return
	}

	sku.ID = id
	sku, err = models.UpdateSKU(sku)
	if err != nil {
		c.JSON(400, gin.H{"error": i18n.Translate(c, "Failed to update SKU")})
		return
	}

	c.JSON(200, gin.H{"message": i18n.Translate(c, "SKU updated successfully"), "sku": sku})
}

func DeleteSKU(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{"error": i18n.Translate(c, "Unable to parse SKU ID")})
		return
	}

	sku, err := models.DeleteSKU(id)
	if err != nil {
		c.JSON(400, gin.H{"error": i18n.Translate(c, "Failed to delete SKU")})
		return
	}

	c.JSON(200, gin.H{"message": i18n.Translate(c, "SKU deleted successfully"), "sku": sku})
}
