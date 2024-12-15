package todo_test

import (
	"context"
	"my-first-api/internal/db"
	"my-first-api/internal/todo"
	"reflect"
	"testing"
)

type MockDB struct {
	items []db.Item
}

func (m *MockDB) InsertItem(ctx context.Context, item db.Item) error {
	m.items = append(m.items, item)
	return nil
}

func (m *MockDB) GetAllItems(ctx context.Context) ([]db.Item, error) {
	return m.items, nil
}

func TestService_Search(t *testing.T) {

	tests := []struct {
		name            string
		toDosToAdd      []string
		query           string
		expectedResults []string
	}{
		{
			name: "given a todo of shop and a search of sh, I should get a shop back",
			toDosToAdd: []string{
				"shop for a carrot",
				"feed the rabbit",
			},
			query: "sh",
			expectedResults: []string{
				"shop for a carrot",
			},
		},
		{
			name: "search with case insensitive",
			toDosToAdd: []string{
				"clean the home",
				"Clean the car",
			},
			query: "cl",
			expectedResults: []string{
				"clean the home",
				"Clean the car",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &MockDB{}
			svc := todo.NewService(m)
			for _, toDoToAdd := range tt.toDosToAdd {
				err := svc.Add(toDoToAdd)
				if err != nil {
					t.Error(err)
				}
			}
			r, err := svc.Search(tt.query)
			if err != nil {
				t.Error(err)
			}
			if got := r; !reflect.DeepEqual(got, tt.expectedResults) {
				t.Errorf("Search() = %v, want %v", got, tt.expectedResults)
			}
		})
	}
}
