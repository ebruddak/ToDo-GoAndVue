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
	dbClientGroup := configs.GetCollection(configs.DB, "groups")

	TodoRepositoryDb := repository.NewTodoRepositoryDb(dbClientTodo)
	UserRepositoryDb := repository.NewUserRepositoryDb(dbClientUser)
	GroupRepositoryDb := repository.NewGroupRepositoryDb(dbClientGroup)

	td := handlers.TodoHandler{Service: services.NewTodoService(TodoRepositoryDb)}
	us := handlers.UserHandler{Service: services.NewUserService(UserRepositoryDb)}
	gr := handlers.GroupHandler{Service: services.NewGroupService(GroupRepositoryDb)}

	appRoute.Post("/api/group", gr.CreateGroup)
	appRoute.Get("/api/groups", gr.GetAllGroup)
	appRoute.Delete("/api/group/:id", gr.DeleteGroup)
	appRoute.Get("/api/group/:id", gr.GetGroup)
	appRoute.Put("/api/group", gr.UpdateGroup)

	appRoute.Post("/api/todo", td.CreateTodo)
	appRoute.Get("/api/todos", td.GetAllTodo)
	appRoute.Get("/api/complated", td.GetAllComplatedTodo)
	appRoute.Delete("/api/todo/:id", td.DeleteTodo)
	appRoute.Get("/api/todo/:id", td.GetTodo)
	appRoute.Put("/api/todo", td.UpdateTodo)
	appRoute.Put("/api/complete/:id", td.Complete)

	appRoute.Post("/api/register", us.CreateUser)
	appRoute.Post("/api/login", us.Login)
	appRoute.Get("/api/user", us.User)
	appRoute.Post("/api/logout", us.Logout)

	appRoute.Listen(":3000")
}
