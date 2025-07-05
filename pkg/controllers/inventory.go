package controllers

import (
	"strconv"

	"github.com/Abhishek-Omniful/IMS/pkg/models"
	"github.com/gin-gonic/gin"
	"github.com/omniful/go_commons/i18n"
	"github.com/omniful/go_commons/log"
)

var logger = log.DefaultLogger()

// inventory
func UpsertInventory(c *gin.Context) {
	var inventory models.Inventory
	if err := c.ShouldBindJSON(&inventory); err != nil {
		c.JSON(400, gin.H{"error": i18n.Translate(c, "Invalid request payload")})
		return
	}
	if err := models.UpsertInventory(&inventory); err != nil {
		c.JSON(500, gin.H{"error": i18n.Translate(c, "Failed to upsert inventory")})
		return
	}
	c.JSON(200, gin.H{"message": i18n.Translate(c, "Inventory upserted successfully")})
}

func GetInventoryByHub(c *gin.Context) {
	hubID, err := strconv.ParseInt(c.Param("hub_id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{"error": i18n.Translate(c, "Invalid hub_id")})
		return
	}
	inventory, err := models.GetInventoryByHub(hubID)
	if err != nil {
		c.JSON(500, gin.H{"error": i18n.Translate(c, "Failed to fetch inventory")})
		return
	}
	c.JSON(200, inventory)
}

func GetInventoryBySKU(c *gin.Context) {
	skuID, err := strconv.ParseInt(c.Param("sku_id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{"error": i18n.Translate(c, "Invalid sku_id")})
		return
	}
	inventory, err := models.GetInventoryBySKU(skuID)
	if err != nil {
		c.JSON(500, gin.H{"error": i18n.Translate(c, "Failed to fetch inventory")})
		return
	}
	c.JSON(200, inventory)
}

func GetInventoryBySKUAndHub(c *gin.Context) {
	skuID, err1 := strconv.ParseInt(c.Param("sku_id"), 10, 64)
	hubID, err2 := strconv.ParseInt(c.Param("hub_id"), 10, 64)
	if err1 != nil || err2 != nil {
		c.JSON(400, gin.H{"error": i18n.Translate(c, "Invalid sku_id or hub_id")})
		return
	}
	inv, err := models.GetInventoryBySKUAndHub(skuID, hubID)
	if err != nil {
		c.JSON(500, gin.H{"error": i18n.Translate(c, "Failed to fetch inventory")})
		return
	}
	c.JSON(200, inv)
}

func GetAllInventory(c *gin.Context) {
	inventory, err := models.GetAllInventory()
	if err != nil {
		c.JSON(500, gin.H{"error": i18n.Translate(c, "Failed to fetch inventory")})
		return
	}
	c.JSON(200, inventory)
}

func CheckInventoryStatus(c *gin.Context) {
	skuID, err1 := strconv.ParseInt(c.Query("sku_id"), 10, 64)
	hubID, err2 := strconv.ParseInt(c.Query("hub_id"), 10, 64)
	quantity, err3 := strconv.Atoi(c.Query("quantity"))
	if err1 != nil || err2 != nil || err3 != nil {
		c.JSON(400, gin.H{"error": i18n.Translate(c, "Parsing error for sku_id, hub_id or quantity")})
		return
	}
	logger.Infof("Checking inventory for SKU: %d, Hub: %d, Quantity: %d", skuID, hubID, quantity)
	if models.CheckInventoryStatus(skuID, hubID, quantity) {
		c.JSON(200, gin.H{"IsValid": true, "message": i18n.Translate(c, "Inventory is available")})
	} else {
		c.JSON(404, gin.H{"IsValid": false, "message": i18n.Translate(c, "Inventory is not available")})
	}
}

func ValidateOrderRequest(c *gin.Context) {
	var validationResponse = &models.ValidationResponse{}
	skuID, err1 := strconv.ParseInt(c.Param("sku_id"), 10, 64)
	hubID, err2 := strconv.ParseInt(c.Param("hub_id"), 10, 64)
	if err1 != nil || err2 != nil {
		c.JSON(400, gin.H{"IsValid": false, "Error": i18n.Translate(c, "Unable to parse sku_id or hub_id")})
		return
	}
	if !models.ValidateOrderByHubAndSKU(hubID, skuID) {
		validationResponse.IsValid = false
		c.JSON(400, gin.H{"IsValid": false, "Error": i18n.Translate(c, "Invalid SKU or Hub")})
		return
	}
	validationResponse.IsValid = true
	validationResponse.Error = nil
	logger.Infof("Order validation successful for SKU: %d and Hub: %d", skuID, hubID)
	c.JSON(200, gin.H{"IsValid": true, "Error": nil})
}
