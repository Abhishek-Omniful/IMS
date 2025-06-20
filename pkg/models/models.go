package models

import (
	"context"
	"errors"
	"fmt"
	"log"
	"strconv"

	"github.com/Abhishek-Omniful/IMS/mycontext"
	"github.com/Abhishek-Omniful/IMS/pkg/appinit"
	"github.com/omniful/go_commons/db/sql/postgres"
	"github.com/omniful/go_commons/redis"
	"gorm.io/gorm"
)

type Category struct {
	ID           int64  `json:"id"`
	CategoryName string `json:"category_name"`
}

type Product struct {
	ID                 int64  `json:"id"`
	ProductName        string `json:"product_name"`
	SellerId           int64  `json:"seller_id"`
	GeneralDescription string `json:"general_description"`
	CategoryID         int64  `json:"category_id"`
}

type Tenant struct {
	ID                int64  `json:"id"`
	TenantName        string `json:"tenant_name"`
	RegisteredAddress string `json:"registered_address"`
	TenantContact     string `json:"tenant_contact"`
	TenantEmail       string `json:"tenant_email"`
}

type Hub struct {
	ID             int64  `json:"id"`
	TenantID       int64  `json:"tenant_id"`
	ManagerName    string `json:"manager_name"`
	ManagerContact string `json:"manager_contact"`
	ManagerEmail   string `json:"manager_email"`
}

type Seller struct {
	ID            int64  `json:"id"`
	HubID         int64  `json:"hub_id"`
	TenantID      int64  `json:"tenant_id"`
	SellerName    string `json:"seller_name"`
	SellerContact string `json:"seller_contact"`
	SellerEmail   string `json:"seller_email"`
}

type SKU struct {
	ID          int64  `json:"id"`
	SellerID    int64  `json:"seller_id"`
	ProductID   int64  `json:"product_id"`
	Images      string `json:"images"`
	Description string `json:"description"`
	Fragile     bool   `json:"fragile"`
	Dimensions  string `json:"dimensions"`
}

type Inventory struct {
	SKUID     int64 `json:"sku_id"`
	HubID     int64 `json:"hub_id"`
	Quantity  int   `json:"quantity"`   //check and deautl are put
	UnitPrice int   `json:"unit_price"` //check and deautl are put
}

type Address struct {
	ID           int64  `json:"id"`
	EntityID     int64  `json:"entity_id"`
	EntityType   string `json:"entity_type"`
	AddressLine1 string `json:"address_line1"`
	AddressLine2 string `json:"address_line2"`
	Pincode      string `json:"pincode"`
	City         string `json:"city"`
	State        string `json:"state"`
	Country      string `json:"country"`
}

type ValidateOrderRequest struct {
	SKUID string `json:"sku_id"`
	HubID string `json:"hub_id"`
}

type ValidationResponse struct {
	IsValid bool
	Error   error
}

var db *postgres.DbCluster
var ctx context.Context
var redisClient *redis.Client

func init() {
	db = appinit.GetDB()
	if db == nil {
		log.Panic("Failed to connect to the database")
	}
	log.Println("Connected to the database successfully")
	// migrations.RunMigration() only once

	ctx = mycontext.GetContext()

	redisClient = appinit.GetRedis()
	log.Println("Connected to Redis successfully")
}

// hubs
func GetHubs() (*[]Hub, error) {
	var hubs []Hub
	result := db.GetMasterDB(ctx).Find(&hubs)
	return &hubs, result.Error

}

func CreateHub(hub *Hub) (*Hub, error) {
	result := db.GetMasterDB(ctx).Create(hub)
	return hub, result.Error
}

func DeleteHub(id int64) (*Hub, error) {
	var hub Hub
	result := db.GetMasterDB(ctx).Where("id = ?", id).Find(&hub)
	if result.Error != nil {
		return nil, result.Error
	}
	deleteError := db.GetMasterDB(ctx).Delete(&hub).Error
	if deleteError != nil {
		return nil, deleteError
	}
	return &hub, result.Error
}

func UpdateHub(hub *Hub) (*Hub, error) {
	result := db.GetMasterDB(ctx).Save(hub)
	if result.Error != nil {
		return nil, result.Error
	}
	return hub, nil
}

// tenants
func GetTenants() (*[]Tenant, error) {
	var tenants []Tenant
	result := db.GetMasterDB(ctx).Find(&tenants)
	return &tenants, result.Error
}

func CreateTenant(tenant *Tenant) (*Tenant, error) {
	result := db.GetMasterDB(ctx).Create(tenant)
	return tenant, result.Error
}

func DeleteTenant(id int64) (*Tenant, error) {
	var tenant Tenant
	result := db.GetMasterDB(ctx).Where("id = ?", id).Find(&tenant)
	if result.Error != nil {
		return nil, result.Error
	}
	delErr := db.GetMasterDB(ctx).Delete(&tenant).Error
	if delErr != nil {
		return nil, delErr
	}
	return &tenant, nil
}

func UpdateTenant(tenant *Tenant) (*Tenant, error) {
	result := db.GetMasterDB(ctx).Save(tenant)
	if result.Error != nil {
		return nil, result.Error
	}
	return tenant, nil
}

// skus
func GetSKUs() (*[]SKU, error) {
	var skus []SKU
	result := db.GetMasterDB(ctx).Find(&skus)
	return &skus, result.Error
}

func CreateSKU(sku *SKU) (*SKU, error) {
	result := db.GetMasterDB(ctx).Create(sku)
	return sku, result.Error
}

