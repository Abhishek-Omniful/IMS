package models

var GetProducts = func() (*[]Product, error) {
	var products []Product
	result := db.GetMasterDB(ctx).Find(&products)
	return &products, result.Error
}

var CreateProduct = func(product *Product) (*Product, error) {
	result := db.GetMasterDB(ctx).Create(product)
	return product, result.Error
}

var UpdateProduct = func(product *Product) (*Product, error) {
	result := db.GetMasterDB(ctx).Save(product)
	if result.Error != nil {
		return nil, result.Error
	}
	return product, nil
}

var DeleteProduct = func(id int64) (*Product, error) {
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
