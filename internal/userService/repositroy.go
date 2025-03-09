package userService

import (
	"gorm.io/gorm"
)

type UserRepository interface {
	// CreateUser - Передаем в функцию user типа User из orm.go
	// возвращаем созданный User и ошибку
	CreateUser(user User) (User, error)
	// GetAllUsers - Возвращаем массив из всех задач в БД и ошибку
	GetAllUsers() ([]User, error)
	// UpdateUserByID - Передаем id и User, возвращаем обновленный User
	// и ошибку
	UpdateUserByID(id uint, user User) (User, error)
	// DeleteUserByID - Передаем id для удаления, возвращаем только ошибку
	DeleteUserByID(id uint) error
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *userRepository {
	return &userRepository{db: db}
}

func (r *userRepository) CreateUser(user User) (User, error) {
	result := r.db.Create(&user)
	if result.Error != nil {
		return User{}, result.Error
	}
	return user, nil
}

func (r *userRepository) GetAllUsers() ([]User, error) {
	var uss []User
	err := r.db.Find(&uss).Error
	return uss, err
}

func (r *userRepository) UpdateUserByID(id uint, user User) (User, error) {
	var existingUser User
	if err := r.db.First(&existingUser, id).Error; err != nil {
		return User{}, err
	}
	existingUser.Username = user.Username
	existingUser.Password = user.Password
	existingUser.Email = user.Email

	if err := r.db.Save(&existingUser).Error; err != nil {
		return User{}, err
	}

	return existingUser, nil
}

func (r *userRepository) DeleteUserByID(id uint) error {
	var user User
	result := r.db.Where("id = ?", id).Delete(&user)
	if result.Error != nil {
		return result.Error
	}
	return nil

}
