package todo

import (
	"errors"
	"strings"
)

type Item struct {
	Task   string
	Status string
}
type Service struct {
	todos []Item
}

func NewService() *Service {
	return &Service{
		todos: make([]Item, 0),
	}
}

func (svc *Service) Add(todo string) error {
	for _, t := range svc.todos {
		if strings.ToLower(t.Task) == strings.ToLower(todo) {
			return errors.New("todo is not unique")
		}
	}
	svc.todos = append(svc.todos, Item{
		Task:   todo,
		Status: "TO_BE_STARTED",
	})
	return nil
}

func (svc *Service) GetAll() []Item {
	return svc.todos
}

func (svc *Service) Search(query string) []string {
	var results []string
	for _, t := range svc.todos {
		if strings.Contains(strings.ToLower(t.Task), strings.ToLower(query)) {
			results = append(results, t.Task)
		}
	}
	return results
}
