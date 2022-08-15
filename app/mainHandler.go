package app

import (
	"forLearnCurrent/models"
	"forLearnCurrent/services"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
	"net/http"
)

type MainHandler struct {
	Service services.MainService
}

func (h MainHandler) CreateMain(c *fiber.Ctx) error {
	var main models.Main

	if err := c.BodyParser(&main); err != nil {
		return c.Status(http.StatusBadRequest).JSON(err.Error())
	}

	result, err := h.Service.MainInsert(main)

	if err != nil || !result.Status {
		return err
	}

	return c.Status(http.StatusCreated).JSON(result)
}

func (h MainHandler) GetAllMain(c *fiber.Ctx) error {
	result, err := h.Service.MainGetAll()

	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(err.Error())
	}

	return c.Status(http.StatusOK).JSON(result)
}

func (h MainHandler) GetByIdMain(c *fiber.Ctx) error {
	query := c.Params("id")
	cnv, _ := primitive.ObjectIDFromHex(query)

	result, err := h.Service.MainGetById(cnv)

	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(err.Error())
	}

	return c.Status(http.StatusOK).JSON(result)
}

func (h MainHandler) DeleteMain(c *fiber.Ctx) error  {
	query := c.Params("id")
	cnv, _ := primitive.ObjectIDFromHex(query)

	result, err := h.Service.MainDelete(cnv)

	if err != nil || !result {
		log.Fatalln(err)
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"State": false})
	}
	return c.Status(http.StatusOK).JSON(fiber.Map{"State": true})
}