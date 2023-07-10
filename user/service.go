package user

type Service interface {
	Store(input InputUser) (User, error)
	FetchAll() ([]User, error)
	FetchById(id int) (User, error)
	Update(id int, inputTask InputUser) (User, error)
	Delete(id int) error
}

type service struct {
	repository Repository
}

func NewService(r Repository) *service {
	return &service{r}
}

func (s *service) Store(input InputUser) (User, error) {
	var user User
	user.Username = input.Username
	user.Password = input.Password

	newTask, err := s.repository.Store(user)
	if err != nil {
		return user, err
	}

	return newTask, nil
}

func (s *service) FetchAll() ([]User, error) {
	tasks, err := s.repository.FetchAll()
	if err != nil {
		return tasks, err
	}

	return tasks, nil
}

func (s *service) FetchById(id int) (User, error) {
	user, err := s.repository.FetchById(id)
	if err != nil {
		return user, err
	}

	return user, nil
}

func (s *service) Update(id int, input InputUser) (User, error) {
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
