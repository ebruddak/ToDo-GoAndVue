package handlers

import (
	"net/http"

	"github.com/ebruddak/ToDo-GoAndVue/tree/main/Go/models"
	"github.com/ebruddak/ToDo-GoAndVue/tree/main/Go/services"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type GroupHandler struct {
	Service services.GroupService
}

func (h GroupHandler) CreateGroup(c *fiber.Ctx) error {
	var group models.Group

	if err := c.BodyParser(&group); err != nil {
		return c.Status(http.StatusBadRequest).JSON(err.Error())
	}

	result, err := h.Service.GroupInsert(group)

	if err != nil || result.Status == false {
		return err
	}

	return c.Status(http.StatusCreated).JSON(result)
}
func (h GroupHandler) UpdateGroup(c *fiber.Ctx) error {
	var group models.Group

	if err := c.BodyParser(&group); err != nil {
		return c.Status(http.StatusBadRequest).JSON(err.Error())
	}

	result, err := h.Service.GroupUpdate(group)

	if err != nil || result == false {
		return err
	}

	return c.Status(http.StatusCreated).JSON(result)
}
func (h GroupHandler) GetAllGroup(c *fiber.Ctx) error {

	query := c.Params("id")
	cnv, _ := primitive.ObjectIDFromHex(query)
	result, err := h.Service.GroupGetAll(cnv)

	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(err.Error())
	}

	return c.Status(http.StatusOK).JSON(result)
}
func (h GroupHandler) GetGroup(c *fiber.Ctx) error {
	query := c.Params("id")
	cnv, _ := primitive.ObjectIDFromHex(query)
	result, err := h.Service.GetGroup(cnv)

	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(err.Error())
	}

	return c.Status(http.StatusOK).JSON(result)
}

func (h GroupHandler) DeleteGroup(c *fiber.Ctx) error {
	query := c.Params("id")
	cnv, _ := primitive.ObjectIDFromHex(query)

	result, err := h.Service.GroupDelete(cnv)

	if err != nil || result == false {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"State": false})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{"State": true})
}
