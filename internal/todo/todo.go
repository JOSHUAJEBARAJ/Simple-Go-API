package todo

import "errors"

type Item struct {
	Task   string
	Status string
}

type Service struct {
	todo []Item
}

// constructor helps to access the variable
func NewService() *Service {
	return &Service{
		todo: make([]Item, 0),
	}
}

func (svc *Service) Add(s string) error {
	for _, v := range svc.todo {
		if v.Task == s {
			return errors.New("Duplicate entry is not allowed")
		}
	}
	svc.todo = append(svc.todo, Item{Task: s, Status: "TO BE DONE"})
	return nil
}

func (svc *Service) GetAll() []Item {
	return svc.todo
}
