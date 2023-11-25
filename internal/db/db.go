package db

import (
	"sync"

	"github.com/google/uuid"
)

type Todo struct {
	ID        string `json:"id"`
	Text      string `json:"text"`
	Completed bool   `json:"completed"`
}

type InMemoryDB struct {
	sync.RWMutex

	db map[string]Todo
}

func NewInMemoryDB() *InMemoryDB {
	return &InMemoryDB{
		db: make(map[string]Todo),
	}
}

func (m *InMemoryDB) All() []Todo {
	todos := make([]Todo, len(m.db))
	i := 0
	for _, v := range m.db {
		todos[i] = v
		i++
	}
	return todos
}

func (m *InMemoryDB) AddTodo(todo string) {
	m.Lock()
	defer m.Unlock()

	newID := uuid.New().String()

	newTodo := Todo{
		ID:        newID,
		Text:      todo,
		Completed: false,
	}

	m.db[newID] = newTodo
}

func (m *InMemoryDB) RemoveTodo(id string) {
	m.Lock()
	defer m.Unlock()

	delete(m.db, id)
}

func (m *InMemoryDB) Clear() {
	m.Lock()
	defer m.Unlock()

	m.db = make(map[string]Todo)
}
