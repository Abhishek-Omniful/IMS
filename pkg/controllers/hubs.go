package controllers

import (
	//"fmt"
	"strconv"

	"github.com/Abhishek-Omniful/IMS/pkg/models"
	"github.com/gin-gonic/gin"
	"github.com/omniful/go_commons/i18n"
)

func GetHubs(c *gin.Context) {
	hubs, err := models.GetHubs()
	if err != nil {
		logger.Error("Failed to fetch hubs", err)
		c.JSON(400, gin.H{"error": i18n.Translate(c, "Failed to fetch hubs")})
		return
	}
	c.JSON(200, gin.H{"hubs": hubs})
}

func CreateHub(c *gin.Context) {

	var hub = &models.Hub{}
	if err := c.ShouldBindJSON(hub); err != nil {
		c.JSON(400, gin.H{"error": i18n.Translate(c, "Failed to parse request")})
		return
	}

	hub, err := models.CreateHub(hub)
	if err != nil {
		c.JSON(400, gin.H{"error": i18n.Translate(c, "Failed to create hub")})
		return
	}

	c.JSON(201, gin.H{"message": i18n.Translate(c, "Hub created successfully"), "hub": hub})
}

func UpdateHub(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{"error": i18n.Translate(c, "Unable to parse hub ID")})
		return
	}

	var hub = &models.Hub{}
	if err := c.ShouldBindJSON(hub); err != nil {
		c.JSON(400, gin.H{"error": i18n.Translate(c, "Failed to update hub")})
		return
	}

	hub.ID = id
	hub, err = models.UpdateHub(hub)
	if err != nil {
		c.JSON(400, gin.H{"error": i18n.Translate(c, "Failed to update hub")})
		return
	}

	c.JSON(200, gin.H{"message": i18n.Translate(c, "Hub updated successfully"), "hub": hub})
}

func DeleteHub(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{"error": i18n.Translate(c, "Unable to parse hub ID")})
		return
	}

	hub, err := models.DeleteHub(id)
	if err != nil {
		c.JSON(400, gin.H{"error": i18n.Translate(c, "Failed to delete hub")})
		return
	}

	c.JSON(200, gin.H{"message": i18n.Translate(c, "Hub deleted successfully"), "hub": hub})
}
