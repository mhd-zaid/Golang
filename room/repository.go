package room

import (
	"errors"

	"gorm.io/gorm"
)

type Repository interface {
	Store(room Room) (Room, error)
	FetchAll() ([]Room, error)
	FetchById(id int) (Room, error)
	Update(id int, inputRoom InputRoom) (Room, error)
	Delete(id int) error
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) Store(room Room) (Room, error) {
	err := r.db.Create(&room).Error
	if err != nil {
		return room, err
	}

	return room, nil
}

func (r *repository) FetchAll() ([]Room, error) {
	var tasks []Room
	err := r.db.Find(&tasks).Error
	if err != nil {
		return tasks, err
	}

	return tasks, nil
}

func (r *repository) FetchById(id int) (Room, error) {
	var room Room

	err := r.db.Where(&Room{ID: id}).First(&room).Error
	if err != nil {
		return room, err
	}

	return room, nil
}

func (r *repository) Update(id int, inputRoom InputRoom) (Room, error) {
	room, err := r.FetchById(id)
	if err != nil {
		return room, err
	}

	room.Name = inputRoom.Name

	err = r.db.Save(&room).Error
	if err != nil {
		return room, err
	}

	return room, nil
}

func (r *repository) Delete(id int) error {
	room := &Room{ID: id}
	tx := r.db.Delete(room)
	if tx.Error != nil {
		return tx.Error
	}

	if tx.RowsAffected == 0 {
		return errors.New("Room not found")
	}

	return nil
}
