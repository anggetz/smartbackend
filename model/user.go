package model

import "database/sql"

type User struct {
	ID              uint
	Name            string
	NIK             string
	BlokID          int
	GroupLocationID int
	Blok            Blok
	GroupLocation   GroupLocation
	Jabatan         string
	RoleApp         string
	Email           string
	Password        string
	IsActive        bool
	Platform        sql.NullString
	*Base
}

func (u *User) TableName() string {
	return "sys_users"
}
