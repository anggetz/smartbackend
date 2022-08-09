package impl

import (
	"smartservice/core"
	"smartservice/model"
	"smartservice/services"
)

type GroupLocation struct{}

func NewMasterGroupLocation() services.GroupLocation {
	return &GroupLocation{}
}

func (g *GroupLocation) Get() (*[]model.GroupLocation, error) {
	db, err := core.NewDatabase().GetConnection()
	if err != nil {
		return nil, err
	}

	defer db.Close()

	groupLocations := []model.GroupLocation{}

	tx := db.GormQuery().Find(&groupLocations)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return &groupLocations, nil
}
