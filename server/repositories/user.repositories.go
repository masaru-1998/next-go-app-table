package repositories

import (
	"myapp/models"

	"gorm.io/gorm"
)

type UserRepository interface {
				GetAllUserByEmail(users *[]models.User, email string) error
				CreateUser(user *models.User) error
}

type userRepository struct {
				db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository{
				return &userRepository{db}
}

/*
メールアドレスから全てのユーザを種痘
*/
func (userRepository *userRepository) GetAllUserByEmail(users *[]models.User, email string) error{
				if err := userRepository.db.Where("email = ?", email).Find(users).Error; err != nil {
								return err
				}
				return nil
}

/*
ユーザの登録処理
*/

func (userRepository *userRepository) CreateUser(user *models.User) error {
				if err := userRepository.db.Create(user).Error; err != nil {
					return err
				}
				return nil
}