package main

import (
	"./handlers"
	"./configs"
	"./repository"
	"./services"
	"github.com/gofiber/fiber"
)

func main() {
	appRoute := fiber.New()
	configs.ConnectDB()

	dbClient := configs.GetCollection(configs.DB, "todos")
	TodoRepositoryDb := repository.NewTodoRepositoryDb(dbClient)

	td := handlers.TodoHandler{Service: services.NewTodoService(TodoRepositoryDb)}

	appRoute.Post("/api/todo", td.CreateTodo)
	appRoute.Get("/api/todos", td.GetAllTodo)
	appRoute.Delete("/api/todo/:id", td.DeleteTodo)
	appRoute.Listen(":3000")
}