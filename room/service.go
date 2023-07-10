package room

import (
	"golang/message"
	"os/user"

	"gorm.io/gorm"
)

type Service interface {
	Store(input InputRoom) (Room, error)
	FetchAll() ([]Room, error)
	FetchById(id int) (Room, error)
	Update(id int, inputTask InputRoom) (Room, error)
	Delete(id int) error
}

type service struct {
	repository Repository
}

func NewService(r Repository) *service {
	return &service{r}
}

func (s *service) Store(input InputRoom) (Room, error) {
	var room Room
	room.Name = input.Name

	newTask, err := s.repository.Store(room)
	if err != nil {
		return room, err
	}

	return newTask, nil
}

func (s *service) FetchAll() ([]Room, error) {
	tasks, err := s.repository.FetchAll()
	if err != nil {
		return tasks, err
	}

	return tasks, nil
}

func (s *service) FetchById(id int) (Room, error) {
	room, err := s.repository.FetchById(id)
	if err != nil {
		return room, err
	}

	return room, nil
}

func (s *service) Update(id int, input InputRoom) (Room, error) {
	uTask, err := s.repository.Update(id, input)
	if err != nil {
		return uTask, err
	}

	return uTask, nil
}

func (s *service) Delete(id int) error {
	err := s.repository.Delete(id)
	if err != nil {
		return err
	}

	return nil
}

func GetUsersByRoomID(db *gorm.DB, roomID int) ([]user.User, error) {
	var users []user.User
	err := db.Where("room_id = ?", roomID).Find(&users).Error
	if err != nil {
		return nil, err
	}
	return users, nil
}

func GetMessagesByRoomID(db *gorm.DB, roomID int) ([]message.Message, error) {
	var messages []message.Message
	err := db.Where("room_id = ?", roomID).Find(&messages).Error
	if err != nil {
		return nil, err
	}
	return messages, nil
}
