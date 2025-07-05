package controllers

import (
	"strconv"

	"github.com/Abhishek-Omniful/IMS/pkg/models"
	"github.com/gin-gonic/gin"
	"github.com/omniful/go_commons/i18n"
)

// tenants
func GetTenants(c *gin.Context) {
	tenants, err := models.GetTenants()
	if err != nil {
		c.JSON(400, gin.H{"error": i18n.Translate(c, "Failed to fetch tenants")})
		return
	}
	c.JSON(200, gin.H{"tenants": tenants})
}

func CreateTenant(c *gin.Context) {
	var tenant = &models.Tenant{}
	if err := c.ShouldBindJSON(tenant); err != nil {
		c.JSON(400, gin.H{"error": i18n.Translate(c, "Failed to bind tenant JSON")})
		return
	}

	tenant, err := models.CreateTenant(tenant)
	if err != nil {
		c.JSON(400, gin.H{"error": i18n.Translate(c, "Failed to create tenant")})
		return
	}

	c.JSON(201, gin.H{"message": i18n.Translate(c, "Tenant created successfully"), "tenant": tenant})
}

func UpdateTenant(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{"error": i18n.Translate(c, "Invalid tenant ID")})
		return
	}

	var tenant = &models.Tenant{}
	if err := c.ShouldBindJSON(tenant); err != nil {
		c.JSON(400, gin.H{"error": i18n.Translate(c, "Failed to bind tenant JSON")})
		return
	}

	tenant.ID = id
	tenant, err = models.UpdateTenant(tenant)
	if err != nil {
		c.JSON(400, gin.H{"error": i18n.Translate(c, "Failed to update tenant")})
		return
	}

	c.JSON(200, gin.H{"message": i18n.Translate(c, "Tenant updated successfully"), "tenant": tenant})
}

func DeleteTenant(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{"error": i18n.Translate(c, "Invalid tenant ID")})
		return
	}

	tenant, err := models.DeleteTenant(id)
	if err != nil {
		c.JSON(400, gin.H{"error": i18n.Translate(c, "Failed to delete tenant")})
		return
	}

	c.JSON(200, gin.H{"message": i18n.Translate(c, "Tenant deleted successfully"), "tenant": tenant})
}
