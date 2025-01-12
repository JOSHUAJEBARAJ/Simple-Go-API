package todo

import "errors"

type Service struct {
	todo []string
}

// constructor helps to access the variable
func NewService() *Service {
	return &Service{
		todo: make([]string, 0),
	}
}

func (svc *Service) Add(s string) error {
	for _, v := range svc.todo {
		if v == s {
			return errors.New("Duplicate entry is not allowed")
		}
	}
	svc.todo = append(svc.todo, s)
	return nil
}

func (svc *Service) GetAll() []string {
	return svc.todo
}
