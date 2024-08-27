package services

import (
    "github.com/klnswamy1702/todo-app/backend/models"
    "github.com/klnswamy1702/todo-app/backend/repositories"
)

type TodoService interface {
    GetAllTodos() ([]models.Todo, error)
    GetTodoByID(id string) (models.Todo, error)
    CreateTodo(todo models.Todo) error
    UpdateTodo(id string, todo models.Todo) error
    DeleteTodo(id string) error
}

type todoService struct {
    repository repositories.TodoRepository
}

func NewTodoService(repo repositories.TodoRepository) TodoService {
    return &todoService{repository: repo}
}

func (s *todoService) GetAllTodos() ([]models.Todo, error) {
    return s.repository.GetAll()
}

func (s *todoService) GetTodoByID(id string) (models.Todo, error) {
    return s.repository.GetByID(id)
}

func (s *todoService) CreateTodo(todo models.Todo) error {
    return s.repository.Create(todo)
}

func (s *todoService) UpdateTodo(id string, todo models.Todo) error {
    return s.repository.Update(id, todo)
}

func (s *todoService) DeleteTodo(id string) error {
    return s.repository.Delete(id)
}
