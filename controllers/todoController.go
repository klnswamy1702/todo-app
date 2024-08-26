package controllers

import (
    "context"
    "encoding/json"
    "net/http"
    "go-todo-app/models"
    "log"
    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/bson/primitive"
    "github.com/gorilla/mux"
)

// Get collection from MongoDB
var collection *mongo.Collection

func SetCollection(c *mongo.Collection) {
    collection = c
}

// Create a new Todo
func CreateTodoEndpoint(response http.ResponseWriter, request *http.Request) {
    response.Header().Set("Content-Type", "application/json")
    var todo models.Todo
    _ = json.NewDecoder(request.Body).Decode(&todo)
    todo.ID = primitive.NewObjectID()
    result, err := collection.InsertOne(context.TODO(), todo)
    if err != nil {
        log.Fatal(err)
    }
    json.NewEncoder(response).Encode(result)
}

// Get all Todos
func GetTodosEndpoint(response http.ResponseWriter, request *http.Request) {
    response.Header().Set("Content-Type", "application/json")
    var todos []models.Todo
    cursor, err := collection.Find(context.TODO(), bson.M{})
    if err != nil {
        log.Fatal(err)
    }
    defer cursor.Close(context.TODO())
    for cursor.Next(context.TODO()) {
        var todo models.Todo
        cursor.Decode(&todo)
        todos = append(todos, todo)
    }
    json.NewEncoder(response).Encode(todos)
}

// Update a Todo
func UpdateTodoEndpoint(response http.ResponseWriter, request *http.Request) {
    response.Header().Set("Content-Type", "application/json")
    params := mux.Vars(request)
    id, _ := primitive.ObjectIDFromHex(params["id"])
    var todo models.Todo
    _ = json.NewDecoder(request.Body).Decode(&todo)
    filter := bson.M{"_id": id}
    update := bson.M{"$set": todo}
    result, err := collection.UpdateOne(context.TODO(), filter, update)
    if err != nil {
        log.Fatal(err)
    }
    json.NewEncoder(response).Encode(result)
}

// Delete a Todo
func DeleteTodoEndpoint(response http.ResponseWriter, request *http.Request) {
    response.Header().Set("Content-Type", "application/json")
    params := mux.Vars(request)
    id, _ := primitive.ObjectIDFromHex(params["id"])
    filter := bson.M{"_id": id}
    result, err := collection.DeleteOne(context.TODO(), filter)
    if err != nil {
        log.Fatal(err)
    }
    json.NewEncoder(response).Encode(result)
}

