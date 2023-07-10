package message

import "gorm.io/gorm"

type Service interface {
	Store(input InputMessage) (Message, error)
	FetchAll() ([]Message, error)
	FetchById(id int) (Message, error)
	Update(id int, inputTask InputMessage) (Message, error)
	Delete(id int) error
}

type service struct {
	repository Repository
}

func NewService(r Repository) *service {
	return &service{r}
}

func (s *service) Store(input InputMessage) (Message, error) {
	var message Message
	message.RoomID = input.RoomID
	message.SenderID = input.SenderID
	message.Content = input.Content

	newTask, err := s.repository.Store(message)
	if err != nil {
		return message, err
	}

	return newTask, nil
}

func (s *service) FetchAll() ([]Message, error) {
	tasks, err := s.repository.FetchAll()
	if err != nil {
		return tasks, err
	}

	return tasks, nil
}

func (s *service) FetchById(id int) (Message, error) {
	message, err := s.repository.FetchById(id)
	if err != nil {
		return message, err
	}

	return message, nil
}

func (s *service) Update(id int, input InputMessage) (Message, error) {
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

func GetMessagesByUserID(db *gorm.DB, userID int) ([]Message, error) {
	var messages []Message
	err := db.Where("sender_id = ?", userID).Find(&messages).Error
	if err != nil {
		return nil, err
	}
	return messages, nil
}
