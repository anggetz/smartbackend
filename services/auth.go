package services

import "smartservice/model"

type Auth interface {
	Login(username, password string) (string, error)
	Register(user *model.User) (*model.User, error)
}
