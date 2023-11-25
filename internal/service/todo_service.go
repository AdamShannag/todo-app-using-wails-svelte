package service

import "todo-app/internal/db"

type TodoService struct {
	db *db.InMemoryDB
}

func NewTodoService(db *db.InMemoryDB) *TodoService {
	return &TodoService{db}
}

func (t *TodoService) GetTodos() []db.Todo {
	return t.db.All()
}

func (t *TodoService) AddTodo(todo string) {
	t.db.AddTodo(todo)
}

func (t *TodoService) RemoveTodo(id string) {
	t.db.RemoveTodo(id)
}

func (t *TodoService) ClearAll() {
	t.db.Clear()
}
