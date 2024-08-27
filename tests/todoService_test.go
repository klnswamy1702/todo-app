package tests

import (
    "errors"
    "github.com/klnswamy1702/todo-app/backend/models"
    "github.com/klnswamy1702/todo-app/backend/services"
    "github.com/stretchr/testify/assert"
    "testing"
)

// Mock Repository Implementation
type mockTodoRepository struct {
    todos []models.Todo
}

func (m *mockTodoRepository) GetAll() ([]models.Todo, error) {
    return m.todos, nil
}

func (m *mockTodoRepository) GetByID(id string) (models.Todo, error) {
    for _, todo := range m.todos {
        if todo.ID.Hex() == id {
            return todo, nil
        }
    }
    return models.Todo{}, errors.New("todo not found")
}

func (m *mockTodoRepository) Create(todo models.Todo) error {
    m.todos = append(m.todos, todo)
    return nil
}

func (m *mockTodoRepository) Update(id string, updatedTodo models.Todo) error {
    for i, todo := range m.todos {
        if todo.ID.Hex() == id {
            m.todos[i] = updatedTodo
            return nil
        }
    }
    return errors.New("todo not found")
}

func (m *mockTodoRepository) Delete(id string) error {
    for i, todo := range m.todos {
        if todo.ID.Hex() == id {
            m.todos = append(m.todos[:i], m.todos[i+1:]...)
            return nil
        }
    }
    return errors.New("todo not found")
}

// Test Suite
func TestGetAllTodos(t *testing.T) {
    mockRepo := &mockTodoRepository{
        todos: []models.Todo{
            {Title: "Test Todo 1", Description: "Description 1"},
            {Title: "Test Todo 2", Description: "Description 2"},
        },
    }
    service := services.NewTodoService(mockRepo)

    todos, err := service.GetAllTodos()

    assert.NoError(t, err)
    assert.Equal(t, 2, len(todos))
    assert.Equal(t, "Test Todo 1", todos[0].Title)
}

func TestGetTodoByID_Success(t *testing.T) {
    mockRepo := &mockTodoRepository{
        todos: []models.Todo{
            {ID: primitive.NewObjectID(), Title: "Test Todo 1", Description: "Description 1"},
        },
    }
    service := services.NewTodoService(mockRepo)

    todo, err := service.GetTodoByID(mockRepo.todos[0].ID.Hex())

    assert.NoError(t, err)
    assert.Equal(t, "Test Todo 1", todo.Title)
}

func TestGetTodoByID_NotFound(t *testing.T) {
    mockRepo := &mockTodoRepository{
        todos: []models.Todo{},
    }
    service := services.NewTodoService(mockRepo)

    _, err := service.GetTodoByID("non-existing-id")

    assert.Error(t, err)
    assert.Equal(t, "todo not found", err.Error())
}

func TestCreateTodo_Success(t *testing.T) {
    mockRepo := &mockTodoRepository{
        todos: []models.Todo{},
    }
    service := services.NewTodoService(mockRepo)

    newTodo := models.Todo{Title: "New Todo", Description: "New Description"}

    err := service.CreateTodo(newTodo)
    assert.NoError(t, err)

    todos, _ := service.GetAllTodos()
    assert.Equal(t, 1, len(todos))
    assert.Equal(t, "New Todo", todos[0].Title)
}

func TestUpdateTodo_Success(t *testing.T) {
    mockRepo := &mockTodoRepository{
        todos: []models.Todo{
            {ID: primitive.NewObjectID(), Title: "Test Todo 1", Description: "Description 1"},
        },
    }
    service := services.NewTodoService(mockRepo)

    updatedTodo := models.Todo{Title: "Updated Title", Description: "Updated Description"}
    err := service.UpdateTodo(mockRepo.todos[0].ID.Hex(), updatedTodo)

    assert.NoError(t, err)

    todo, _ := service.GetTodoByID(mockRepo.todos[0].ID.Hex())
    assert.Equal(t, "Updated Title", todo.Title)
    assert.Equal(t, "Updated Description", todo.Description)
}

func TestUpdateTodo_NotFound(t *testing.T) {
    mockRepo := &mockTodoRepository{
        todos: []models.Todo{},
    }
    service := services.NewTodoService(mockRepo)

    updatedTodo := models.Todo{Title: "Updated Title", Description: "Updated Description"}
    err := service.UpdateTodo("non-existing-id", updatedTodo)

    assert.Error(t, err)
    assert.Equal(t, "todo not found", err.Error())
}

func TestDeleteTodo_Success(t *testing.T) {
    mockRepo := &mockTodoRepository{
        todos: []models.Todo{
            {ID: primitive.NewObjectID(), Title: "Test Todo 1", Description: "Description 1"},
        },
    }
    service := services.NewTodoService(mockRepo)

    err := service.DeleteTodo(mockRepo.todos[0].ID.Hex())

    assert.NoError(t, err)

    todos, _ := service.GetAllTodos()
    assert.Equal(t, 0, len(todos))
}

func TestDeleteTodo_NotFound(t *testing.T) {
    mockRepo := &mockTodoRepository{
        todos: []models.Todo{},
    }
    service := services.NewTodoService(mockRepo)

    err := service.DeleteTodo("non-existing-id")

    assert.Error(t, err)
    assert.Equal(t, "todo not found", err.Error())
}
