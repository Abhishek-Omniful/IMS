package models

import (
	"context"
	"errors"
	"fmt"

	"github.com/Abhishek-Omniful/IMS/mycontext"
	"github.com/Abhishek-Omniful/IMS/pkg/services"
	"github.com/omniful/go_commons/db/sql/postgres"
	"github.com/omniful/go_commons/log"
	"github.com/omniful/go_commons/redis"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Product struct {
	ID                 int64  `json:"id"`
	ProductName        string `json:"product_name"`
	SellerId           int64  `json:"seller_id"`
	GeneralDescription string `json:"general_description"`
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
	TenantID  int64 `json:"tenant_id"`
	SkuID     int64 `json:"sku_id"`
	HubID     int64 `json:"hub_id"`
	Quantity  int   `json:"quantity"`
	UnitPrice int   `json:"unit_price"`
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
var logger *log.Logger

func init() {
	logger = log.DefaultLogger()
	db = services.GetDB()
	redisClient = services.GetRedis()
	ctx = mycontext.GetContext()
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

func CreateSeller(seller *Seller) (*Seller, error) {
	result := db.GetMasterDB(ctx).Create(seller)
	return seller, result.Error
}

func UpdateSeller(seller *Seller) (*Seller, error) {
	result := db.GetMasterDB(ctx).Save(seller)
	if result.Error != nil {
		return nil, result.Error
	}
	return seller, nil
}

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

// -- Inventory --

func UpsertInventory(inv *Inventory) error {
	db := db.GetMasterDB(ctx)
	var existing Inventory
	err := db.Where("sku_id = ? AND hub_id = ?", inv.SkuID, inv.HubID).First(&existing).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		return db.Create(inv).Error
	}
	return db.Model(&Inventory{}).
		Where("sku_id = ? AND hub_id = ?", inv.SkuID, inv.HubID).
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

func UpdateInventoryQuantity(skuID, hubID int64, quantityToDeduct int) error {
	logger.Infof("Updating inventory quantity for SKU: %d in Hub: %d by quantity: %d", skuID, hubID, quantityToDeduct)
	var inv Inventory
	tx := db.GetMasterDB(ctx).Begin()
	err := tx.Clauses(clause.Locking{Strength: "UPDATE"}).
		Where("sku_id = ? AND hub_id = ?", skuID, hubID).
		First(&inv).Error
	if err != nil {
		logger.Errorf("Error fetching inventory: %v", err)
		tx.Rollback()
		return err
	}
	inv.Quantity -= quantityToDeduct
	_ = tx.Model(&Inventory{}).
		Where("sku_id = ? AND hub_id = ?", skuID, hubID).
		Updates(map[string]interface{}{"quantity": inv.Quantity}).Error
	logger.Infof("Inventory updated successfully for SKU: %d in Hub: %d", skuID, hubID)
	return tx.Commit().Error
}

func CheckInventoryStatus(skuID, hubID int64, quantityDemanded int) bool {
	var inv Inventory
	err := db.GetMasterDB(ctx).Where("sku_id = ? AND hub_id = ?", skuID, hubID).First(&inv).Error
	if err != nil || inv.Quantity < quantityDemanded {
		return false
	}
	logger.Infof("Inventory check passed for SKU %d in Hub %d: Available Quantity = %d, Required Quantity = %d", skuID, hubID, inv.Quantity, quantityDemanded)
	UpdateInventoryQuantity(skuID, hubID, quantityDemanded)
	return true
}

func Validator(hubid int64, skuid int64) bool {
	var inv Inventory
	err := db.GetMasterDB(ctx).Where("sku_id = ? AND hub_id = ?", skuid, hubid).First(&inv).Error
	return err == nil
}

func ValidateOrder(hubID int64, skuID int64) bool {
	redisKey := fmt.Sprintf("hub:%d:sku:%d", hubID, skuID)
	val, err := redisClient.Get(ctx, redisKey)
	logger.Infof("Checking Redis for validation: %s", redisKey)
	if err == nil && val == "valid" {
		var hub Hub
		err := db.GetMasterDB(ctx).Where("id = ?", hubID).First(&hub).Error
		if err != nil {
			logger.Warnf("Hub does not exist: %d", hubID)
			redisClient.Del(ctx, redisKey)
			return false
		}
		var sku SKU
		err = db.GetMasterDB(ctx).Where("id = ?", skuID).First(&sku).Error
		if err != nil {
			logger.Warnf("SKU does not exist: %d", skuID)
			redisClient.Del(ctx, redisKey)
			return false
		}
		logger.Infof("Order validated from Redis cache.")
		return true
	}
	isValid := Validator(hubID, skuID)
	if isValid {
		storeRedis(hubID, skuID)
		return true
	}
	return false
}

func storeRedis(hubID int64, skuID int64) {
	key := fmt.Sprintf("hub:%d:sku:%d", hubID, skuID)
	ok, err := redisClient.Set(ctx, key, "valid", 0)
	if err != nil || !ok {
		logger.Errorf("Failed to store validation for hub %d and sku %d in Redis: %v", hubID, skuID, err)
	}
	logger.Infof("Stored validation for hub %d and sku %d in Redis", hubID, skuID)
}
