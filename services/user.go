package services

import "smartservice/model"

type User interface {
	Create(User model.User)
}
