package models

// skus
var GetSKUs = func() (*[]SKU, error) {
	var skus []SKU
	result := db.GetMasterDB(ctx).Find(&skus)
	return &skus, result.Error
}

var CreateSKU = func(sku *SKU) (*SKU, error) {
	result := db.GetMasterDB(ctx).Create(sku)
	return sku, result.Error
}

var DeleteSKU = func(id int64) (*SKU, error) {
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

var UpdateSKU = func(sku *SKU) (*SKU, error) {
	result := db.GetMasterDB(ctx).Save(sku)
	if result.Error != nil {
		return nil, result.Error
	}
	return sku, nil
}
