package models

import (
	"errors"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// -- Inventory --

var UpsertInventory = func(inv *Inventory) error {
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

var GetInventoryByHub = func(hubID int64) ([]Inventory, error) {
	var inventory []Inventory
	err := db.GetMasterDB(ctx).Where("hub_id = ?", hubID).Find(&inventory).Error
	return inventory, err
}

var GetInventoryBySKU = func(skuID int64) ([]Inventory, error) {
	var inventory []Inventory
	err := db.GetMasterDB(ctx).Where("sku_id = ?", skuID).Find(&inventory).Error
	return inventory, err
}

var GetInventoryBySKUAndHub = func(skuID, hubID int64) (*Inventory, error) {
	var inv Inventory
	err := db.GetMasterDB(ctx).Where("sku_id = ? AND hub_id = ?", skuID, hubID).First(&inv).Error
	return &inv, err
}

var GetAllInventory = func() (*[]Inventory, error) {
	var inventory []Inventory
	err := db.GetMasterDB(ctx).Find(&inventory).Error
	return &inventory, err
}

var UpdateInventoryQuantity = func(skuID, hubID int64, quantityToDeduct int) error {
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

var CheckInventoryStatus = func(skuID, hubID int64, quantityDemanded int) bool {
	var inv Inventory
	err := db.GetMasterDB(ctx).Where("sku_id = ? AND hub_id = ?", skuID, hubID).First(&inv).Error
	if err != nil || inv.Quantity < quantityDemanded {
		return false
	}
	logger.Infof("Inventory check passed for SKU %d in Hub %d: Available Quantity = %d, Required Quantity = %d", skuID, hubID, inv.Quantity, quantityDemanded)
	UpdateInventoryQuantity(skuID, hubID, quantityDemanded)
	return true
}
