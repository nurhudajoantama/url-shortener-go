package repository

import (
	"gorm.io/gorm"
)

type User struct {
	ID       string
	Name     string
	Usermame string
	Password string
}

type userRepository struct {
	DB *gorm.DB
}

type UserRepository interface {
	FindByUsername(username string, u *User) (*User, error)
	Create(u *User) (*User, error)
	// Update(user *User) (*User, error)
	// Delete(user *User) error
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{
		DB: db,
	}
}

func (r *userRepository) FindByUsername(username string, u *User) (*User, error) {
	res := r.DB.Where("username = ?", username).First(&u)
	if err := res.Error; err != nil {
		return nil, err
	}
	return u, nil
}

func (r *userRepository) Create(u *User) (*User, error) {
	res := r.DB.Create(&u)
	if err := res.Error; err != nil {
		return nil, err
	}
	return u, nil
}

// func (r *Repository) Create(user *User) (*User, error)
