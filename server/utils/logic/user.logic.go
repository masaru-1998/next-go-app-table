package logic

import "golang.org/x/crypto/bcrypt"

type UserLogic interface {
	CreateHashPassword(password string) []byte
}

type userLogic struct {}

func NewUserLogic() UserLogic {
	return &userLogic{}
}

func (userLogic *userLogic) CreateHashPassword(password string) []byte {
			hassPassword, _ :=  bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
			return hassPassword
}