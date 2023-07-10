package user

import (
	"errors"

	"gorm.io/gorm"
)

type Repository interface {
	Store(user User) (User, error)
	FetchAll() ([]User, error)
	FetchById(id int) (User, error)
	Update(id int, inputUser InputUser) (User, error)
	Delete(id int) error
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) Store(user User) (User, error) {
	err := r.db.Create(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}

func (r *repository) FetchAll() ([]User, error) {
	var tasks []User
	err := r.db.Find(&tasks).Error
	if err != nil {
		return tasks, err
	}

	return tasks, nil
}

func (r *repository) FetchById(id int) (User, error) {
	var user User

	err := r.db.Where(&User{ID: id}).First(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}

func (r *repository) Update(id int, inputUser InputUser) (User, error) {
	user, err := r.FetchById(id)
	if err != nil {
		return user, err
	}

	user.Username = inputUser.Username
	user.Password = inputUser.Password

	err = r.db.Save(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}

func (r *repository) Delete(id int) error {
	user := &User{ID: id}
	tx := r.db.Delete(user)
	if tx.Error != nil {
		return tx.Error
	}

	if tx.RowsAffected == 0 {
		return errors.New("User not found")
	}

	return nil
}
