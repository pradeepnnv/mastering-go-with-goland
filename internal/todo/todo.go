package todo

type Service struct {
	todos []string
}

func NewService() *Service {
	return &Service{
		todos: make([]string, 0),
	}
}

func (svc *Service) Add(todo string) {
	svc.todos = append(svc.todos, todo)
}

func (svc *Service) GetAll() []string {
	return svc.todos
}
