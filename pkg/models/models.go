package models

import (
	"context"
	"log"

	"github.com/Abhishek-Omniful/IMS/mycontext"
	"github.com/Abhishek-Omniful/IMS/pkg/appinit"
	"github.com/omniful/go_commons/db/sql/postgres"
)

type Category struct {
	ID           int64  `json:"id"`
	CategoryName string `json:"category_name"`
}

type Product struct {
	ID                 int64  `json:"id"`
	ProductName        string `json:"product_name"`
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
	HubID       int64  `json:"hub_id"`
	SellerID    int64  `json:"seller_id"`
	ProductID   int64  `json:"product_id"`
	Images      string `json:"images"`
	Description string `json:"description"`
	UnitPrice   int    `json:"unit_price"`
	Fragile     bool   `json:"fragile"`
	Dimensions  string `json:"dimensions"`
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

func init() {
	db = appinit.GetDB()
	if db == nil {
		log.Panic("Failed to connect to the database")
	}
	log.Println("Connected to the database successfully")
	// migrations.RunMigration() only once
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
