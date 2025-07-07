package models

import "fmt"

func isValidHubSKUPair(hubid int64, skuid int64) bool {
	var inv Inventory
	err := db.GetMasterDB(ctx).Where("sku_id = ? AND hub_id = ?", skuid, hubid).First(&inv).Error
	return err == nil
}

var ValidateOrderByHubAndSKU = func(hubID int64, skuID int64) bool {
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
	isValid := isValidHubSKUPair(hubID, skuID)
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
