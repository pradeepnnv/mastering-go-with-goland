package todo_test

import (
	"github.com/spf13/viper"
	"log"
	"my-first-api/internal/db"
	"my-first-api/internal/todo"
	"reflect"
	"testing"
)

func InitializeDb() *db.DB {
	viper.SetEnvPrefix("DB")
	viper.AutomaticEnv()
	viper.SetDefault("PORT", "5432")

	if !viper.IsSet("USER") || !viper.IsSet("PASSWORD") || !viper.IsSet("HOST") || !viper.IsSet("NAME") {
		log.Fatal("DB Details not provided")
	}

	d, err := db.New(
		viper.GetString("USER"),
		viper.GetString("PASSWORD"),
		viper.GetString("HOST"),
		viper.GetInt("PORT"),
		viper.GetString("NAME"),
	)
	if err != nil {
		log.Fatal(err)
	}
	return d
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
			d := InitializeDb()
			svc := todo.NewService(d)
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