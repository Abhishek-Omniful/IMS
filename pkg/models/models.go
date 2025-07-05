package models

import (
	"context"

	"github.com/Abhishek-Omniful/IMS/mycontext"
	dbService "github.com/Abhishek-Omniful/IMS/pkg/integrations/db"
	redisService "github.com/Abhishek-Omniful/IMS/pkg/integrations/redis"
	"github.com/omniful/go_commons/db/sql/postgres"
	"github.com/omniful/go_commons/log"
	"github.com/omniful/go_commons/redis"
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
	db = dbService.GetDB()
	redisClient = redisService.GetRedis()
	ctx = mycontext.GetContext()
}
