package message

import (
	"errors"

	"gorm.io/gorm"
)

type Repository interface {
	Store(message Message) (Message, error)
	FetchAll() ([]Message, error)
	FetchById(id int) (Message, error)
	Update(id int, inputMessage InputMessage) (Message, error)
	Delete(id int) error
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) Store(message Message) (Message, error) {
	//Check if the user exists
	//var user user.User
	// errUser := r.db.First(&user, message.SenderID).Error
	// if errUser != nil {
	// 	return message, fmt.Errorf("user does not exist")
	// }

	// // Check if the room exists
	// var room room.Room
	// errRoom := r.db.First(&room, message.RoomID).Error
	// if errRoom != nil {
	// 	return message, fmt.Errorf("room does not exist")
	// }

	// Create the message
	err := r.db.Create(&message).Error
	if err != nil {
		return message, err
	}

	return message, nil
}

func (r *repository) FetchAll() ([]Message, error) {
	var tasks []Message
	err := r.db.Find(&tasks).Error
	if err != nil {
		return tasks, err
	}

	return tasks, nil
}

func (r *repository) FetchById(id int) (Message, error) {
	var message Message

	err := r.db.Where(&Message{ID: id}).First(&message).Error
	if err != nil {
		return message, err
	}

	return message, nil
}

func (r *repository) Update(id int, inputMessage InputMessage) (Message, error) {
	message, err := r.FetchById(id)
	if err != nil {
		return message, err
	}

	message.Content = inputMessage.Content

	err = r.db.Save(&message).Error
	if err != nil {
		return message, err
	}

	return message, nil
}

func (r *repository) Delete(id int) error {
	message := &Message{ID: id}
	tx := r.db.Delete(message)
	if tx.Error != nil {
		return tx.Error
	}

	if tx.RowsAffected == 0 {
		return errors.New("Message not found")
	}

	return nil
}
