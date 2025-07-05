package models

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
