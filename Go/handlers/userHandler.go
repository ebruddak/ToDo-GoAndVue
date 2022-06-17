package handlers

import (
	"net/http"
	"time"

	"../dtos"
	"../services"
	"../utils"
	"github.com/gofiber/fiber"
)

type UserHandler struct {
	Service services.UserService
}

func (h UserHandler) CreateUser(c *fiber.Ctx) error {
	var user dtos.RegisterDTO

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
	if err := c.BodyParser(&user); err != nil {
		return err
	}
	result, err := h.Service.Login(user)

	if err != nil {
		return c.JSON(fiber.Map{
			"message": err.Error(),
		})
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

}

func (h UserHandler) User(c *fiber.Ctx) error {
	cookie := c.Cookies("jwt")

	id, _ := utils.ParseJwt(cookie)
	result, err := h.Service.User(id)
	if err != nil {
	}

	return c.JSON(result)
}

func (h UserHandler) Logout(c *fiber.Ctx) error {
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
