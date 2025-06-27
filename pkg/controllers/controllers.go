package controllers

import (
	"strconv"

	"github.com/Abhishek-Omniful/IMS/pkg/models"
	"github.com/gin-gonic/gin"
	"github.com/omniful/go_commons/i18n"
	"github.com/omniful/go_commons/log"
)

var logger = log.DefaultLogger()

func ServeHome(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": i18n.Translate(c, "Welcome to the IMS Service"),
	})
}

// validators
func ValidateHubAndSKU(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": i18n.Translate(c, "Hub and SKU validated successfully"),
	})
}

func ValidateAndUpdateInventory(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": i18n.Translate(c, "Inventory validated and updated successfully"),
	})
}

// hubs
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
		c.JSON(400, gin.H{"error": i18n.Translate(c, "Unable to parse sku_id or hub_id")})
		return
	}
	if !models.ValidateOrder(hubID, skuID) {
		validationResponse.IsValid = false
		c.JSON(400, gin.H{"error": i18n.Translate(c, "Order validation failed"), "validationResponse": validationResponse})
		return
	}
	validationResponse.IsValid = true
	validationResponse.Error = nil
	logger.Infof("Order validation successful for SKU: %d and Hub: %d", skuID, hubID)
	c.JSON(200, gin.H{"IsValid": true})
}
