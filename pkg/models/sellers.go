package models

// seller
var GetSellers = func() (*[]Seller, error) {
	var sellers []Seller
	result := db.GetMasterDB(ctx).Find(&sellers)
	return &sellers, result.Error
}

var CreateSeller = func(seller *Seller) (*Seller, error) {
	result := db.GetMasterDB(ctx).Create(seller)
	return seller, result.Error
}

var UpdateSeller = func(seller *Seller) (*Seller, error) {
	result := db.GetMasterDB(ctx).Save(seller)
	if result.Error != nil {
		return nil, result.Error
	}
	return seller, nil
}

var DeleteSeller = func(id int64) (*Seller, error) {
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
