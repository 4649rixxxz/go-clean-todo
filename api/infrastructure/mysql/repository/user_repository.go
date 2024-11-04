package repository

import (
	userDomain "go-clean-todo/domain/user"
	"go-clean-todo/infrastructure/mysql"

	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository() userDomain.UserRepository {
	return &userRepository{
		db: mysql.GetDB(),
	}
}

func (r *userRepository) CreateUser(user *userDomain.User) (*userDomain.User, error) {
	userORM := mysql.User{
		Email:    user.Email(),
		Password: user.Password(),
	}
	err := r.db.Create(&userORM).Error
	if err != nil {
		return nil, err
	}

	return userDomain.Reconstruct(
		userORM.UserID,
		userORM.Email,
		userORM.Password,
		userORM.CreatedAt,
		userORM.UpdatedAt,
		&userORM.DeletedAt.Time,
	), nil
}

func (r *userRepository) FetchByEmail(email string) (*userDomain.User, error) {
	var userORM mysql.User
	if err := r.db.Where("email = ?", email).First(&userORM).Error; err != nil {
		return nil, err
	}

	return userDomain.Reconstruct(
		userORM.UserID,
		userORM.Email,
		userORM.Password,
		userORM.CreatedAt,
		userORM.UpdatedAt,
		&userORM.DeletedAt.Time,
	), nil
}

func (r *userRepository) FetchByUserID(userID uint) (*userDomain.User, error) {
	var userORM mysql.User
	if err := r.db.First(&userORM, userID).Error; err != nil {
		return nil, err
	}

	return userDomain.Reconstruct(
		userORM.UserID,
		userORM.Email,
		userORM.Password,
		userORM.CreatedAt,
		userORM.UpdatedAt,
		&userORM.DeletedAt.Time,
	), nil
}
