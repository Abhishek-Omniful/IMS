package controllers

import (
	"strconv"

	"github.com/Abhishek-Omniful/IMS/pkg/models"
	"github.com/gin-gonic/gin"
	"github.com/omniful/go_commons/i18n"
)

// products
func GetProducts(c *gin.Context) {
	products, err := models.GetProducts()
	if err != nil {
		c.JSON(400, gin.H{"error": i18n.Translate(c, "Failed to fetch products")})
		return
	}
	c.JSON(200, gin.H{"products": products})
}

func CreateProduct(c *gin.Context) {
	var product = &models.Product{}
	if err := c.ShouldBindJSON(product); err != nil {
		c.JSON(400, gin.H{"error": i18n.Translate(c, "Failed to bind product JSON")})
		return
	}
	product, err := models.CreateProduct(product)
	if err != nil {
		c.JSON(400, gin.H{"error": i18n.Translate(c, "Failed to create product")})
		return
	}
	c.JSON(201, gin.H{"message": i18n.Translate(c, "Product created successfully"), "product": product})
}

func UpdateProduct(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{"error": i18n.Translate(c, "Invalid product ID")})
		return
	}
	var product = &models.Product{}
	if err := c.ShouldBindJSON(product); err != nil {
		c.JSON(400, gin.H{"error": i18n.Translate(c, "Failed to bind product JSON")})
		return
	}
	product.ID = id
	product, err = models.UpdateProduct(product)
	if err != nil {
		c.JSON(400, gin.H{"error": i18n.Translate(c, "Failed to update product")})
		return
	}
	c.JSON(200, gin.H{"message": i18n.Translate(c, "Product updated successfully"), "product": product})
}

func DeleteProduct(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{"error": i18n.Translate(c, "Invalid product ID")})
		return
	}
	product, err := models.DeleteProduct(id)
	if err != nil {
		c.JSON(400, gin.H{"error": i18n.Translate(c, "Failed to delete product")})
		return
	}
	c.JSON(200, gin.H{"message": i18n.Translate(c, "Product deleted successfully"), "product": product})
}
