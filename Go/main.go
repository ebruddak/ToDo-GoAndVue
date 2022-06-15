package main

import (
	"./configs"
	"./handlers"
	"./repository"
	"./services"
	"github.com/gofiber/fiber"
	"github.com/gofiber/fiber/middleware/cors"
)

func main() {
	appRoute := fiber.New()

	appRoute.Use(cors.New(cors.Config{
		AllowCredentials: true,
	}))

	configs.ConnectDB()

	dbClientTodo := configs.GetCollection(configs.DB, "todos")
	dbClientUser := configs.GetCollection(configs.DB, "users")

	TodoRepositoryDb := repository.NewTodoRepositoryDb(dbClientTodo)
	UserRepositoryDb := repository.NewUserRepositoryDb(dbClientUser)

	td := handlers.TodoHandler{Service: services.NewTodoService(TodoRepositoryDb)}
	us := handlers.UserHandler{Service: services.NewUserService(UserRepositoryDb)}

	appRoute.Post("/api/todo", td.CreateTodo)
	appRoute.Get("/api/todos", td.GetAllTodo)
	appRoute.Delete("/api/todo/:id", td.DeleteTodo)

	appRoute.Post("/api/user", us.CreateUser)
	appRoute.Post("/api/login", us.Login)

	appRoute.Listen(":3000")
}
