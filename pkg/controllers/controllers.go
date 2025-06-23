package controllers

import (
	"fmt"
	"log"
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
			"error": "Failed to parse request",
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

// sellers
func GetSellers(c *gin.Context) {
	sellers, err := models.GetSellers()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to fetch sellers"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"sellers": sellers})
}

func CreateSeller(c *gin.Context) {
	var seller = &models.Seller{}
	err := c.ShouldBindJSON(seller)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to bind seller JSON"})
		return
	}

	seller, err = models.CreateSeller(seller)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to create seller"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Seller created successfully",
		"seller":  seller,
	})
}

func UpdateSeller(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid seller ID"})
		return
	}

	var seller = &models.Seller{}
	err = c.ShouldBindJSON(seller)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to bind seller JSON"})
		return
	}

	seller.ID = id
	seller, err = models.UpdateSeller(seller)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to update seller"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Seller updated successfully",
		"seller":  seller,
	})
}

func DeleteSeller(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid seller ID"})
		return
	}

	seller, err := models.DeleteSeller(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to delete seller"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Seller deleted successfully",
		"seller":  seller,
	})
}

// products
func GetProducts(c *gin.Context) {
	products, err := models.GetProducts()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to fetch products"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"products": products})
}

func CreateProduct(c *gin.Context) {
	var product = &models.Product{}
	fmt.Println(product)
	err := c.ShouldBindJSON(product)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to bind product JSON"})
		return
	}

	product, err = models.CreateProduct(product)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to create product"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Product created successfully",
		"product": product,
	})
}

func UpdateProduct(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid product ID"})
		return
	}

	var product = &models.Product{}
	err = c.ShouldBindJSON(product)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to bind product JSON"})
		return
	}

	product.ID = id
	product, err = models.UpdateProduct(product)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to update product"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Product updated successfully",
		"product": product,
	})
}

func DeleteProduct(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid product ID"})
		return
	}

	product, err := models.DeleteProduct(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to delete product"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Product deleted successfully",
		"product": product,
	})
}

// inventory
func UpsertInventory(c *gin.Context) {
	var inventory models.Inventory

	if err := c.ShouldBindJSON(&inventory); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	err := models.UpsertInventory(&inventory)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to upsert inventory"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Inventory upserted successfully"})
}

func GetInventoryByHub(c *gin.Context) {
	hubIDStr := c.Param("hub_id")
	hubID, err := strconv.ParseInt(hubIDStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid hub_id"})
		return
	}

	inventory, err := models.GetInventoryByHub(hubID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch inventory"})
		return
	}

	c.JSON(http.StatusOK, inventory)
}

func GetInventoryBySKU(c *gin.Context) {
	skuIDStr := c.Param("sku_id")
	skuID, err := strconv.ParseInt(skuIDStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid sku_id"})
		return
	}

	inventory, err := models.GetInventoryBySKU(skuID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch inventory"})
		return
	}

	c.JSON(http.StatusOK, inventory)
}

func GetInventoryBySKUAndHub(c *gin.Context) {
	skuIDStr := c.Param("sku_id")
	hubIDStr := c.Param("hub_id")

	skuID, err1 := strconv.ParseInt(skuIDStr, 10, 64)
	hubID, err2 := strconv.ParseInt(hubIDStr, 10, 64)

	if err1 != nil || err2 != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid sku_id or hub_id"})
		return
	}

	inv, err := models.GetInventoryBySKUAndHub(skuID, hubID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch inventory"})
		return
	}

	c.JSON(http.StatusOK, inv)
}

func GetAllInventory(c *gin.Context) {
	inventory, err := models.GetAllInventory()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch inventory"})
		return
	}

	c.JSON(http.StatusOK, inventory)
}

func CheckInventoryStatus(c *gin.Context) {
	skuIDStr := c.Query("sku_id")
	hubIDStr := c.Query("hub_id")
	quantityStr := c.Query("quantity")

	skuID, err1 := strconv.ParseInt(skuIDStr, 10, 64)
	hubID, err2 := strconv.ParseInt(hubIDStr, 10, 64)
	quantity, err3 := strconv.Atoi(quantityStr)

	if err1 != nil || err2 != nil || err3 != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Parsing error for sku_id, hub_id or quantity"})
		return
	}
	log.Println("Checking inventory for SKU:", skuID, "Hub:", hubID, "Quantity:", quantity)
	isAvailable := models.CheckInventoryStatus(skuID, hubID, quantity)
	
	if isAvailable {
		c.JSON(http.StatusOK, gin.H{"IsValid": true, "message": "Inventory is available"})
	} else {
		c.JSON(http.StatusNotFound, gin.H{"IsValid": false, "message": "Inventory is not available"})
	}
}

// validate SKU and HUB
func ValidateOrderRequest(c *gin.Context) {
	var validationResponse = &models.ValidationResponse{}
	skuIDStr := c.Param("sku_id")
	hubIDStr := c.Param("hub_id")

	skuID, err1 := strconv.ParseInt(skuIDStr, 10, 64)
	hubID, err2 := strconv.ParseInt(hubIDStr, 10, 64)

	if err1 != nil || err2 != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Unable to parse sku_id or hub_id",
		})
		return
	}

	isValidOrder := models.ValidateOrder(hubID, skuID)

	if !isValidOrder {
		validationResponse.IsValid = false

		c.JSON(http.StatusBadRequest, gin.H{
			"error":              "Order validation failed",
			"validationResponse": validationResponse,
		})
		return
	}

	validationResponse.IsValid = true
	validationResponse.Error = nil
	log.Println("Order validation successful for SKU:", skuID, "and Hub:", hubID, "and is valid:", validationResponse.IsValid)
	c.JSON(http.StatusOK, gin.H{"IsValid": true})

}
