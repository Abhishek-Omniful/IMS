package models

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
