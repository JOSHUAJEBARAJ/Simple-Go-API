package todo

import (
	"context"
	"errors"
	"fmt"

	"github.com/JOSHUAJEBARAJ/Simple-Go-API/internal/db"
)

type Item struct {
	Task   string
	Status string
}

type Service struct {
	db *db.DB
}

// constructor helps to access the variable
func NewService(db *db.DB) *Service {
	return &Service{
		db: db,
	}
}

func (svc *Service) Add(s string) error {
	todos, err := svc.db.GetAllItems(context.Background())
	if err != nil {
		return err
	}
	for _, v := range todos {
		if v.Task == s {
			return errors.New("Duplicate entry is not allowed")
		}
	}
	if err := svc.db.InsertItem(context.Background(), db.Item{
		Task:   s,
		Status: "TO_BE_STARTED",
	}); err != nil {
		return fmt.Errorf("Failed to insert item %w", err)
	}
	return nil

}

// func (svc *Service) Search(s string) []string {
// 	var result []string
// 	for _, v := range svc.todo {
// 		if strings.Contains(v.Task, s) {
// 			result = append(result, v.Task)
// 		}
// 	}
// 	return result
// }

func (svc *Service) GetAll() ([]Item, error) {
	var results []Item
	items, err := svc.db.GetAllItems(context.Background())
	if err != nil {
		return nil, err
	}
	for _, item := range items {
		results = append(results, Item{
			Task:   item.Task,
			Status: item.Status,
		})
	}
	return results, nil
}
