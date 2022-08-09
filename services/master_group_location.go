package services

import "smartservice/model"

type GroupLocation interface {
	Get() (*[]model.GroupLocation, error)
}
