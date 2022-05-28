package user

import "gorm.io/gorm"

type Repository interface {
	AddUser(user User) (User, error)
	FindByEmail(emal string) (User, error)
	FindByID(id int) (User, error)
	FindAll() ([]User, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) AddUser(user User) (User, error) {
	err := r.db.Create(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}

func (r *repository) FindByID(id int) (User, error) {
	var user User

	err := r.db.Where("id = ?", id).Find(&user).Error
	if err != nil {
		return user, err
	}

	return user, err
}

func (r *repository) FindByEmail(email string) (User, error) {
	var user User

	err := r.db.Where("email = ?", email).Find(&user).Error
	if err != nil {
		return user, err
	}

	return user, err
}

func (r *repository) FindAll() ([]User, error) {
	var user []User

	err := r.db.Find(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}
