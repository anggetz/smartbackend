package impl

import (
	"smartservice/core"
	"smartservice/helpers"
	"smartservice/model"
	"smartservice/services"

	"github.com/dgrijalva/jwt-go"
)

type Auth struct {
}

func NewAuth() services.Auth {
	return &Auth{}
}

func (a *Auth) Login(username, password string) (string, error) {
	db, err := core.NewDatabase().GetConnection()
	if err != nil {
		return "", err
	}

	defer db.Close()

	user := &model.User{}

	tx := db.GormQuery().First(&user, " username = ? and password = ?", username, password)
	if tx.Error != nil {
		return "", tx.Error
	}

	//generate jwt
	token, err := helpers.NewJwtAuth().SetJwtKey("smartservicejwtkey").CreateToken(jwt.MapClaims{
		"username":  user.Name,
		"blokrumah": user.BlokID,
		"jabatan":   user.Jabatan,
		"groupuser": user.GroupLocationID,
	})
	if err != nil {
		return "", err
	}

	return token, nil
}

func (a *Auth) Register(user *model.User) (*model.User, error) {
	db, err := core.NewDatabase().GetConnection()
	if err != nil {
		return nil, err
	}

	encryptedPassword, err := helpers.NewEncrypt().SetKeyChiper("permatakwangsanchiper").Do(user.Password)

	user.Password = encryptedPassword

	defer db.Close()
	tx := db.GormQuery().Create(user)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return user, nil
}
