package controllers

import (
	"net/http"
	"strconv"

	"github.com/Abhishek-Omniful/IMS/pkg/models"
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
	var hubs, err = models.GetHubs()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to fetch hubs"})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"hubs": hubs,
	})
}

func CreateHub(c *gin.Context) {
	var hub = &models.Hub{}
	err := c.ShouldBindBodyWithJSON(hub) //json bytes ->  struct (unmarshall)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to create hub",
		})
		return
	}

	hub, err = models.CreateHub(hub)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to create hub",
		})
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"message": "Hub created successfully",
		"hub":     hub,
	})
}

func DeleteHub(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Unable to parse hub ID",
		})
		return
	}

	hub, err := models.DeleteHub(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to delete hub",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Hub deleted successfully",
		"hub":     hub,
	})
}

func UpdateHub(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Unable to parse hub ID",
		})
		return
	}

	var hub = &models.Hub{}
	err = c.ShouldBindBodyWithJSON(hub) //json bytes ->  struct (unmarshall)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to update hub",
		})
		return
	}

	hub.ID = id // Set the ID for the update operation

	hub, err = models.UpdateHub(hub)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to update hub",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Hub updated successfully",
		"hub":     hub,
	})
}

// tenants
func GetTenants(c *gin.Context) {
	tenants, err := models.GetTenants()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to fetch tenants"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"tenants": tenants})
}

func CreateTenant(c *gin.Context) {
	var tenant = &models.Tenant{}
	err := c.ShouldBindJSON(tenant)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to bind tenant JSON"})
		return
	}

	tenant, err = models.CreateTenant(tenant)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to create tenant"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Tenant created successfully",
		"tenant":  tenant,
	})
}

func UpdateTenant(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid tenant ID"})
		return
	}

	var tenant = &models.Tenant{}
	err = c.ShouldBindJSON(tenant)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to bind tenant JSON"})
		return
	}

	tenant.ID = id
	tenant, err = models.UpdateTenant(tenant)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to update tenant"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Tenant updated successfully",
		"tenant":  tenant,
	})
}

func DeleteTenant(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid tenant ID"})
		return
	}

	tenant, err := models.DeleteTenant(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to delete tenant"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Tenant deleted successfully",
		"tenant":  tenant,
	})
}

// skus

func GetSKUs(c *gin.Context) {
	skus, err := models.GetSKUs()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to fetch SKUs",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"skus": skus,
	})
}

func CreateSKU(c *gin.Context) {
	var sku = &models.SKU{}
	err := c.ShouldBindBodyWithJSON(sku)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to create SKU",
		})
		return
	}

	sku, err = models.CreateSKU(sku)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to create SKU",
		})
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"message": "SKU created successfully",
		"sku":     sku,
	})
}

func DeleteSKU(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Unable to parse SKU ID",
		})
		return
	}

	sku, err := models.DeleteSKU(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to delete SKU",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "SKU deleted successfully",
		"sku":     sku,
	})
}

func UpdateSKU(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Unable to parse SKU ID",
		})
		return
	}

	var sku = &models.SKU{}
	err = c.ShouldBindBodyWithJSON(sku)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to update SKU",
		})
		return
	}

	sku.ID = id // Important for .Save()
	sku, err = models.UpdateSKU(sku)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to update SKU",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "SKU updated successfully",
		"sku":     sku,
	})
}
