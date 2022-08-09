package impl

import (
	"smartservice/core"
	"smartservice/model"
	"smartservice/services"
)

type Blok struct{}

func NewMasterBlok() services.Blok {
	return &Blok{}
}

func (g *Blok) Get() (*[]model.Blok, error) {
	db, err := core.NewDatabase().GetConnection()
	if err != nil {
		return nil, err
	}

	defer db.Close()

	Bloks := []model.Blok{}

	tx := db.GormQuery().Find(&Bloks)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return &Bloks, nil
}

func (b *Blok) GetByGroupLocation(id int) (*[]model.Blok, error) {
	db, err := core.NewDatabase().GetConnection()
	if err != nil {
		return nil, err
	}

	defer db.Close()

	Bloks := []model.Blok{}

	tx := db.GormQuery().Find(&Bloks, "id_group_location = ?", id)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return &Bloks, nil

}
