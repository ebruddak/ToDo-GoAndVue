package handlers

import (
	"../dtos"
	"../models"
	"../services"
	"github.com/gofiber/fiber"

	// "go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
	// "../utils"
	"time"
)

type UserHandler struct {
	Service services.UserService
}

func (h UserHandler) CreateUser(c *fiber.Ctx) error {
	var user models.User

	if err := c.BodyParser(&user); err != nil {
		return err
	}

	result, err := h.Service.UserInsert(user)

	if err != nil || result.Status == false {
		return err
	}

	return c.Status(http.StatusCreated).JSON(result)
}

func (h UserHandler) Login(c *fiber.Ctx) error {
	var user dtos.LoginDTO

	result, err := h.Service.Login(user)

	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	} else {
		cookie := fiber.Cookie{
			Name:     "jwt",
			Value:    result,
			Expires:  time.Now().Add(time.Hour * 24),
			HTTPOnly: true,
		}

		c.Cookie(&cookie)

		return c.JSON(fiber.Map{
			"message": "success",
		})
	}

	// return c.Status(http.StatusCreated).JSON(result)
}

// func User(c *fiber.Ctx) error {
// 	cookie := c.Cookies("jwt")

// 	id, _ := utils.ParseJwt(cookie)

// 	var user models.User

// 	// result, err := h.Service.User(id)

// 	database.DB.Where("id = ?", id).First(&user)

// 	return c.JSON(user)
// }

func Logout(c *fiber.Ctx) error {
	cookie := fiber.Cookie{
		Name:     "jwt",
		Value:    "",
		Expires:  time.Now().Add(-time.Hour),
		HTTPOnly: true,
	}

	c.Cookie(&cookie)

	return c.JSON(fiber.Map{
		"message": "success",
	})
}
