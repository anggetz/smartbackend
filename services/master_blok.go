package services

import "smartservice/model"

type Blok interface {
	Get() (*[]model.Blok, error)
	GetByGroupLocation(id int) (*[]model.Blok, error)
}