func DeleteSKU(id int64) (*SKU, error) {
	var sku SKU
	result := db.GetMasterDB(ctx).Where("id = ?", id).Find(&sku)
	if result.Error != nil {
		return nil, result.Error
	}
	deleteError := db.GetMasterDB(ctx).Delete(&sku).Error
	if deleteError != nil {
		return nil, deleteError
	}
	return &sku, result.Error
}

func UpdateSKU(sku *SKU) (*SKU, error) {
	result := db.GetMasterDB(ctx).Save(sku)
	if result.Error != nil {
		return nil, result.Error
	}
	return sku, nil
}

// seller
func GetSellers() (*[]Seller, error) {
	var sellers []Seller
	result := db.GetMasterDB(ctx).Find(&sellers)
	return &sellers, result.Error
}

// CreateSeller creates a new seller
func CreateSeller(seller *Seller) (*Seller, error) {
	result := db.GetMasterDB(ctx).Create(seller)
	return seller, result.Error
}

// UpdateSeller updates an existing seller
func UpdateSeller(seller *Seller) (*Seller, error) {
	result := db.GetMasterDB(ctx).Save(seller)
	if result.Error != nil {
		return nil, result.Error
	}
	return seller, nil
}

// DeleteSeller deletes a seller by ID
func DeleteSeller(id int64) (*Seller, error) {
	var seller Seller
	result := db.GetMasterDB(ctx).Where("id = ?", id).Find(&seller)
	if result.Error != nil {
		return nil, result.Error
	}
	deleteError := db.GetMasterDB(ctx).Delete(&seller).Error
	if deleteError != nil {
		return nil, deleteError
	}
	return &seller, nil
}

//products

func GetProducts() (*[]Product, error) {
	var products []Product
	result := db.GetMasterDB(ctx).Find(&products)
	return &products, result.Error
}

func CreateProduct(product *Product) (*Product, error) {
	result := db.GetMasterDB(ctx).Create(product)
	return product, result.Error
}

func UpdateProduct(product *Product) (*Product, error) {
	result := db.GetMasterDB(ctx).Save(product)
	if result.Error != nil {
		return nil, result.Error
	}
	return product, nil
}

func DeleteProduct(id int64) (*Product, error) {
	var product Product
	result := db.GetMasterDB(ctx).Where("id = ?", id).Find(&product)
	if result.Error != nil {
		return nil, result.Error
	}
	delErr := db.GetMasterDB(ctx).Delete(&product).Error
	if delErr != nil {
		return nil, delErr
	}
	return &product, nil
}

func Validator(hubid int64, skuid int64) bool {
	var inv Inventory
	err := db.GetMasterDB(ctx).Where("sku_id = ? AND hub_id = ?", skuid, hubid).First(&inv).Error
	return err == nil
}

// inventory
func UpsertInventory(inv *Inventory) error {
	db := db.GetMasterDB(ctx)

	// Try to find existing record
	var existing Inventory
	err := db.Where("sku_id = ? AND hub_id = ?", inv.SKUID, inv.HubID).First(&existing).Error

	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		// Insert new record
		return db.Create(inv).Error
	}

	// Update existing record
	return db.Model(&Inventory{}).
		Where("sku_id = ? AND hub_id = ?", inv.SKUID, inv.HubID).
		Updates(map[string]interface{}{
			"quantity":   inv.Quantity,
			"unit_price": inv.UnitPrice,
		}).Error
}

func GetInventoryByHub(hubID int64) ([]Inventory, error) {
	var inventory []Inventory
	err := db.GetMasterDB(ctx).Where("hub_id = ?", hubID).Find(&inventory).Error
	return inventory, err
}

func GetInventoryBySKU(skuID int64) ([]Inventory, error) {
	var inventory []Inventory
	err := db.GetMasterDB(ctx).Where("sku_id = ?", skuID).Find(&inventory).Error
	return inventory, err
}

func GetInventoryBySKUAndHub(skuID, hubID int64) (*Inventory, error) {
	var inv Inventory
	err := db.GetMasterDB(ctx).Where("sku_id = ? AND hub_id = ?", skuID, hubID).First(&inv).Error
	return &inv, err
}

func GetAllInventory() (*[]Inventory, error) {
	var inventory []Inventory
	err := db.GetMasterDB(ctx).Find(&inventory).Error
	return &inventory, err
}

// validate order
func ValidateOrder(order *ValidateOrderRequest) bool {
	hubid := order.HubID
	skuid := order.SKUID

	hubIDInt64, err := strconv.ParseInt(hubid, 10, 64)

	if err != nil {
		log.Printf("Error parsing hub ID: %v", err)
		return false
	}
	skuIDInt64, err := strconv.ParseInt(skuid, 10, 64)
	if err != nil {
		log.Printf("Error parsing SKU ID: %v", err)
		return false
	}

	// Compose a Redis key like "hub:123:sku:456"
	redisKey := fmt.Sprintf("hub:%s:sku:%s", hubid, skuid)

	// Check in Redis
	val, err := redisClient.Get(ctx, redisKey)
	log.Println("Checking Redis for validation:", redisKey)
	if err == nil && val == "valid" {
		log.Println("Order validated from Redis cache.")
		return true
	}

	isValid := Validator(hubIDInt64, skuIDInt64)

	if isValid {
		storeRedis(hubid, skuid)
		return true
	}
	return false

}

func storeRedis(hubID string, skuID string) {
	key := fmt.Sprintf("hub:%s:sku:%s", hubID, skuID)
	ok, err := redisClient.Set(ctx, key, "valid", 0)
	if err != nil || !ok {
		log.Panicf("Failed to store validation for hub %s and sku %s in Redis", hubID, skuID)
	}
	log.Printf("Stored validation for hub %s and sku %s in Redis", hubID, skuID)
}
