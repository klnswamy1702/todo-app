package repositories

import (
    "context"
    "github.com/klnswamy1702/todo-app/backend/config"
    "github.com/klnswamy1702/todo-app/backend/models"
    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/bson/primitive"
    "go.mongodb.org/mongo-driver/mongo"
)

type TodoRepository interface {
    GetAll() ([]models.Todo, error)
    GetByID(id string) (models.Todo, error)
    Create(todo models.Todo) error
    Update(id string, todo models.Todo) error
    Delete(id string) error
}

type todoRepository struct {
    db *mongo.Collection
}

func NewTodoRepository() TodoRepository {
    return &todoRepository{
        db: config.DB.Collection("todos"),
    }
}

func (r *todoRepository) GetAll() ([]models.Todo, error) {
    var todos []models.Todo
    cursor, err := r.db.Find(context.TODO(), bson.M{})
    if err != nil {
        return nil, err
    }
    defer cursor.Close(context.TODO())
    for cursor.Next(context.TODO()) {
        var todo models.Todo
        if err := cursor.Decode(&todo); err != nil {
            return nil, err
        }
        todos = append(todos, todo)
    }
    return todos, nil
}

func (r *todoRepository) GetByID(id string) (models.Todo, error) {
    var todo models.Todo
    objectID, err := primitive.ObjectIDFromHex(id)
    if err != nil {
        return todo, err
    }
    err = r.db.FindOne(context.TODO(), bson.M{"_id": objectID}).Decode(&todo)
    return todo, err
}

func (r *todoRepository) Create(todo models.Todo) error {
    _, err := r.db.InsertOne(context.TODO(), todo)
    return err
}

func (r *todoRepository) Update(id string, todo models.Todo) error {
    objectID, err := primitive.ObjectIDFromHex(id)
    if err != nil {
        return err
    }
    _, err = r.db.UpdateOne(context.TODO(), bson.M{"_id": objectID}, bson.M{"$set": todo})
    return err
}

func (r *todoRepository) Delete(id string) error {
    objectID, err := primitive.ObjectIDFromHex(id)
    if err != nil {
        return err
    }
    _, err = r.db.DeleteOne(context.TODO(), bson.M{"_id": objectID})
    return err
}
