package repository

import (
	"github.com/shabacha/pkg/domain/model"
	"github.com/shabacha/pkg/usecase/repository"
	"golang.org/x/crypto/bcrypt"

	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) repository.UserRepository {
	return &userRepository{db}
}

func (ur *userRepository) FindAll(u []*model.User) ([]*model.User, error) {
	err := ur.db.Find(&u).Error

	if err != nil {
		return nil, err
	}

	return u, nil
}

func (ur *userRepository) Create(u *model.User) (*model.User, error) {
	// Hash the user's password before storing it in the database
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	u.Password = string(hashedPassword)

	// Call Create() on the database object to store the user
	if err := ur.db.Create(u).Error; err != nil {
		return nil, err
	}

	return u, nil
}

func (ur *userRepository) GetById(id int) (*model.User, error) {
	var u model.User
	ur.db.Where("id = ?", id).First(&u)
	return &u, nil
}

func (ur *userRepository) Update(u *model.User, id int) (*model.User, error) {
	if err := ur.db.Model(&model.User{}).Where("id = ?", id).Updates(&u).Error; err != nil {
		return nil, err
	}
	return u, nil
}
func (ur *userRepository) Delete(id int) error {
	if err := ur.db.Model(&model.User{}).Where("id = ?", id).Delete(&model.User{}).Error; err != nil {
		return err
	}
	return nil
}
func (ur *userRepository) GetUserByUserName(u *model.User, username string) (*model.User, error) {
	if err := ur.db.Model(&model.User{}).Where("username = ?", username).First(&u).Error; err != nil {
		return nil, err
	}
	return u, nil
}
